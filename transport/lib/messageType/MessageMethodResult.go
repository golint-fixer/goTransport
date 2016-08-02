package messageType

import (
	"github.com/iain17/goTransport/transport/lib/MessageDefinition"
	"github.com/iain17/goTransport/transport/lib/Message"
)

type messageMethodResult struct {
	Message.Message
	Result       bool `json:"result"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	MessageDefinition.Set(NewMessageMethodResult(false, nil))
}

func NewMessageMethodResult(result bool, parameters []interface{}) *messageMethodResult {
	return &messageMethodResult{
		Message: Message.NewMessage(MessageDefinition.MessageTypeMethodResult),
		Result: result,
		Parameters: parameters,
	}
}

func (message messageMethodResult) Validate() error {
	return nil
}

func (message messageMethodResult) Run() error {
	return 	nil
}