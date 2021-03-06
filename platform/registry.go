package main

import (
	"github.com/pouyanh/anycast/lib/actor"
	"github.com/pouyanh/anycast/lib/logrus"
	"github.com/pouyanh/anycast/lib/nats"
	"github.com/pouyanh/anycast/lib/redis"

	"github.com/pouyanh/anycast/platform/driven/repository"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Registry struct {
	Dictionary     actor.Dictionary
	SyncBroker     actor.SyncBroker
	AsyncBroker    actor.AsyncBroker
	LevelledLogger actor.LevelledLogger

	Repositories struct {
		Servants prosecution.ServantRepository
	}
}

func SetupRegistry() (*Registry, error) {
	registry := new(Registry)

	if v, err := logrus.NewLevelledLogger(actor.DEBUG); nil != err {
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
