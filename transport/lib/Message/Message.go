package Message

import (
	"log"
	"strings"
	"strconv"
	"errors"
	"fmt"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/MessageDefinition"
	"encoding/json"
)

var current_id uint64
const headerDelimiter = "\f"

func init() {
	current_id = uint64(0)
}

type Message struct {
	Id uint64 `json:"id"`
	Type interfaces.MessageType `json:"type"`
}

func NewMessage(message_type interfaces.MessageType) Message {
	log.Printf("NewMessage called: %d", current_id)
	return Message{
		Type: message_type,
	}
}

func validate() interfaces.IMessage {
	return &Message{}
}

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
	return MessageDefinition.Get(interfaces.MessageType(message_type), parts[1])
}

func serialize(message interfaces.IMessage) string {
	json, err := json.Marshal(message)
	if err != nil {
		log.Print(err)
		return ""
	}
	return strconv.Itoa(int(message.GetType())) + headerDelimiter + string(json)
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

func (message *Message) Validate(manager interfaces.MessageManager, session sockjs.Session) error {
	return errors.New(fmt.Sprint("MessageType %d has not implemented Validate()", message.Type))
}

func (message *Message) Run(manager interfaces.MessageManager, session sockjs.Session) error {
	return errors.New(fmt.Sprint("MessageType %d has not implemented run()", message.Type))
}

func (message *Message) Reply(replyMessage interfaces.IMessage, manager interfaces.MessageManager, session sockjs.Session) {
	replyMessage.SetId(message.GetId())
	manager.Send(serialize(replyMessage), session)
}

func (message *Message) Send(manager interfaces.MessageManager, session sockjs.Session) {
	current_id++
	message.SetId(current_id)
	manager.Send(serialize(message), session)
}

func Start(message interfaces.IMessage, manager interfaces.MessageManager, session sockjs.Session) error {
	if err := message.Validate(manager, session); err != nil {
		return err
	}
	if err := message.Run(manager, session); err != nil {
		return err
	}
	return nil
}