package Message

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
)

type MessageManager interface {
	Listen(session sockjs.Session)
}

type messageManager struct {

}

func NewMessageManager() MessageManager {
	return &messageManager{}
}

func (messageManager *messageManager) Listen(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			message := UnSerialize(msg)
			if message == nil {
				log.Print("Invalid message")
				continue
			}
			message.Start()
			continue
		}
		break
	}
}