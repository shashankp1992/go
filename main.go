package main

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}
func handler() {
	fmt.Println("Hello Lambda from Go")
}
