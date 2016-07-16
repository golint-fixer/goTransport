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
			//var message Message
			//err = json.Unmarshal([]byte(msg), &message)
			//if err != nil {
			//	log.Print(err)
			//}
			//message.Transport = transport
			//message.Session = session
			//message.Json = []byte(msg)
			//message.Call()
			continue
		}
		break
	}
}