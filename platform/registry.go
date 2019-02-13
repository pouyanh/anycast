package main

import (
	"github.com/pouyanh/anycast/lib/logrus"
	"github.com/pouyanh/anycast/lib/nats"
	"github.com/pouyanh/anycast/lib/port"
	"github.com/pouyanh/anycast/lib/redis"

	"github.com/pouyanh/anycast/platform/prosecution"
	"github.com/pouyanh/anycast/platform/services/repository"
)

type Registry struct {
	Dictionary     port.Dictionary
	SyncBroker     port.SyncBroker
	AsyncBroker    port.AsyncBroker
	LevelledLogger port.LevelledLogger

	Repositories struct {
		Servants prosecution.ServantRepository
	}
}

func SetupRegistry() (*Registry, error) {
	registry := new(Registry)

	if v, err := logrus.NewLevelledLogger(port.DEBUG); nil != err {
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

	registry.Repositories.Servants = repository.OnlineServantRepository{
		Dictionary: registry.Dictionary,
	}

	return registry, nil
}
