package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/beeceej/sls-go-examples/sentence-builder/pkg/model"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func handler(d model.Data) (string, error) {
	return fmt.Sprintf("%s %s", d.Subject, phrases[rand.Intn(5)]), nil
}

func main() { lambda.Start(handler) }

var phrases = []string{
	"coded quickly",
	"cautiously jumped",
	"was notorius for mispelling thengs",
	"loves his dog",
	"always swipes right",
}
