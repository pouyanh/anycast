package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/lib/infrastructure/logrus"
	"github.com/pouyanh/anycast/lib/infrastructure/nats"
	"github.com/pouyanh/anycast/lib/infrastructure/redis"
	"github.com/pouyanh/anycast/platform/application/selection"
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
		panic(fmt.Errorf("error occurred during infrastructure setup: %s", err))
	} else {
		services = v
	}

	// Create the application
	slcApp := &selection.Application{
		Services: *services,
	}

	// Run the application
	if err := slcApp.Start(); nil != err {
		panic(fmt.Errorf("error occurred during application start: %s", err))
	}

	// Handle shutdown
	if err := <-waitForShutdown(slcApp); nil != err {
		panic(err)
	}
}

func registerFlags() {
	flag.StringVar(&CfgNatsUri, "nats", "nats://nats.race:4222", "NATS URI")
	flag.StringVar(&CfgMongoDsn, "mongo", "mongodb://mongo.race/", "Mongo DB DSN")
	flag.StringVar(&CfgMysqlDsn, "mysql", "race:phi0lambda@tcp(mysql.race)/?parseTime=true", "Mysql DB DSN")
	flag.StringVar(&CfgRedisAddress, "redis", "redis.race:6379", "Redis Address")
}

func setupInfrastructure() (*infrastructure.Services, error) {
	services := new(infrastructure.Services)

	if v, err := logrus.NewLevelledLogger(infrastructure.DEBUG); nil != err {
		return nil, err
	} else {
		services.LevelledLogger = v
	}

	if v, err := nats.NewPubSubMessaging(CfgNatsUri); nil != err {
		return nil, err
	} else {
		services.AsyncBroker = v
	}

	if v, err := nats.NewReqRepMessaging(CfgNatsUri); nil != err {
		return nil, err
	} else {
		services.SyncBroker = v
	}

	if v, err := redis.NewKeyValueStorage(CfgRedisAddress); nil != err {
		return nil, err
	} else {
		services.KeyValueStorage = v
	}

	return services, nil
}

func waitForShutdown(apps ...application.Application) chan error {
	chshutdown := make(chan os.Signal)
	signal.Notify(chshutdown, syscall.SIGTERM, syscall.SIGINT)

	cherr := make(chan error)

	go func() {
		defer close(cherr)

		sig := <-chshutdown
		fmt.Printf("caught sig: %+v", sig)

		for _, app := range apps {
			if err := app.Stop(); nil != err {
				cherr <- err
			}
		}
	}()

	return cherr
}
