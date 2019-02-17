package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pouyanh/anycast/lib/actor"
)

// Redis

type cache struct {
	pool *redis.Pool
}

func (s cache) Put(key string, value interface{}) error {
	var b []byte

	switch v := value.(type) {
	case []byte:
		b = v

	case string:
		b = []byte(v)

	case fmt.Stringer:
		b = []byte(v.String())

	default:
		if data, err := json.Marshal(v); err != nil {
			return err
		} else {
			b = data
		}
	}

	_, err := s.pool.Get().Do("SET", key, b)

	return err
}

func (s cache) Has(key string) bool {
	if v, err := redis.Bool(s.pool.Get().Do("EXISTS", key)); err != nil {
		return false
	} else {
		return v
	}
}

func (s cache) Get(key string) interface{} {
	if v, err := s.pool.Get().Do("GET", key); err != nil {
		//TODO: Log error

		return nil
	} else {
		return v
	}
}

func (s cache) GetDefault(key string, value interface{}) interface{} {
	if v := s.Get(key); nil == v {
		return value
	} else {
		return v
	}
}

func (s cache) Delete(key string) error {
	if _, err := s.pool.Get().Do("DEL", key); err != nil {
		return err
	}

	return nil
}

func NewDictionary(dsn string) (actor.Dictionary, error) {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", dsn)
			if err != nil {
				return nil, err
			}

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")

			return err
		},
	}

	if _, err := pool.Get().Do("PING"); err != nil {
		return nil, err
	}

	return cache{pool}, nil
}
