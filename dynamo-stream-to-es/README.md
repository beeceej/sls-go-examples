# Dynamo Stream To ES

To build for AWS Lambda (linux binaries):

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

## Environment Variables

- ES_URL: the URL to your Elastic Search Cluster
- SLS_DEPLOYMENT_BUCKET: The S3 bucket SLS will deploy to

## Note

You will need to edit line 10 and line 12 in serverless.yml to refer to your own dynamo table arns.
