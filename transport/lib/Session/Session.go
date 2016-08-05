package Session

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/Message"
	"github.com/iain17/goTransport/transport/lib/messageType"
)

type session struct {
	socket sockjs.Session
	client interfaces.Client
}

func NewSession(socket sockjs.Session, client interfaces.Client) interfaces.Session {
	return &session{
		socket: socket,
		client: client,
	}
}

func (session *session) GetClient() interfaces.Client {
	return session.client
}

func (session *session) Messaged(data string) {
	message := Message.UnSerialize(data)
	if message == nil {
		log.Print("Invalid message")
		return
	}
	message.Initialize(session)

	error := Message.Start(message)
	if error != nil {
		message.Reply(messageType.NewMessageError(error))
	}
}

func (session *session) Send(message string) {
	session.socket.Send(message)
	//log.Print("Send:", message)
}