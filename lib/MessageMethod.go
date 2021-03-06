package lib

import (
	"errors"
	"github.com/fanliao/go-promise"
	"github.com/iain17/goTransport/lib/interfaces"
	"reflect"
)

type messageMethod struct {
	Message
	Name       string        `json:"name"`
	Parameters []interface{} `json:"parameters"`
	promise    *promise.Promise
}

func init() {
	SetMessageDefinition(newMessageMethod("", nil))
}

func newMessageMethod(name string, parameters []interface{}) *messageMethod {
	return &messageMethod{
		Message:    NewMessage(MessageTypeMethod),
		Name:       name,
		Parameters: parameters,
	}
}

func (message *messageMethod) Sending() error {
	message.promise = promise.NewPromise()
	return nil
}

func (message *messageMethod) GetPromise() *promise.Promise {
	return message.promise
}

//Received a request to call a method on our side. Figure out which method it is, and dynamically call it with the sent along parameters.
func (message *messageMethod) Received(previousMessage interfaces.IMessage) error {
	//Catch any panics
	defer func() {
		// recover from panic if one occurred. Set err to nil otherwise.
		if recover() != nil {
			message.Reply(newMessageError(errors.New("Panic whilst running method")))
		}
	}()

	//Prepend the session to the parameters
	message.Parameters = append([]interface{}{message.GetSession()}, message.Parameters...)

	//Get the requested method
	method := message.GetSession().GetClient().GetMethod(message.Name)
	if method == nil {
		return errors.New("[404]: Unknown method:" + message.Name)
	}
	rpcMethod := reflect.ValueOf(method)

	//Setup the sent parameters
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

	//Reply with the returned result by the called method
	message.Reply(newMessageMethodResult(true, result))
	return nil
}
