package main

import (
	"fmt"

	"github.com/pouyanh/anycast/platform/driver/subscription"
	"github.com/pouyanh/anycast/platform/prosecution"
)

// Attach driver adapters to driver ports of applications
func AttachDrivers(registry *Registry, prosecutor prosecution.Prosecutor) error {
	if _, err := subscription.BindProsecutor(registry.AsyncBroker, prosecutor); nil != err {
		return fmt.Errorf("error on nats registration: %s", err)
	} else {
		// TODO: Manage driver shutdown
	}

	return nil
}
