# Snowblower

A lightweight high-performance Snowplow implementation in Go. Besides language choice, Snowblower differs from the official first party implementations its use of SNS/SQS as the data transport between stages and in its use of JSON-serialized Event records instead of Thrift-serialized CollectorPayloads.

One advantage to using SNS/SQS instead of Kinesis is that SQS scales transparently without explicit provisioning instruction.

## Performance and Cost

In initial testing, the collector service requires between 10 and 20 times fewer front-end compute resources than the Scala-based Snowplow Kinesis collector, based on the observation that we scaled down from 24 c3.xlarge machines to 2 on our initial deployment.

## Running

Snowblower has two commands:

- `collect` Runs the collector, sending events to SNS or SQS, acting as the second stage in a Snowplow pipeline.
- `enrich` Pulls events from SQS, enriches them, and sends them into storage into Postgres or Redshift, acting as the third stage in a Snowplow pipeline.


## Configuration

The following environment variables configure the operation of Snowblower:

- `SNS_TOPIC` Must contain the ARN of the SNS topic to send events to. **REQUIRED**
- `PORT` Optionally sets the port that the server listens to. Defaults to 8080.
- `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` Amazon Web Services credentials. If not set, Snowblower will attempt to use IAM Roles.
- `COOKIE_DOMAIN` if not set, a domain won't be set on the session cookie
