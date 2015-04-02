Snowblower
==========

A Golang based Snowplow collector that pushes received events directly to an SNS topic with a minimum of processing. From there, messages can be routed to any number of SQS queues or other SNS topic subscriptions as you wish.

In initial testing, this service is between 10 and 20 times more efficient than the Scala-based Snowplow Kinesis collector, based on the observation that we scaled down from 24 c3.xlarge machines to 2 in intial deployment. Besides the language difference, the performance improvements may be as much a result of writing to SNS instead of Kinesis or the minimization of processing of payloads and direct transmission of those payloads as byte arrays instead of using Thrift serialization.

Configuration
-------------

The following environment variables configure the operation of Snowblower:

- `SNS_TOPIC` Must contain the ARN of the SNS topic to send events to. **REQUIRED**
- `PORT` Optionally sets the port that the server listens to. Defaults to 8000.
- `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` Amazon Web Services credentials. If not set, Snowblower will attempt to use IAM Roles.
