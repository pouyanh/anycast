package main

import (
	"flag"
	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/lib/infrastructure/nats"

	"github.com/pouyanh/anycast/platform/application/sos"
)

var (
	CfgNatsUri      string
	CfgMongoDsn     string
	CfgMysqlDsn     string
	CfgRedisAddress string
)

var (
	PubSubMessaging infrastructure.PubSubMessaging
	ReqRepMessaging infrastructure.ReqRepMessaging
)

func init() {
	registerFlags()
}

func main() {
	flag.Parse()

	// Setup infrastructures
	setupInfrastructures()

	// Create the application
	sosApp := sos.Application{}
	sosApp.Setup()

	// Run the application
	sosApp.Run()
}

func registerFlags() {
	flag.StringVar(&CfgNatsUri, "nats", "nats://nats.any:4222", "NATS URI")
	flag.StringVar(&CfgMongoDsn, "mongo", "mongodb://mongo.any/", "Mongo DB DSN")
	flag.StringVar(&CfgMysqlDsn, "mysql", "race:phi0lambda@tcp(mysql.race)/?parseTime=true", "Mysql DB DSN")
	flag.StringVar(&CfgRedisAddress, "redis", "redis.race:6379", "Redis Address")
}

func setupInfrastructures() {
	if v, err := nats.NewPubSubMessaging(CfgNatsUri); nil != err {
		PubSubMessaging = v
	} else {
		panic(err)
	}

	if v, err := nats.NewReqRepMessagingProvider(CfgNatsUri); nil != err {
		ReqRepMessaging = v
	} else {
		panic(err)
	}
}
