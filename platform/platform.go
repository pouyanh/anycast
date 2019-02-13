package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pouyanh/anycast/platform/butler"
)

var (
	CfgNatsUri      string
	CfgMongoDsn     string
	CfgMysqlDsn     string
	CfgRedisAddress string
)

func init() {
	RegisterFlags()
}

func main() {
	flag.Parse()

	// Setup drivers
	var registry *Registry
	if v, err := SetupRegistry(); nil != err {
		panic(fmt.Errorf("error on drivers setup: %s", err))
	} else {
		registry = v
	}

	// Create Applications
	btlr := butler.NewButler(registry.LevelledLogger, registry.Repositories.Servants)

	// Attach the Drivers
	if err := AttachDrivers(registry, btlr); nil != err {
		panic(fmt.Errorf("drivers attachment failed: %s", err))
	}

	// Handle shutdown
	if err := <-WaitForShutdown(); nil != err {
		panic(err)
	}
}

func RegisterFlags() {
	flag.StringVar(&CfgNatsUri, "nats", "nats://nats.race:4222", "NATS URI")
	flag.StringVar(&CfgMongoDsn, "mongo", "mongodb://mongo.race/", "Mongo DB DSN")
	flag.StringVar(&CfgMysqlDsn, "mysql", "race:phi0lambda@tcp(mysql.race)/?parseTime=true", "Mysql DB DSN")
	flag.StringVar(&CfgRedisAddress, "redis", "redis.race:6379", "Redis Address")
}

func WaitForShutdown() chan error {
	chshutdown := make(chan os.Signal)
	signal.Notify(chshutdown, syscall.SIGTERM, syscall.SIGINT)

	cherr := make(chan error)

	go func() {
		defer close(cherr)

		sig := <-chshutdown
		fmt.Printf("caught sig: %+v", sig)

		// TODO: Shutdown drivers, applications and services
	}()

	return cherr
}
