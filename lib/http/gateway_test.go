package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"testing"
	"time"
)

var addr net.TCPAddr

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	addr.Port = 1025 + rand.Intn(65535-1025)
}

func TestGatewaySetup(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()

	<-time.After(time.Millisecond * 100)

	if _, err := http.DefaultClient.Get(
		fmt.Sprintf("http://%s", addr.String()),
	); nil != err {
		t.Errorf("GET request initiation error: %s", err)
	} else if err := gw.Shutdown(context.Background()); nil != err {
		t.Errorf("shutdown error: %s", err)
	}
}

func TestGatewayEchoHandler(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	gw.Handle("/echo", http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			io.Copy(rw, req.Body)
		},
	))

	<-time.After(time.Millisecond * 100)

	b := []byte("Pouyan")
	if resp, err := http.DefaultClient.Post(fmt.Sprintf(
		"http://%s/%s",
		addr.String(),
		"echo",
	), "text/plain", bytes.NewReader(b)); nil != err {
		t.Errorf("error on POST: %s", err)
	} else if body, err := ioutil.ReadAll(resp.Body); nil != err {
		t.Errorf("error on reading response body: %s", err)
	} else if !bytes.Equal(b, body) {
		t.Errorf("expected `%s` got `%s`", b, body)
	}
}

func TestGatewayRequestMethod(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	<-time.After(time.Millisecond * 100)
}

func TestGatewayRequestHeaders(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	<-time.After(time.Millisecond * 100)
}

func TestGatewayResponseStatus(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	<-time.After(time.Millisecond * 100)
}

func TestGatewayResponseHeaders(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	<-time.After(time.Millisecond * 100)
}

func TestGatewayAsAdapter(t *testing.T) {

}
