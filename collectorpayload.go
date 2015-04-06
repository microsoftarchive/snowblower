package main

// CollectorPayloadSchema ...
const CollectorPayloadSchema = "iglu:com.snowplowanalytics.snowplow/CollectorPayload/thrift/1-0-0"

// CollectorPayload defines the structure of data posted from Snowplow trackers
type CollectorPayload struct {
	Schema        string   `json:"schema"`
	IPAddress     string   `json:"ipAddress"`
	Timestamp     int64    `json:"timestamp"`
	Encoding      string   `json:"encoding"`
	Collector     string   `json:"collector"`
	UserAgent     string   `json:"userAgent,omitempty"`
	RefererURI    string   `json:"refererUri,omitempty"`
	Path          string   `json:"path,omitempty"`
	QueryString   string   `json:"querystring,omitempty"`
	Body          string   `json:"body,omitempty"`
	Headers       []string `json:"headers,omitempty"`
	ContentType   string   `json:"contentType,omitEmpty"`
	Hostname      string   `json:"hostname,omitEmpty"`
	NetworkUserID string   `json:"networkUserId,omitEmpty"`
}
