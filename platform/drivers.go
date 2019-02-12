package main

import (
	"fmt"

	"github.com/pouyanh/anycast/platform/butler"
	"github.com/pouyanh/anycast/platform/butler/drivers/nats"
)

// Attach driver adapters to driver ports of applications
func AttachDrivers(registry *Registry, btlrapp butler.Butler) error {
	if err := nats.Register(registry.AsyncBroker, btlrapp); nil != err {
		return fmt.Errorf("error on nats registration: %s", err)
	}

	return nil
}
