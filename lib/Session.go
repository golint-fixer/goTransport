package lib

import (
	"errors"
	"github.com/iain17/goTransport/lib/interfaces"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"sync"
)

type session struct {
	socket          sockjs.Session
	client          interfaces.Client
	currentId       uint64
	currentId_mutex *sync.Mutex
}

func NewSession(socket sockjs.Session, client interfaces.Client) interfaces.Session {
	return &session{
		socket:          socket,
		client:          client,
		currentId:       0,
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

func (session *session) Messaged(data string) error {
	log.Printf("Received: %s", data)
	message := UnSerialize(data)
	if message == nil {
		log.Print("Invalid message")
		return errors.New("Invalid message")
	}
	message.Initialize(session)

	err := message.Received()
	if err != nil {
		message.Reply(NewMessageError(err))
		return err
	}
	return nil
}

func (session *session) Send(message string) {
	session.socket.Send(message)
	log.Print("Sending to client:", message)
}

func (session *session) Call(name string, parameters []interface{}) {
	message := NewMessageMethod(name, parameters)
	message.Initialize(session)
	Send(message)
}
