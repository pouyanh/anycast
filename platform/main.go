package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pouyanh/anycast/lib/application"
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
	sosApp := &sos.Application{
		Services: *services,
	}

	// Handle shutdown
	handleShutdown(sosApp)

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

func handleShutdown(apps ...application.Application) {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)

		for _, app := range apps {
			if err := app.Stop(); nil != err {
				os.Exit(1)
			}
		}

		os.Exit(0)
	}()
}
