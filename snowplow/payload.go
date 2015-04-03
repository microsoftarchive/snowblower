package snowplow

import "encoding/json"

// Payload defines the structure of data posted from Snowplow trackers
type Payload struct {
	Schema string            `json:"schema"`
	Data   []json.RawMessage `json:"data"`
}
