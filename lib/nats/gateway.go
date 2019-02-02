package nats

import (
	"net/http"

	"github.com/nats-io/go-nats"
	"github.com/pouyanh/anycast/lib/infrastructure"
)

type Gateway interface {
	infrastructure.Gateway

	Unhandle(pattern string)
}

type gateway struct {
	conn *nats.Conn
}

func (gw gateway) Handle(pattern string, handler http.Handler) {

}

func (gw gateway) Unhandle(pattern string) {

}

func NewGateway(url string, options ...nats.Option) (Gateway, error) {
	gw := &gateway{}
	if conn, err := nats.Connect(url, options...); nil != err {
		return nil, err
	} else {
		gw.conn = conn
	}

	return gw, nil
}
