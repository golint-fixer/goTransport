package lib

import (
	"errors"
	"github.com/iain17/goTransport/lib/interfaces"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"sync"
)

type session struct {
	socket         sockjs.Session
	client         interfaces.IClient
	currentId      uint64
	currentIDMutex *sync.Mutex
	messages       []interfaces.IMessage
}

//Creates a new session.
func NewSession(socket sockjs.Session, client interfaces.IClient) interfaces.ISession {
	return &session{
		socket:         socket,
		client:         client,
		currentId:      0,
		currentIDMutex: new(sync.Mutex),
	}

}

func (session *session) GetClient() interfaces.IClient {
	return session.client
}

func (session *session) GetCurrentId() uint64 {
	return session.currentId
}

func (session *session) SetCurrentId(id uint64) {
	session.currentIDMutex.Lock()
	session.currentId = id
	session.currentIDMutex.Unlock()
}

func (session *session) IncrementCurrentId() {
	session.currentIDMutex.Lock()
	session.currentId++
	session.currentIDMutex.Unlock()
}

func (session *session) Messaged(data string) error {
	log.Printf("Received: %s", data)
	message := UnSerialize(data)
	if message == nil {
		log.Print("Invalid message received.")
		return errors.New("Invalid message received.")
	}
	message.Initialize(session)

	//Set the previous message that was sent for this message id.
	err := message.Received(session.GetPreviousMessage(message))
	if err != nil {
		message.Reply(newMessageError(err))
		return err
	}
	return nil
}

func (session *session) SetPreviousMessage(message interfaces.IMessage) {
	session.messages[message.GetId()] = message
}

func (session *session) GetPreviousMessage(message interfaces.IMessage) interfaces.IMessage {
	return session.messages[message.GetId()]
}

func (session *session) Send(message string) {
	session.socket.Send(message)
	log.Print("Sending to client:", message)
}

func (session *session) Call(name string, parameters []interface{}) {
	message := newMessageMethod(name, parameters)
	message.Initialize(session)
	Send(message)
}
