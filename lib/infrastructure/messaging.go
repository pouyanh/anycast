package infrastructure

import "time"

type Message struct {
	Topic string
	Reply string
	Data  []byte
}

// Message Queue - Publish Subscribe service
type AsyncBroker interface {
	Publish(topic string, reply string, data []byte) error
	Subscribe(topic string) (<-chan Message, error)
	Unsubscribe(topic string) error
}

// Message Queue - Request Reply service
type SyncBroker interface {
	Request(topic string, message []byte, timeout time.Duration) ([]byte, error)
}

// Message Queue - Push Pull service
type Pipeline interface {
	Push([]byte) error
	Pull() ([]byte, error)
}
