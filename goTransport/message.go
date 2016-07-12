package goTransport

import (
	"log"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
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

func (message *Message) Handle() {
	log.Printf("Received new message %d with type %d", message.Id, message.Type)
	parser := GetParser(message.Type)
	if parser == nil {
		log.Print("No parser found for message type: %d", message.Type)
		return
	}
	handler, err := parser.Parse(message, parser.ReturnMessageType)
	if err != nil {
		log.Print(err)
		return
	}

	handler.Call()
}

func (message *Message) Reply() {
	log.Print("repyl!")
	//message.Session.Send()
}