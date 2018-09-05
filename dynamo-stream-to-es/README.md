# Dynamo Stream To ES

This serverless project acts as an example for:
* Creating a Dynamo DB Table via Cloudformation
* Creating a 1 node Elasticserach Cluster via Cloudformation
* Creating a generic Go function which maps the keyspace from Dynamo DB to ElasticSearch

to deploy the example you must first have npm installed on your system [https://nodejs.org/en/download/](https://nodejs.org/en/download/)

cd into `sls-go-examples/dynamo-stream-to-es`

run npm install which will install the serverless CLI. Once you have serverless on your machine you may run `./node_modules/serverless/bin/serverless deploy` This will deploy the serverless project to your AWS Account. This example will very likely take ~15min to deploy (because of elasticsearch), and in production the deployment of data stores (dynamodb, rds variants, elasticsearch) should probably not be tied to an application deployment.



## To build for AWS Lambda (linux binaries):

`$ make`

This project depends on the `dep` tool, if you don't have it you may install it [here](https://github.com/golang/dep)

To deploy:

`$ npm install`

`$ export ES_URL=$your_es_url`

`$ export SLS_DEPLOYMENT_BUCKET=$your_deployment_bucket`

`$ make`

`$ node_modules/serverless/bin/serverless deploy`

## Seeding Your Dynamo Table with Data

make sure `cmd/seed-dynamo/main.go` refers to the table you've created, then run

`$ go run cmd/seed-dynamo/main.go`
