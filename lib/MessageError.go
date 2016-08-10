package lib

import (
	"log"
)

type messageError struct {
	Message
	Reason string `json:"reason"`
}

func init() {
	SetMessageDefinition(NewMessageError(nil))
}

func NewMessageError(reason error) *messageError {
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

func (message messageError) Received() error {
	log.Print(message.Reason)
	return nil
}
