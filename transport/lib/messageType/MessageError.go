package messageType

import (
	"log"
	"github.com/iain17/goTransport/transport/lib/Message"
	"github.com/iain17/goTransport/transport/lib/MessageDefinition"
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

func (message messageError) Validate() error {
	return nil
}

func (message messageError) Run() error {
	log.Print(message.Reason)
	return 	nil
}