# Sentence Builder

To build for AWS Lambda (linux binaries):

`$ make`

This project depends on the `dep` tool, if you don't have it you may install it [here](https://github.com/golang/dep)

To deploy:

`$ npm install`

`$ export SLS_DEPLOYMENT_BUCKET=$your_deployment_bucket`

`$ make`

`$ node_modules/serverless/bin/serverless deploy`

## To Test

You should see something like this in your terminal once the application has deployed.

```endpoints:
  POST - https://6ordzy7rob.execute-api.us-east-1.amazonaws.com/dev/begin
functions:
  begin-sentence-builder: sentence-builder-dev-begin-sentence-builder
  subject: sentence-builder-dev-subject
  phrase: sentence-builder-dev-phrase
```

Take `POST - https://6ordzy7rob.execute-api.us-east-1.amazonaws.com/dev/begin`

and run

`$ curl -XPOST https://6ordzy7rob.execute-api.us-east-1.amazonaws.com/dev/begin -d '{"subject":"Gopher"}'`

to start a state machine execution

## Environment Variables

- SLS_DEPLOYMENT_BUCKET: The S3 bucket SLS will deploy to
