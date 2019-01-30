package main

import (
	"os"
	"fmt"
	"flag"
	"syscall"
	"os/signal"

	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/lib/infrastructure/nats"
	"github.com/pouyanh/anycast/lib/infrastructure/logrus"
	"github.com/pouyanh/anycast/servant/application/servant"
)

var (
	CfgNatsUri string
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
	srvApp := &servant.Application{
		Services: *services,
	}

	// Run the application
	if err := srvApp.Start(); nil != err {
		panic(fmt.Errorf("error occurred during application start: %s", err))
	}

	// Handle shutdown
	if err := <-waitForShutdown(srvApp); nil != err {
		panic(err)
	}
}

func registerFlags() {
	flag.StringVar(&CfgNatsUri, "nats", "nats://nats.race:4222", "NATS URI")
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
