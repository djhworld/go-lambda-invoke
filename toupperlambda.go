package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

func toUpperHandler(input string) (string, error) {
	return strings.ToUpper(input), nil
}

func main() {
	lambda.Start(toUpperHandler)
}
