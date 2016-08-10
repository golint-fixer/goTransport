package lib

import (
	"encoding/json"
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
	"reflect"
)

func Build(definition reflect.Type, data string) interfaces.IMessage {
	_message := reflect.New(definition).Interface()
	err := json.Unmarshal([]byte(data), &_message)
	if err != nil {
		log.Print("Error unmarshalling the message", err)
		return nil
	}
	__message := reflect.ValueOf(_message).Elem().Elem()
	if message, ok := __message.Addr().Interface().(interfaces.IMessage); ok {
		return message
	}

	log.Print("Could not cast message to IMessage interface. Invalid MessageType")
	return nil
}
