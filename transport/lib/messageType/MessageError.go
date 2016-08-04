package messageType

import (
	"log"
	"github.com/iain17/goTransport/transport/lib/Message"
)

type messageError struct {
	Message.Message
	Reason error `json:"reason"`
}

func init() {
	Message.Set(NewMessageError(nil))
}

func NewMessageError(reason error) *messageError {
	return &messageError{
		Message: Message.NewMessage(Message.MessageTypeError),
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