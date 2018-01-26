package golambdainvoke

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/rpc"
	"time"

	"github.com/aws/aws-lambda-go/lambda/messages"
	lc "github.com/aws/aws-lambda-go/lambdacontext"
)

const functioninvokeRPC = "Function.Invoke"

//Run a Go based lambda, passing the configured payload
//note that 'payload' can be anything that can be encoded by encoding/json
func Run(port int, payload interface{}) ([]byte, error) {
	return RunWithClientContext(port, payload, nil)
}

//RunWithClientContext a Go based lambda, passing the configured payload and ClientContext
//note that 'payload' can be anything that can be encoded by encoding/json
func RunWithClientContext(port int, payload interface{}, clientContext *lc.ClientContext) ([]byte, error) {
	request, err := createInvokeRequest(payload, clientContext)

	if err != nil {
		return nil, err
	}

	// 2. Open a TCP connection to the lambda
	client, err := rpc.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	// 3. Issue an RPC request for the Function.Invoke method
	var response messages.InvokeResponse

	if err = client.Call(functioninvokeRPC, request, &response); err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(response.Error.Message)
	}

	return response.Payload, nil
}

func createInvokeRequest(payload interface{}, clientContext *lc.ClientContext) (*messages.InvokeRequest, error) {
	payloadEncoded, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var clientContextEncoded []byte
	if clientContext != nil {
		b, err := json.Marshal(clientContext)

		if err != nil {
			return nil, err
		}

		clientContextEncoded = b
	}

	var t time.Time = time.Now()
	return &messages.InvokeRequest{
		Payload:      payloadEncoded,
		RequestId:    "0",
		XAmznTraceId: "",
		Deadline: messages.InvokeRequest_Timestamp{
			Seconds: int64(t.Unix()),
			Nanos:   int64(t.Nanosecond()),
		},
		InvokedFunctionArn:    "",
		CognitoIdentityId:     "",
		CognitoIdentityPoolId: "",
		ClientContext:         clientContextEncoded,
	}, nil
}
