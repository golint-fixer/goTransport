package Message

import "log"

type MessageMethod struct {
	Message
	Name       string `json:"name"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	Set(MessageTypeMethod, MessageMethod{})
}

func (message MessageMethod) Validate() error {
	log.Print(message.Name)
	return nil
}