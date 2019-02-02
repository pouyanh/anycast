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

	gw.Handle("/resource", http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(req.Method))
		},
	))

	<-time.After(time.Millisecond * 100)

	methods := []string{http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPut}
	for _, method := range methods {
		req, _ := http.NewRequest(
			method,
			fmt.Sprintf(
				"http://%s/%s",
				addr.String(),
				"resource",
			),
			nil,
		)

		if resp, err := http.DefaultClient.Do(req); nil != err {
			t.Errorf("error on http `%s` request: %s", method, err)
		} else if body, err := ioutil.ReadAll(resp.Body); nil != err {
			t.Errorf("error on reading response body: %s", err)
		} else if !bytes.Equal([]byte(method), body) {
			t.Errorf("expected `%s` got `%s`", method, body)
		}
	}
}

func TestGatewayRequestHeaders(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	header, value := "phi0lambda", "Pouyan"

	gw.Handle("/headers", http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(req.Header.Get(header)))
		},
	))

	<-time.After(time.Millisecond * 100)

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"http://%s/%s",
			addr.String(),
			"headers",
		),
		nil,
	)
	req.Header.Add(header, value)

	if resp, err := http.DefaultClient.Do(req); nil != err {
		t.Errorf("error on http `%s` request: %s", http.MethodPost, err)
	} else if body, err := ioutil.ReadAll(resp.Body); nil != err {
		t.Errorf("error on reading response body: %s", err)
	} else if !bytes.Equal([]byte(value), body) {
		t.Errorf("expected `%s` got `%s`", value, body)
	}
}

func TestGatewayResponseStatus(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	status := http.StatusAccepted

	gw.Handle("/state", http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(status)
		},
	))

	<-time.After(time.Millisecond * 100)

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"http://%s/%s",
			addr.String(),
			"state",
		),
		nil,
	)

	if resp, err := http.DefaultClient.Do(req); nil != err {
		t.Errorf("error on http `%s` request: %s", http.MethodGet, err)
	} else if status != resp.StatusCode {
		t.Errorf("expected `%d` got `%d`", status, resp.StatusCode)
	}
}

func TestGatewayResponseHeaders(t *testing.T) {
	gw := NewGateway(addr.String(), nil)
	go gw.ListenAndServe()
	defer gw.Shutdown(context.Background())

	header, value := "phi0lambda", "Pouyan"

	gw.Handle("/headers", http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set(header, value)
		},
	))

	<-time.After(time.Millisecond * 100)

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"http://%s/%s",
			addr.String(),
			"headers",
		),
		nil,
	)

	if resp, err := http.DefaultClient.Do(req); nil != err {
		t.Errorf("error on http `%s` request: %s", http.MethodGet, err)
	} else if v := resp.Header.Get(header); value != v {
		t.Errorf("expected `%s` got `%s`", value, v)
	}
}
