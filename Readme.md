Snowblower
==========

A golang based Snowplow collector that pushes received events to an SNS topic. You can then route the messages to any number of SQS queues or other SNS topic subscriptions.

Configuration
-------------

The following environment variables configure the operation of Snowblower:

- `SNS_TOPIC` Must contain the ARN of the SNS topic to send events to. **REQUIRED**
- `PORT` Optionally sets the port that the server listens to. Defaults to 8000.
- `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` Amazon Web Services credentials. If not set, Snowblower will attempt to use IAM Roles.
