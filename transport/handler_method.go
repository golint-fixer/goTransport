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
	log.Print(m.Data)
	if message.Transport.getRPCMethod(m.Data.Name) == nil  {
		return errors.New("[404]: Unknown method:"+m.Data.Name)
	}
	return nil
}

func (m HandleMethod) Run(message *Message) (interface{}, error) {
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
		return nil, errors.New("The number of parameters sent do not match the amount required.")
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
	return result, nil
}