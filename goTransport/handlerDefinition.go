package goTransport

import (
	"sync"
	"reflect"
)

type IHandler interface {
	Validate(*Message) error
	Run(*Message) error
}

type HandlerDefinition struct {
	Type reflect.Type
	ReturnMessageType MessageType
}

var definitions map[MessageType]*HandlerDefinition
var definitions_mutex = new(sync.Mutex)

func initStorage() {
	definitions = make(map[MessageType]*HandlerDefinition)
}

func SetHandlerDefinition(receiveMessageType MessageType, returnMessageType MessageType, handler IHandler) {
	definitions_mutex.Lock()
	definitions[receiveMessageType] = &HandlerDefinition{
		Type: reflect.TypeOf(handler),
		ReturnMessageType:returnMessageType,
	}
	definitions_mutex.Unlock()
}

func GetHandlerDefinition(receiveMessageType MessageType) *HandlerDefinition {
	if a, ok := definitions[receiveMessageType]; ok {
		return a
	}
	return nil
}