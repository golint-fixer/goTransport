package Message

import (
	"sync"
	"reflect"
)

var definitions map[MessageType]reflect.Type
var definitions_mutex = new(sync.Mutex)

func init() {
	definitions = make(map[MessageType]reflect.Type)
}

func Set(messageType MessageType, definition IMessage) {
	definitions_mutex.Lock()
	definitions[messageType] = reflect.TypeOf(definition)
	definitions_mutex.Unlock()
}

func Get(messageType MessageType, data string) IMessage {
	definition := definitions[messageType]
	if definition == nil {
		return nil
	}
	return build(definition, data)
}