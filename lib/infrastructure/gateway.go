package infrastructure

import (
	"io"
)

type Handler func(reader io.Reader, writer io.Writer) error

type Gateway interface {
	Handle(pattern string, handler Handler) error
}

type multiGateways struct {
	gws []Gateway
}

func StickGateways(gateways ...Gateway) Gateway {
	return multiGateways{gws: gateways}
}

func (mgw multiGateways) Handle(pattern string, handler Handler) error {
	for _, gw := range mgw.gws {
		if err := gw.Handle(pattern, handler); nil != err {
			return err
		}
	}

	return nil
}
