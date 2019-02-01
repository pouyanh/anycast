package infrastructure

import (
	"net/http"
)

type Gateway interface {
	Handle(pattern string, handler http.Handler)
}

type multiGateways struct {
	gws []Gateway
}

func CombineGateways(gateways ...Gateway) Gateway {
	return multiGateways{gws: gateways}
}

func (mgw multiGateways) Handle(pattern string, handler http.Handler) {
	for _, gw := range mgw.gws {
		gw.Handle(pattern, handler)
	}
}
