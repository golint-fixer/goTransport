package lib

import (
	"log"
	"strings"
	"strconv"
	"errors"
	"fmt"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

type IMessage interface {
	GetType() MessageType
	//setReply()
	Validate(manager MessageManager, session sockjs.Session) error
	Run(manager MessageManager, session sockjs.Session) error
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

func NewMessage(message_type MessageType) Message {
	current_id++
	return Message{
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

func (message Message) Validate(manager MessageManager, session sockjs.Session) error {
	return errors.New(fmt.Sprint("MessageType %d has not implemented Validate()", message.Type))
}

func (message Message) Run(manager MessageManager, session sockjs.Session) error {
	return errors.New(fmt.Sprint("MessageType %d has not implemented run()", message.Type))
}

func Start(message IMessage, manager MessageManager, session sockjs.Session) bool {
	if err := message.Validate(manager, session); err != nil {
		log.Print(err)
		//manager.Send(newMessageError(err))
		return false
	}
	if err := message.Run(manager, session); err != nil {
		log.Print(err)
		//manager.Send(newMessageError(err))
		return false
	}
	return true
}