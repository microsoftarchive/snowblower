package aws

type SNSMessage struct {
	Type      string `json:"Type"`
	MessageID string `json:"MessageID"`
	Message   string `json:"Message"`
}