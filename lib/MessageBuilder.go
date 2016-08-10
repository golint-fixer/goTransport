package lib

import (
	"encoding/json"
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
	"reflect"
)

func build(definition reflect.Type, jsonData string) interfaces.IMessage {
	messageInterface := reflect.New(definition).Interface()
	err := json.Unmarshal([]byte(jsonData), &messageInterface)
	if err != nil {
		log.Print("Error unmarshalling the message", err)
		return nil
	}
	if message, ok := reflect.ValueOf(messageInterface).Elem().Interface().(interfaces.IMessage); ok {
		return message
	}

	log.Print("Could not cast message to IMessage interface. Invalid MessageType")
	return nil
}
