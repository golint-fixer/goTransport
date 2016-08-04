package messageType

import (
	"log"
	"errors"
	"reflect"
	"github.com/iain17/goTransport/transport/lib/Message"
)

type messageMethod struct {
	Message.Message
	Name       string `json:"name"`
	Parameters []interface{} `json:"parameters"`
}


func init() {
	Message.Set(NewMessageMethod("", nil))
}

func NewMessageMethod(name string, parameters []interface{}) *messageMethod {
	return &messageMethod{
		Message: Message.NewMessage(Message.MessageTypeMethod),
		Name: name,
		Parameters: parameters,
	}
}

func (message *messageMethod) Validate() error {
	log.Print(message.Name)

	if message.GetManager().GetMethod(message.Name) == nil  {
		return errors.New("[404]: Unknown method:"+message.Name)
	}
	return nil
}

func (message *messageMethod) Run() error {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if (recover() != nil) {
			message.Reply(NewMessageError(errors.New("Panic whilst running method")))
		}
	}()

	rpcMethod := reflect.ValueOf(message.GetManager().GetMethod(message.Name))
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
	message.Reply(NewMessageMethodResult(true, result))
	return 	nil
}