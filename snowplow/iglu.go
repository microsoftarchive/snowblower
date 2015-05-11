package snowplow

type Iglu struct {
	Schema string      `json:"schema"`
	Data   interface{} `json:"data"`
}