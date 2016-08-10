package lib

import (
	"sync"
	"reflect"
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

const (
	MessageTypeTest interfaces.MessageType = iota
	MessageTypeMethod
	MessageTypeMethodResult
	MessageTypeError
)

var definitions map[interfaces.MessageType]reflect.Type
var definitions_mutex = new(sync.Mutex)

func init() {
	definitions = make(map[interfaces.MessageType]reflect.Type)
}

func SetMessageDefinition(definition interfaces.IMessage) {
	definitions_mutex.Lock()
	definitions[definition.GetType()] = reflect.TypeOf(definition)
	definitions_mutex.Unlock()
}

func GetMessageDefinition(messageType interfaces.MessageType) reflect.Type {
	definition := definitions[messageType]
	if definition == nil {
		return nil
	}
	return definition
}