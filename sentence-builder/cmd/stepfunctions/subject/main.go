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

var adjectives = []string{
	"big And scary",
	"cute and cuddly",
	"old",
	"naive",
	"bubbly",
	"twisted",
	"hairy",
}

func handler(d model.Data) (interface{}, error) {
	d.Subject = fmt.Sprintf("The %s %s", adjectives[rand.Intn(6)], d.Subject)

	return d, nil
}

func main() { lambda.Start(handler) }
