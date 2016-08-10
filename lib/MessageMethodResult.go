package lib

import (
	"errors"
	"github.com/iain17/goTransport/lib/interfaces"
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
	if previousMessageMethod, ok := previousMessage.(*messageMethod); ok {
		promise := previousMessageMethod.GetPromise()
		if message.Result {
			promise.Resolve(message.Parameters[0])
		} else {
			promise.Reject(errors.New("Result false"))
		}
		return nil
	} else {
		return errors.New("Invalid or no previousMessage")
	}
}
