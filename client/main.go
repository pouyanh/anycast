package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pouyanh/anycast/client/application/client"
	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/lib/infrastructure/logrus"
	"github.com/pouyanh/anycast/lib/infrastructure/nats"
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
		panic(fmt.Errorf("error occured during infrastrucure setup: %s", err))
	} else {
		services = v
	}

	// Create the application
	clientApp := &client.Application{
		Services: *services,
	}

	// Handle shutdown
	handleShutdown(clientApp)

	// Run the application
	if err := clientApp.Start(); nil != err {
		panic(fmt.Errorf("error occured during application start: %s", err))
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
		services.PubSubMessaging = v
	}

	if v, err := nats.NewReqRepMessagingProvider(CfgNatsUri); nil != err {
		return nil, err
	} else {
		services.ReqRepMessaging = v
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
