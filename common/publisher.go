package common

type Publisher interface {
	Publish(message string)
}
