package Session

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/Message"
	"github.com/iain17/goTransport/transport/lib/messageType"
	"sync"
)

type session struct {
	socket sockjs.Session
	client interfaces.Client
	currentId uint64
	currentId_mutex *sync.Mutex
}

func NewSession(socket sockjs.Session, client interfaces.Client) interfaces.Session {
	return &session{
		socket: socket,
		client: client,
		currentId: 0,
		currentId_mutex: new(sync.Mutex),
	}
}

func (session *session) GetClient() interfaces.Client {
	return session.client
}

func (session *session) GetCurrentId() uint64 {
	return session.currentId
}

func (session *session) SetCurrentId(id uint64) {
	session.currentId_mutex.Lock()
	session.currentId = id
	session.currentId_mutex.Unlock()
}

func (session *session) IncrementCurrentId() {
	session.currentId_mutex.Lock()
	session.currentId++
	session.currentId_mutex.Unlock()
}

func (session *session) Messaged(data string) {
	log.Printf("Received: %s", data)
	message := Message.UnSerialize(data)
	if message == nil {
		log.Print("Invalid message")
		return
	}
	message.Initialize(session)

	err := message.Received()
	if err != nil {
		message.Reply(messageType.NewMessageError(err))
	}
}

func (session *session) Send(message string) {
	session.socket.Send(message)
	log.Print("Sending to client:", message)
}

func (session *session) Call(name string, parameters []interface{}) {
	message := messageType.NewMessageMethod(name, parameters)
	message.Initialize(session)
	Message.Send(message)
}