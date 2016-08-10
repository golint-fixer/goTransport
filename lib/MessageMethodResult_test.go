package lib

import "testing"

func TestMessageMethodResult_Received(t *testing.T) {
	var parameters []interface{}
	parameters = append(parameters, "1")
	parameters = append(parameters, 2)
	message := newMessageMethodResult(true, parameters)
	result := message.Received(nil)
	if result != nil {
		t.Fatal(result)
	}
}

func TestMessageMethodResult_Sending(t *testing.T) {
	var parameters []interface{}
	parameters = append(parameters, "1")
	parameters = append(parameters, 2)
	message := newMessageMethodResult(true, parameters)
	result := message.Sending()
	if result != nil {
		t.Fatal(result)
	}
}
