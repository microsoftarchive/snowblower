package collector

// Publisher ...
type Publisher interface {
	Publish(message string)
}
