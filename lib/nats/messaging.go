package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/pouyanh/anycast/lib/infrastructure"
)

type broker struct {
	conn *nats.Conn

	// Subscriptions list
	subs map[string]*nats.Subscription

	// Channel which when subscription is pushed it'll be added to subs
	chsubs chan *nats.Subscription

	// Channel which when topic is pushed it'll be deleted from subs
	chunsubs chan string
}

func (s broker) Publish(topic string, reply string, data []byte) error {
	return s.conn.PublishMsg(&nats.Msg{
		Subject: topic,
		Reply:   reply,
		Data:    data,
	})
}

func (s *broker) Subscribe(topic string) (<-chan infrastructure.Message, error) {
	if nil == s.chsubs {
		s.chsubs = make(chan *nats.Subscription)
		s.subs = make(map[string]*nats.Subscription)
		go func() {
			for {
				select {
				case sub := <-s.chsubs:
					s.subs[sub.Subject] = sub

				case topic := <-s.chunsubs:
					delete(s.subs, topic)
				}
			}
		}()
	}

	chmsg := make(chan *nats.Msg)
	if v, err := s.conn.ChanSubscribe(topic, chmsg); err != nil {
		return nil, err
	} else {
		s.chsubs <- v
	}

	ch := make(chan infrastructure.Message)
	go func() {
		for {
			msg := <-chmsg
			ch <- infrastructure.Message{
				Topic: msg.Subject,
				Reply: msg.Reply,
				Data:  msg.Data,
			}
		}
	}()

	return ch, nil
}

func (s *broker) Unsubscribe(topic string) error {
	if sub, ok := s.subs[topic]; !ok {
		return fmt.Errorf("topic `%s` is not in subscription list", topic)
	} else if err := sub.Unsubscribe(); err != nil {
		return err
	} else {
		s.chunsubs <- topic

		return nil
	}
}

func (s broker) Request(topic string, message []byte, timeout time.Duration) ([]byte, error) {
	if msg, err := s.conn.Request(topic, message, timeout); err != nil {
		return nil, err
	} else {
		return msg.Data, nil
	}
}

func NewAsyncBroker(url string) (infrastructure.AsyncBroker, error) {
	if conn, err := nats.Connect(url); err != nil {
		return nil, err
	} else {
		return &broker{conn: conn}, nil
	}
}

func NewSyncBroker(url string) (infrastructure.SyncBroker, error) {
	if conn, err := nats.Connect(url); err != nil {
		return nil, err
	} else {
		return broker{conn: conn}, nil
	}
}
