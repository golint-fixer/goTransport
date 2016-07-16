package Message

import "log"

type MessageMethod struct {

}

func init() {
	Set(MessageTypeMethod, MessageMethod{})
}


func (message MessageMethod) Validate() error {
	return nil
}

func (message MessageMethod) Run() (interface{}, error) {
	return nil, nil
}

func (message MessageMethod) Start() bool {
	log.Print("Start message")
	return false
}