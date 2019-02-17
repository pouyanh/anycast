package nats

import (
	"bytes"
	"io"
	"net/http"

	"github.com/nats-io/go-nats"
	"github.com/pouyanh/anycast/lib/actor"
)

type Gateway interface {
	actor.HttpMux

	Unhandle(pattern string)
}

type gateway struct {
	conn *nats.Conn
}

func (gw gateway) Handle(pattern string, handler http.Handler) {
	if subs, err := gw.conn.Subscribe(pattern, HttpToNats(gw.conn, handler)); nil != err {

	} else if subs.IsValid() {

	}
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

func HttpToNats(conn *nats.Conn, handler http.Handler) nats.MsgHandler {
	return func(msg *nats.Msg) {
		if req, err := http.NewRequest(
			http.MethodPost,
			msg.Sub.Subject,
			bytes.NewReader(msg.Data),
		); nil != err {
			// TODO: Return the message
		} else {
			handler.ServeHTTP(NewHttpResponse(conn, msg), req)
		}
	}
}

type httpResponse struct {
	msg  *nats.Msg
	conn *nats.Conn

	statusCode  int
	headers     http.Header
	wroteHeader bool

	body   *io.PipeWriter
	reader *io.PipeReader
}

func NewHttpResponse(conn *nats.Conn, msg *nats.Msg) *httpResponse {
	ra := &httpResponse{
		conn: conn,
		msg:  msg,

		statusCode:  http.StatusOK,
		headers:     make(http.Header),
		wroteHeader: false,
	}

	ra.reader, ra.body = io.Pipe()

	go func(conn *nats.Conn, msg *nats.Msg, reader *io.PipeReader) {

	}(conn, msg, ra.reader)

	return ra
}

func (rw *httpResponse) Header() http.Header {
	return rw.headers
}

func (rw *httpResponse) Write(data []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}

}

func (rw *httpResponse) WriteHeader(statusCode int) {
	if rw.wroteHeader {
		return
	}

	rw.wroteHeader = true
	rw.statusCode = statusCode
}
