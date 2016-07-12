package goTransport
import (
	//"encoding/json"
	//"errors"
)
import "log"

type HandleMethod struct {
	Message *Message
	Data struct {
		     Name       string `json:"name"`
		     Parameters []interface{} `json:"parameters"`
	     } `json:"data"`
}

func init() {
	SetParser(MessageTypeMethod, &Parser{
		Get: get,
		ReturnMessageType:MessageTypeMethodResult,
	})
}

func get() IHandler {
	var handler HandleMethod
	return handler
}

func (m HandleMethod) getMethod() RPCMethod {
	return m.Message.Transport.getRPCMethod(m.Data.Name)
}

func (m HandleMethod) Test123() {
	log.Print("jaja", m.Data.Name)
	//rpcMethod := m.getMethod()
	//rpcMethod(m.Data.Parameters)
}