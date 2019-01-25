package main

import (
	"fmt"
	"flag"

	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/lib/infrastructure/nats"
	"github.com/pouyanh/anycast/lib/infrastructure/redis"

	"github.com/pouyanh/anycast/platform/application/sos"
)

var (
	CfgNatsUri      string
	CfgMongoDsn     string
	CfgMysqlDsn     string
	CfgRedisAddress string
)

func init() {
	registerFlags()
}

func main() {
	flag.Parse()

	// Setup infrastructure
	var services *infrastructure.Services
	if v, err := setupInfrastructure(); nil != err {
		panic(fmt.Errorf("error occured during infrastrucure setup: %s", err))
	} else {
		services = v
	}

	// Create the application
	sosApp := sos.Application{
		Services: *services,
	}

	// Run the application
	if err := sosApp.Start(); nil != err {
		panic(fmt.Errorf("error occured during application start: %s", err))
	}
}

func registerFlags() {
	flag.StringVar(&CfgNatsUri, "nats", "nats://nats.any:4222", "NATS URI")
	flag.StringVar(&CfgMongoDsn, "mongo", "mongodb://mongo.any/", "Mongo DB DSN")
	flag.StringVar(&CfgMysqlDsn, "mysql", "race:phi0lambda@tcp(mysql.race)/?parseTime=true", "Mysql DB DSN")
	flag.StringVar(&CfgRedisAddress, "redis", "redis.race:6379", "Redis Address")
}

func setupInfrastructure() (*infrastructure.Services, error) {
	services := new(infrastructure.Services)

	if v, err := nats.NewPubSubMessaging(CfgNatsUri); nil != err {
		return nil, err
	} else {
		services.PubSubMessaging = v
	}

	if v, err := nats.NewReqRepMessagingProvider(CfgNatsUri); nil != err {
		return nil, err
	} else {
		services.ReqRepMessaging = v
	}

	if v, err := redis.NewKeyValueStorage(CfgRedisAddress); nil != err {
		return nil, err
	} else {
		services.KeyValueStorage = v
	}

	return services, nil
}
