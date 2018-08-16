# SLS Go Examples

This repo contains 2 Serverless Go Examples. Each require an AWS Account and your AWS developer environment to be set up properly. For more info refer to: [https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html) for more info.

## Dynamo Stream to ES [⇒](https://github.com/beeceej/sls-go-examples/tree/master/dynamo-stream-to-es)

An example which showcases the power of Dynamo Streams and elastic search. Unfortunately this example doesn't come with code to deploy an elastic search cluster, or the Dynamo Tables. To play around with it, You must:

- Create a Dynamo Table
  - Make note of the Table ARN
  - Enable Dynamo Streams && Make note of the Table Stream ARN
- Create an ElasticSearch Cluster.
  - Make note of your cluster URL

For demo purposes I only tested with a completely open cluster (not for production use!!), If you do have a more secure cluster, unfortunately you may need to fiddle with the code to get this example talking to elastic search :(

the code in the directory `dynamo-stream-to-es/cmd/seed-dynamo`, can be used to seed random data into your Dynamo Table.

## Sentence Builder [⇒](https://github.com/beeceej/sls-go-examples/tree/master/sentence-builder)

An Example which showcases AWS Step Functions to orchestrate complex workflows in the cloud
