package messageType

import (
	"github.com/iain17/goTransport/transport/lib"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

type messageMethodResult struct {
	lib.Message
	Result       bool `json:"result"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	lib.Set(lib.MessageTypeMethodResult, messageMethodResult{})
}

func (message messageMethodResult) Validate(manager lib.MessageManager, session sockjs.Session) error {
	return nil
}

func (message messageMethodResult) Run(manager lib.MessageManager, session sockjs.Session) error {
	return 	nil
}