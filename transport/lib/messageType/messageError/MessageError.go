package messageType

import (
	"github.com/iain17/goTransport/transport/lib"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"log"
)

type messageError struct {
	lib.Message
	Reason error `json:"reason"`
}

func init() {
	message := New(nil)
	lib.Set(message.Type, message)
}

func New(reason error) *messageError {
	return &messageError{
		Message: lib.NewMessage(lib.MessageTypeError),
		Reason: reason,
	}
}

func (message messageError) Validate(manager lib.MessageManager, session sockjs.Session) error {
	return nil
}

func (message messageError) Run(manager lib.MessageManager, session sockjs.Session) error {
	log.Print(message.Reason)
	return 	nil
}