package main

const simpleSNSEvent = `
{
  "Type" : "Notification",
  "MessageId" : "MESSAGEID",
  "TopicArn" : "TOPICARN",
  "Message" : "{\"schema\":\"iglu:com.snowplowanalytics.snowplow/CollectorPayload/thrift/1-0-0\",\"ipAddress\":\"1.1.1.1\",\"timestamp\":1428343554340,\"encoding\":\"\",\"collector\":\"Snowblower/0.0.1\",\"body\":\"{\\\"schema\\\":\\\"iglu:com.snowplowanalytics.snowplow/payload_data/jsonschema/1-0-2\\\",\\\"data\\\":[{\\\"uid\\\":\\\"9999\\\",\\\"res\\\":\\\"1080x1920\\\",\\\"dtm\\\":\\\"1428343748993\\\",\\\"tz\\\":\\\"Europe\\\\/Berlin\\\",\\\"e\\\":\\\"ue\\\",\\\"p\\\":\\\"mob\\\",\\\"tv\\\":\\\"andr-0.3.9\\\",\\\"tna\\\":\\\"com.wunderlist\\\",\\\"cx\\\":\\\"\\\\n\\\",\\\"ue_px\\\":\\\"\\\\n\\\",\\\"aid\\\":\\\"Wunderlist\\\\/3.3.1-production.497\\\",\\\"eid\\\":\\\"0de7c7ae-ecb3-4c9a-9c30-8b6718819a3c\\\",\\\"lang\\\":\\\"Deutsch\\\"}]}\",\"headers\":[\"Connection: keep-alive\",\"Accept-Encoding: gzip\",\"Content-Type: application/json; charset=utf-8\",\"X-Forwarded-For: 2.2.2.2\",\"X-Forwarded-Port: 443\",\"X-Forwarded-Proto: https\",\"Content-Length: 1176\"],\"contentType\":\"\",\"hostname\":\"\",\"networkUserId\":\"NETWORKID\"}",
  "Timestamp" : "2015-04-06T18:05:54.357Z",
  "SigningCertURL" : "URL",
  "UnsubscribeURL" : "URL"
}
`
