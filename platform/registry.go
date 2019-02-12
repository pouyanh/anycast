package main

import (
	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/lib/logrus"
	"github.com/pouyanh/anycast/lib/nats"
	"github.com/pouyanh/anycast/lib/redis"
)

type Registry struct {
	Dictionary infrastructure.Dictionary
	SyncBroker infrastructure.SyncBroker
	AsyncBroker infrastructure.AsyncBroker
	LevelledLogger infrastructure.LevelledLogger
}

func SetupRegistry() (*Registry, error) {
	registry := new(Registry)

	if v, err := logrus.NewLevelledLogger(infrastructure.DEBUG); nil != err {
		return nil, err
	} else {
		registry.LevelledLogger = v
	}

	if v, err := nats.NewAsyncBroker(CfgNatsUri); nil != err {
		return nil, err
	} else {
		registry.AsyncBroker = v
	}

	if v, err := nats.NewSyncBroker(CfgNatsUri); nil != err {
		return nil, err
	} else {
		registry.SyncBroker = v
	}

	if v, err := redis.NewDictionary(CfgRedisAddress); nil != err {
		return nil, err
	} else {
		registry.Dictionary = v
	}

	return registry, nil
}
