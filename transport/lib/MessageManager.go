package lib

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"sync"
)

type MessageManager interface {
	Listen(session sockjs.Session)
	SetMethod(name string, method CallableMethod)
	GetMethod(name string) CallableMethod
	Send(message IMessage)
}

type messageManager struct {
	methods_mutex *sync.Mutex
	methods map[string]CallableMethod
}

func NewMessageManager() MessageManager {
	return &messageManager{
		methods_mutex: new(sync.Mutex),
		methods: make(map[string]CallableMethod),
	}
}

func (manager *messageManager) Listen(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			message := UnSerialize(msg)
			if message == nil {
				log.Print("Invalid message")
				continue
			}
			Start(message, manager, session)
			continue
		}
		break
	}
}

func (manager *messageManager) Send(message IMessage) {

}