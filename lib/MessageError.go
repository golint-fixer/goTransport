package lib

import (
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
)

type messageError struct {
	Message
	Reason string `json:"reason"`
}

func init() {
	SetMessageDefinition(newMessageError(nil))
}

func newMessageError(reason error) *messageError {
	message := &messageError{
		Message: NewMessage(MessageTypeError),
	}
	if reason != nil {
		message.Reason = reason.Error()
	}
	return message
}

func (message messageError) Sending() error {
	return nil
}

func (message messageError) Received(previousMessage interfaces.IMessage) error {
	log.Print(message.Reason)
	return nil
}
