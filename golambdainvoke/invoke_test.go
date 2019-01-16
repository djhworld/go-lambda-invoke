package golambdainvoke

import (
	"encoding/json"
	"log"
	"testing"
)

// relies on the toupperlambda being available and listening on port 8001
func TestSuccessIT(t *testing.T) {
	input := "hello world"
	expected := "HELLO WORLD"

	response, err := Run(Input{
		Port: 8001,
		Payload: input,
	})

	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	var actual string
	if err = json.Unmarshal(response, &actual); err != nil {
		log.Println(err)
		t.FailNow()
	}

	if actual != expected {
		log.Println(err)
		t.FailNow()
	}
}
