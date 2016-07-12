package goTransport

import (
	"sync"
	"reflect"
)

//Parser
type IHandler interface {
	Call()
}

type Parser struct {
	//Validate the message. and returns a handler
	Type reflect.Type
	ReturnMessageType MessageType
}

var parsers map[MessageType]*Parser
var parsers_mutex = new(sync.Mutex)

func initStorage() {
	parsers = make(map[MessageType]*Parser)
}

func SetParser(messageType MessageType, returnMessageType MessageType, handler IHandler) {
	parsers_mutex.Lock()
	parsers[messageType] = &Parser{
		Type: reflect.TypeOf(handler),
		ReturnMessageType:returnMessageType,
	}
	parsers_mutex.Unlock()
}

func GetParser(messageType MessageType) *Parser {
	if a, ok := parsers[messageType]; ok {
		return a
	}

	return nil
}