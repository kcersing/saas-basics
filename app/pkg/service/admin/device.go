package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"log"
)

var upgrader = websocket.HertzUpgrader{} // use default options

func Create() {

}
func Faces(_ context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
}
