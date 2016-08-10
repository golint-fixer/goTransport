package lib

import (
	"errors"
	"reflect"
)

type messageMethod struct {
	Message
	Name       string `json:"name"`
	Parameters []interface{} `json:"parameters"`
}


func init() {
	SetMessageDefinition(NewMessageMethod("", nil))
}

func NewMessageMethod(name string, parameters []interface{}) *messageMethod {
	return &messageMethod{
		Message: NewMessage(MessageTypeMethod),
		Name: name,
		Parameters: parameters,
	}
}

func (message *messageMethod) Sending() error {
	return nil
}

//Received a request to call a method on our side. Figure out which method it is, and dynamically call it with the sent along parameters.
func (message *messageMethod) Received() error {
	//Catch any panics
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if (recover() != nil) {
			message.Reply(NewMessageError(errors.New("Panic whilst running method")))
		}
	}()

	//Prepend the session to the parameters
	message.Parameters = append([]interface{}{message.GetSession()}, message.Parameters...)

	//Get the requested method
	method := message.GetSession().GetClient().GetMethod(message.Name)
	if method == nil  {
		return errors.New("[404]: Unknown method:"+message.Name)
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
	message.Reply(NewMessageMethodResult(true, result))
	return 	nil
}