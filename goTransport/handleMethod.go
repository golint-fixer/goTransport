package goTransport
import (
	"encoding/json"
	"errors"
)

type handleMethod struct {
	Message *Message
	Data struct {
		     Name       string `json:"name"`
		     Parameters []interface{} `json:"parameters"`
	     } `json:"data"`
}

func init() {
	SetParser(MessageTypeMethod, &Parser{
		Parse: parse,
		ReturnMessageType:MessageTypeMethodResult,
	})
}

func parse(message *Message, returnMessageType MessageType) (Handler, error) {
	var handler handleMethod
	err := json.Unmarshal(message.Json, &handler)
	if err != nil {
		return nil, err
	}
	handler.Message = message
	if handler.getMethod() == nil {
		return nil, errors.New("Method doesn't exist")
	}
	return handler, nil
}

func (m handleMethod) getMethod() RPCMethod {
	return m.Message.Transport.getRPCMethod(m.Data.Name)
}

func (m handleMethod) Call() {
	rpcMethod := m.getMethod()
	rpcMethod(m.Data.Parameters)
}