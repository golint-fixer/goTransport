package messageType

import (
	"log"
	"errors"
	"reflect"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"github.com/iain17/goTransport/transport/lib/interfaces"
	"github.com/iain17/goTransport/transport/lib/Message"
	"github.com/iain17/goTransport/transport/lib/MessageDefinition"
)

type messageMethod struct {
	Message.Message
	Name       string `json:"name"`
	Parameters []interface{} `json:"parameters"`
}


func init() {
	MessageDefinition.Set(NewMessageMethod("", nil))
}

func NewMessageMethod(name string, parameters []interface{}) *messageMethod {
	return &messageMethod{
		Message: Message.NewMessage(MessageDefinition.MessageTypeMethod),
		Name: name,
		Parameters: parameters,
	}
}

func (message *messageMethod) Validate(manager interfaces.MessageManager, session sockjs.Session) error {
	log.Print(message.Name)
	if manager.GetMethod(message.Name) == nil  {
		return errors.New("[404]: Unknown method:"+message.Name)
	}
	return nil
}

func (message *messageMethod) Run(manager interfaces.MessageManager, session sockjs.Session) error {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if (recover() != nil) {
			message.Reply(NewMessageError(errors.New("Panic whilst running method")), manager, session)
		}
	}()

	rpcMethod := reflect.ValueOf(manager.GetMethod(message.Name))
	if len(message.Parameters) != rpcMethod.Type().NumIn() {
		return errors.New("The number of parameters sent do not match the amount required.")
	}
	in := make([]reflect.Value, len(message.Parameters))
	for k, param := range message.Parameters {
		in[k] = reflect.ValueOf(param)
	}

	values := rpcMethod.Call(in)
	var result []interface{}
	for _, value := range values {
		if value.IsValid() {
			result = append(result, value.Interface())
		}
	}
	message.Reply(NewMessageMethodResult(true, result), manager, session)
	return 	nil
}