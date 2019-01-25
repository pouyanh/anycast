package infrastructure

import "time"

type Message struct {
	Topic string
	Reply string
	Data  []byte
}

// Message Queue - Publish Subscribe service
type PubSubMessaging interface {
	Publish(topic string, reply string, data []byte) error
	Subscribe(topic string) (<-chan Message, error)
	Unsubscribe(topic string) error
}

// Message Queue - Request Reply service
type ReqRepMessaging interface {
	Request(topic string, message []byte, timeout time.Duration) ([]byte, error)
}

// Message Queue - Push Pull service
type PushPullMessaging interface {
}

// Message Queue - Exclusive Pair service
type ExclusivePairMessaging interface {
}
