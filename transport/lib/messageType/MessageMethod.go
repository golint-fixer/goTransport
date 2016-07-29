package messageType

import (
	"log"
	"github.com/iain17/goTransport/transport/lib"
	"errors"
	"reflect"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

type messageMethod struct {
	lib.Message
	Name       string `json:"name"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	lib.Set(lib.MessageTypeMethod, messageMethod{})
}

func (message messageMethod) Validate(manager lib.MessageManager, session sockjs.Session) error {
	log.Print(message.Name)
	if manager.GetMethod(message.Name) == nil  {
		return errors.New("[404]: Unknown method:"+message.Name)
	}
	return nil
}

func (message messageMethod) Run(manager lib.MessageManager, session sockjs.Session) error {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if (recover() != nil) {
			manager.Send(newMessageError(errors.New("Panic whilst running method")))
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
	manager.Send(messageMethodResult{
		Result: true,
		Parameters: result,
	})
	return 	nil
}