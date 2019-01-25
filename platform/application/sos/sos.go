package sos

import "github.com/pouyanh/anycast/lib/infrastructure"

type Application struct {
	Services infrastructure.Services
}

func (a *Application) Start() error {
	return nil
}

func (a *Application) Stop() error {
	return nil
}
