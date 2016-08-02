package MessageDefinition

import (
	"sync"
	"reflect"
	"github.com/iain17/goTransport/transport/lib/interfaces"
)

const (
	MessageTypeMethod interfaces.MessageType = iota
	MessageTypeMethodResult
	MessageTypeError
	MessageTypePub
)

var definitions map[interfaces.MessageType]reflect.Type
var definitions_mutex = new(sync.Mutex)

func init() {
	definitions = make(map[interfaces.MessageType]reflect.Type)
}

func Set(definition interfaces.IMessage) {
	definitions_mutex.Lock()
	definitions[definition.GetType()] = reflect.TypeOf(definition)
	definitions_mutex.Unlock()
}

func Get(messageType interfaces.MessageType, data string) reflect.Type {
	definition := definitions[messageType]
	if definition == nil {
		return nil
	}
	return definition
}