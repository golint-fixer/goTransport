package goTransport
import (
	//"encoding/json"
	//"errors"
)
import (
	"log"
)

type HandleMethod struct {
	Message *Message
	Data struct {
		     Name       string `json:"name"`
		     Parameters []interface{} `json:"parameters"`
	     } `json:"data"`
}

func init() {
	SetParser(MessageTypeMethod, MessageTypeMethodResult, HandleMethod{})
}

func (m HandleMethod) getMethod() RPCMethod {
	return m.Message.Transport.getRPCMethod(m.Data.Name)
}

func (m HandleMethod) Call() {
	log.Print("Method called:", m.Data.Name)
	rpcMethod := m.getMethod()
	rpcMethod(m.Data.Parameters)
}