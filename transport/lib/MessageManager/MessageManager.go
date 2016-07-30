package MessageManager

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"sync"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/Message"
	"github.com/iain17/goTransport/transport/lib/messageType"
)

type messageManager struct {
	methods_mutex *sync.Mutex
	methods map[string]interfaces.CallableMethod
}

func NewMessageManager() interfaces.MessageManager {
	return &messageManager{
		methods_mutex: new(sync.Mutex),
		methods: make(map[string]interfaces.CallableMethod),
	}
}

func (manager *messageManager) Listen(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			message := Message.UnSerialize(msg)
			if message == nil {
				log.Print("Invalid message")
				continue
			}
			error := Message.Start(message, manager, session)
			if error != nil {
				message.Reply(messageType.NewMessageError(error), manager, session)
			}
			continue
		}
		break
	}
}

func (manager *messageManager) Send(message string, session sockjs.Session) {
	session.Send(message)
	//log.Print("Send:", message)
}