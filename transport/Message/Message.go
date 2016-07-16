package Message

import (
	"log"
	"strings"
	"strconv"
	"errors"
)

type MessageType int

const (
	MessageTypeMethod MessageType = iota
	MessageTypeMethodResult
	MessageTypePub
)

type IMessage interface {
	GetType() MessageType
	//setReply()
	Validate() error
	Run() error
	Start() bool
	serialize() string
}

var current_id uint64
const headerDelimiter = "\f"

func init() {
	current_id = uint64(0)
}

type Message struct {
	Id uint64
	Type MessageType
}

func NewMessage(message_type MessageType) IMessage {
	current_id++
	return &Message{
		Id: current_id,
		Type: message_type,
	}
}

func UnSerialize(data string) IMessage {
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
	return Get(MessageType(message_type), parts[1])
}

func (message Message) serialize() string {
	log.Print("MessageType has not implemented Encode()")
	return ""
}

func (message Message) GetType() MessageType {
	return message.Type
}

func (message Message) Validate() error {
	return errors.New("MessageType has not implemented Validate()")
}

func (message Message) Run() error {
	return errors.New("MessageType has not implemented Run()")
}

func (message Message) Start() bool {
	if err := message.Validate(); err != nil {
		log.Print(err)
		return false
	}
	if err := message.Run(); err != nil {
		log.Print(err)
		return false
	}
	return true
}