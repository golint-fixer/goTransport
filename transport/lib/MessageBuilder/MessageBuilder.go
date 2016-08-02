package MessageBuilder

import (
	"log"
	"reflect"
	"encoding/json"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

func Build(definition reflect.Type, data string, manager interfaces.MessageManager, session *sockjs.Session) interfaces.IMessage {
	_message := reflect.New(definition).Interface()
	err := json.Unmarshal([]byte(data), &_message)
	if err != nil {
		log.Print("Error unmarshalling the message", err)
		return nil
	}
	__message := reflect.ValueOf(_message).Elem().Elem()
	if message, ok := __message.Addr().Interface().(interfaces.IMessage); ok {
		message.Initialize(manager, session)
		return message
	}
	log.Print("Could not cast message to IMessage interface. Invalid MessageType")
	return nil
}