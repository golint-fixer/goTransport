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
	Transport *Transport `json:"transport,omitempty"`
	Type	MessageType `json:"type"`
	Id	int64 `json:"id"`
	Data	interface{} `json:"data"`
	Session sockjs.Session `json:"session,omitempty"`
	Json	[]byte `json:"json,omitempty"`
}

type Reply struct {
	Success bool `json:"success"`
	Result interface{} `json:"result"`
}

func (message *Message) Call() {
	log.Printf("Received new message %d with type %d", message.Id, message.Type)
	handleDefinition := GetHandlerDefinition(message.Type)
	if handleDefinition == nil {
		log.Printf("No parser found for message type: %d", message.Type)
		return
	}
	handler := reflect.New(handleDefinition.Type).Interface()
	//log.Print(string(message.Json))
	err := json.Unmarshal(message.Json, &handler)
	if err != nil {
		log.Print("Error unmarshalling the message", err)
		message.Reply(Reply{
			Success: false,
			Result: err,
		})
		return
	}
	iHandler := handler.(IHandler)
	err = iHandler.Validate(message)
	if err != nil {
		log.Print("Error validating message", err)
		message.Reply(Reply{
			Success: false,
			Result: err,
		})
		return
	}

	err = iHandler.Run(message)
	if err != nil {
		log.Print("Error running handler", err)
		message.Reply(Reply{
			Success: false,
			Result: err,
		})
	}
	log.Print("Finished")
}

func (message *Message) Reply(reply interface{}) {
	handleDefinition := GetHandlerDefinition(message.Type)
	reply_message := Message{
		Type: handleDefinition.ReturnMessageType,
		Id: message.Id,
		Data: reply,
	}
	json_reply_message, err := json.Marshal(reply_message)
	if err != nil {
		log.Print("Error marshalling reply message", err)
		message.Reply(Reply{
			Success: false,
			Result: err,
		})
		return
	}
	message.Session.Send(string(json_reply_message))
}