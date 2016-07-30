package messageType

import (
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
	"github.com/iain17/goTransport/transport/lib/Message"
	"github.com/iain17/goTransport/transport/lib/MessageDefinition"
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

type messageError struct {
	Message.Message
	Reason error `json:"reason"`
}

func init() {
	MessageDefinition.Set(NewMessageError(nil))
}

func NewMessageError(reason error) *messageError {
	return &messageError{
		Message: Message.NewMessage(MessageDefinition.MessageTypeError),
		Reason: reason,
	}
}

func (message messageError) Validate(manager interfaces.MessageManager, session sockjs.Session) error {
	return nil
}

func (message messageError) Run(manager interfaces.MessageManager, session sockjs.Session) error {
	log.Print(message.Reason)
	return 	nil
}