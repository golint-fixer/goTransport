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
	"github.com/iain17/goTransport/transport/lib/MessageBuilder"
)

var current_id uint64
const headerDelimiter = "\f"

func init() {
	current_id = uint64(0)
}

type Message struct {
	Id uint64 `json:"id"`
	Type interfaces.MessageType `json:"type"`
	manager interfaces.MessageManager `json:"-"`
	session *sockjs.Session `json:"-"`
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

func UnSerialize(data string, manager interfaces.MessageManager, session *sockjs.Session) interfaces.IMessage {
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
	definition := MessageDefinition.Get(interfaces.MessageType(message_type), parts[1])
	if definition == nil {
		log.Print("No definition")
		return nil
	}
	return MessageBuilder.Build(definition, parts[1], manager, session)
}

func serialize(message interfaces.IMessage) string {
	json, err := json.Marshal(message)
	if err != nil {
		log.Print(err)
		return ""
	}
	return strconv.Itoa(int(message.GetType())) + headerDelimiter + string(json)
}

func (message *Message) Initialize(manager interfaces.MessageManager, session *sockjs.Session) {
	message.manager = manager
	message.session = session
}

func (message *Message) GetManager() interfaces.MessageManager {
	return message.manager
}

func (message *Message) GetSession() *sockjs.Session {
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

func (message *Message) Validate() error {
	return errors.New(fmt.Sprint("MessageType %d has not implemented Validate()", message.Type))
}

func (message *Message) Run() error {
	return errors.New(fmt.Sprint("MessageType %d has not implemented run()", message.Type))
}

func (message *Message) Reply(replyMessage interfaces.IMessage) {
	replyMessage.SetId(message.GetId())
	message.manager.Send(serialize(replyMessage), message.session)
}

func (message *Message) Send() {
	current_id++
	message.SetId(current_id)
	message.manager.Send(serialize(message), message.session)
}

func Start(message interfaces.IMessage) error {
	if(message.GetManager() == nil || message.GetSession() == nil) {
		return errors.New(fmt.Sprint("MessageType %d has not been initialized.", message.GetType()))
	}

	if err := message.Validate(); err != nil {
		return err
	}
	if err := message.Run(); err != nil {
		return err
	}
	return nil
}