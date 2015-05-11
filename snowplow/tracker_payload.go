package snowplow

// TrackerPayload defines the structure of data posted from Snowplow trackers
type TrackerPayload struct {
	Schema string  `json:"schema"`
	Data   []Event `json:"data"`
}