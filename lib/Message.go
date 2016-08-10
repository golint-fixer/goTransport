package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iain17/goTransport/lib/interfaces"
	"log"
	"strconv"
	"strings"
)

const headerDelimiter = "\f"

type Message struct {
	Id      uint64                 `json:"id"`
	Type    interfaces.MessageType `json:"type"`
	session interfaces.ISession    `json:"-"`
}

//Constructor
func NewMessage(message_type interfaces.MessageType) Message {
	//log.Printf("NewMessage called: %d", current_id)
	return Message{
		Type: message_type,
	}
}

//func validate() interfaces.IMessage {return &Message{}}

func (message *Message) Initialize(session interfaces.ISession) {
	message.session = session

	if message.GetId() > session.GetCurrentId() {
		session.SetCurrentId(message.GetId())
	}
}

//getters and setters
func (message *Message) GetSession() interfaces.ISession {
	return message.session
}

func (message *Message) GetType() interfaces.MessageType {
	return message.Type
}

func (message *Message) GetId() uint64 {
	return message.Id
}

func (message *Message) SetId(id uint64) {
	message.Id = id
}

//Sending the message.
func (message *Message) Sending() error {
	return errors.New(fmt.Sprintf("MessageType %d has not implemented Sending()", message.Type))
}

//Received the message.
func (message *Message) Received(previousMessage interfaces.IMessage) error {
	return errors.New(fmt.Sprintf("MessageType %d has not implemented Received()", message.Type))
}

//This function converts a message object to the respectable message structure. header and json.
func serialize(message interfaces.IMessage) string {
	json, err := json.Marshal(message)
	if err != nil {
		log.Print(err)
		return ""
	}

	return strconv.Itoa(int(message.GetType())) + headerDelimiter + string(json)
}

//This function converts a string with the correct message structure from message header and json to a message object.
func UnSerialize(data string) interfaces.IMessage {
	parts := strings.Split(data, headerDelimiter)
	if len(parts) != 2 {
		log.Print("Invalid length:", len(parts))
		return nil
	}

	message_type, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Print(err)
		return nil
	}

	definition := GetMessageDefinition(interfaces.MessageType(message_type))
	if definition == nil {
		log.Printf("No definition for type: %d", interfaces.MessageType(message_type))
		return nil
	}

	return build(definition, parts[1])
}

//Reply to this message.
func (message *Message) Reply(replyMessage interfaces.IMessage) {
	if message.GetSession() == nil {
		log.Printf("MessageType %d has not been initialized.", message.GetType())
		return
	}

	replyMessage.SetId(message.GetId())
	replyMessage.Sending()
	message.session.SetPreviousMessage(message)
	message.session.Send(serialize(replyMessage))
}

//Sends the message to the client.
func Send(message interfaces.IMessage) {
	if message.GetSession() == nil {
		log.Printf("MessageType %d has not been initialized.", message.GetType())
		return
	}

	session := message.GetSession()
	session.IncrementCurrentId()
	message.SetId(session.GetCurrentId())
	message.Sending()
	session.SetPreviousMessage(message)
	session.Send(serialize(message))
}
