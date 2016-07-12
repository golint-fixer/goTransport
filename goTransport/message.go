package goTransport

import (
	"log"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"encoding/json"
	"reflect"
)

type MessageType int

const (
	MessageTypeMethod MessageType = iota
	MessageTypeMethodResult
	MessageTypePub
)

type Message struct {
	Transport *Transport
	Type	MessageType `json:"type"`
	Id	int64 `json:"id"`
	Data	interface{} `json:"data"`
	Session sockjs.Session
	Json	[]byte
}

func (message *Message) Call() {
	log.Printf("Received new message %d with type %d", message.Id, message.Type)
	handlerA := GetHandler(message.Type)
	if handlerA == nil {
		log.Print("No parser found for message type: %d", message.Type)
		return
	}
	handlerB := reflect.ValueOf(handlerA.Get())
	handler := reflect.New(handlerB.Type()).Interface()

	log.Print(string(message.Json))
	err := json.Unmarshal(message.Json, &handler)
	if err != nil {
		log.Print(err)
		return
	}
	handler.(IHandler).Test123()
}

func (message *Message) Reply() {
	log.Print("repyl!")
	//message.Session.Send()
}