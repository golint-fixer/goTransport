package lib

import (
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
)

type messageMethodResult struct {
	Message
	Result     bool          `json:"result"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	SetMessageDefinition(newMessageMethodResult(false, nil))
}

func newMessageMethodResult(result bool, parameters []interface{}) *messageMethodResult {
	return &messageMethodResult{
		Message:    NewMessage(MessageTypeMethodResult),
		Result:     result,
		Parameters: parameters,
	}
}

func (message messageMethodResult) Sending() error {
	return nil
}

func (message messageMethodResult) Received(previousMessage interfaces.IMessage) error {
	log.Print(previousMessage)
	log.Print(message.Parameters)
	return nil
}
