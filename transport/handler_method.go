package transport
import (
	"errors"
	"reflect"
	"log"
)

type HandleMethod struct {
	test int
	Data struct {
		     Name       string `json:"name"`
		     Parameters []interface{} `json:"parameters"`
	     } `json:"data"`
}

func init() {
	SetHandlerDefinition(MessageTypeMethod, MessageTypeMethodResult, HandleMethod{})
}

func (m HandleMethod) Validate(message *Message) error {
	////Run some checks.
	if message.Transport.getRPCMethod(m.Data.Name) == nil  {
		return errors.New("[404]: Invalid method.")
	}
	return nil
}

func (m HandleMethod) Run(message *Message) error {
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if (recover() != nil) {
			message.Reply(Reply{
				Success: false,
				Result: errors.New("Panic whilst running method"),
			})
		}
	}()

	rpcMethod := reflect.ValueOf(message.Transport.getRPCMethod(m.Data.Name))
	if len(m.Data.Parameters) != rpcMethod.Type().NumIn() {
		return errors.New("The number of parameters sent do not match the amount required.")
	}
	in := make([]reflect.Value, len(m.Data.Parameters))
	for k, param := range m.Data.Parameters {
		in[k] = reflect.ValueOf(param)
	}

	values := rpcMethod.Call(in)
	var result []interface{}
	for _, value := range values {
		if value.IsValid() {
			log.Print(value.Interface())
			result = append(result, value.Interface())
		}
	}

	message.Reply(Reply{
		Success: true,
		Result: result,
	})
	return nil
}