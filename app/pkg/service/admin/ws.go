package admin

import (
	"saas/app/pkg/do"
	"sync"
)

type ws struct {
	entity            map[string]*do.Entity
	entityChan        chan []*do.Entity
	requestEntityChan chan struct{}
	enteringChan      chan *do.Entity
	leavingChan       chan *do.Entity

	msgChan chan *do.Message
	lock    sync.RWMutex
}

var Ws = &ws{
	entity:            make(map[string]*do.Entity),
	entityChan:        make(chan []*do.Entity),
	requestEntityChan: make(chan struct{}),
	enteringChan:      make(chan *do.Entity),
	leavingChan:       make(chan *do.Entity),
	msgChan:           make(chan *do.Message),
}

func (w *ws) Start() {
	for {
		select {

		case e := <-w.enteringChan:
			w.entity[e.SN] = e
			break
		case e := <-w.leavingChan:
			delete(w.entity, e.SN)
			e.CloseMsgChannel()
			break
		case msg := <-w.msgChan:
			if msg == nil {
				break
			}
			if len(w.entity) > 0 {
				for _, e := range w.entity {
					if e.SN == msg.Entity.SN {
						continue
					}
					e.MsgChannel <- msg
				}
			}
			break
		case <-w.requestEntityChan:
			entitys := make([]*do.Entity, 0, len(w.entity))
			for _, e := range w.entity {
				entitys = append(entitys, e)
			}
			w.entityChan <- entitys
			break
		}
	}
}
func (w *ws) Entering(e *do.Entity) {
	w.enteringChan <- e
}
func (w *ws) Leaving(e *do.Entity) {
	w.leavingChan <- e
}
func (w *ws) Broadcast(msg *do.Message) {
	w.msgChan <- msg
}

func (w *ws) OnLine() []*do.Entity {
	w.requestEntityChan <- struct{}{}
	return <-w.entityChan
}
