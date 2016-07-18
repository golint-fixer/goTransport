package Message

import (
	"log"
	"github.com/iain17/goTransport/transport/Message"
)

type MessageMethod struct {
	Message.Message
	Name       string `json:"name"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	Message.Set(Message.MessageTypeMethod, MessageMethod{})
}

func (message MessageMethod) Validate() error {
	log.Print("yeh")
	log.Print(message.Name)
	return nil
}