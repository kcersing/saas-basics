package do

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"sync"
)

// Entity 实体
type Entity struct {
	SN         string        `json:"sn"`
	Ip         string        `json:"ip"`
	MsgChannel chan *Message `json:"-"`
	lock       sync.RWMutex
	conn       *websocket.Conn
}

// Message 消息
type Message struct {
	Entity *Entity `json:"entity"`
	//Service        string  `json:"service"`
	Content        string `json:"content"`
	Face           *Face  `json:"face"`
	Time           string `json:"time"`
	ClientSentTime string `json:"client_sent_time"`
}

// Face 人脸信息
type Face struct {
	// 会员id
	MemberId string `json:"member_id"`
	// 员工id
	UserId string `json:"user_id"`

	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Sex    string `json:"sex"`
	Level  string `json:"level"`

	FacePic string `json:"face_pic"`
	//人脸特征值
	FaceEigenvalue string `json:"face_eigenvalue"`
	//人脸更新时间
	FacePicUpdatedTime string `json:"face_pic_updated_time"`
}

// SendMsg 发送信息
func (e *Entity) SendMsg() {
	e.lock.RLock()
	defer e.lock.RUnlock()
	for msg := range e.MsgChannel {
		err := e.conn.WriteJSON(msg)
		if err != nil {
			hlog.Error("write:", err)
			break
		}
	}
}

func (e *Entity) ReceiveMsg() {

}

func (e *Entity) CloseMsgChannel() {
	e.lock.Lock()
	defer e.lock.RUnlock()
	close(e.MsgChannel)
}
