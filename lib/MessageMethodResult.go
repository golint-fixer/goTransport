package lib

type messageMethodResult struct {
	Message
	Result     bool          `json:"result"`
	Parameters []interface{} `json:"parameters"`
}

func init() {
	SetMessageDefinition(newMessageMethodResult(false, nil))
}

func newMessageMethodResult(result bool, parameters []interface{}) *messageMethodResult {
	return &messageMethodResult{
		Message:    NewMessage(MessageTypeMethodResult),
		Result:     result,
		Parameters: parameters,
	}
}

func (message messageMethodResult) Sending() error {
	return nil
}

func (message messageMethodResult) Received() error {
	return nil
}
