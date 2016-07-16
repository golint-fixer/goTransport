package Message

import (
	"log"
	"strings"
	"strconv"
)

type MessageType int

const (
	MessageTypeMethod MessageType = iota
	MessageTypeMethodResult
	MessageTypePub
)

type IMessage interface {
	//GetType() MessageType
	//setReply()
	Validate() error
	Run() (interface{}, error)
	Start() bool
	//serialize() string
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

func (message Message) Encode() interface{} {
	return nil
}

func (message Message) GetType() MessageType {
	return message.Type
}

func (message Message) Validate() error {
	return nil
}

func (message Message) Run() (interface{}, error) {
	return nil, nil
}

func (message Message) Start() bool {
	log.Print("Start message")
	return false
}