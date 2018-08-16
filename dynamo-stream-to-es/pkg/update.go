package dstream

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	es "github.com/olivere/elastic"
)

// Elasticsearch is an ES Client which will perform Updates for Dynamo Items
type Elasticsearch struct {
	*es.Client
}

// Update will index the item based on the details given
func (e *Elasticsearch) Update(d *Details, item map[string]events.DynamoDBAttributeValue) (*es.IndexResponse, error) {
	tmp := eventStreamToMap(item)
	var i interface{}
	_ = dynamodbattribute.UnmarshalMap(tmp, &i)
	resp, err := e.Index().
		Id(docID(d, item)).
		Type(docType(d)).
		Index(index(d)).
		BodyJson(i).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func docType(d *Details) (t string) {
	if d.RangeKey != "" {
		t = fmt.Sprintf("%s-%s", d.HashKey, d.RangeKey)
	} else {
		t = d.HashKey
	}
	return t
}

func docID(d *Details, item map[string]events.DynamoDBAttributeValue) (id string) {
	if d.RangeKey != "" {
		id = fmt.Sprintf("%s-%s", item[d.HashKey].String(), item[d.RangeKey].String())
	} else {
		id = item[d.HashKey].String()
	}

	return id
}

func index(d *Details) string {
	return strings.ToLower(d.TableName)
}

// Ugly Hack because
// events.DynamoDBAttributeValue != *dynamodb.AttributeValue
func eventStreamToMap(attribute interface{}) map[string]*dynamodb.AttributeValue {

	// Map to be returned
	m := make(map[string]*dynamodb.AttributeValue)

	tmp := make(map[string]events.DynamoDBAttributeValue)

	switch t := attribute.(type) {
	case map[string]events.DynamoDBAttributeValue:
		tmp = t
	case events.DynamoDBAttributeValue:
		tmp = t.Map()
	}

	for k, v := range tmp {
		switch v.DataType() {
		case events.DataTypeString:
			s := v.String()
			m[k] = &dynamodb.AttributeValue{
				S: &s,
			}
		case events.DataTypeBoolean:
			b := v.Boolean()
			m[k] = &dynamodb.AttributeValue{
				BOOL: &b,
			}
		case events.DataTypeMap:
			m[k] = &dynamodb.AttributeValue{
				M: eventStreamToMap(v),
			}
		}
	}
	return m
}
