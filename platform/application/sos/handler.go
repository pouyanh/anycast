package sos

import (
	"sync"
	"bytes"
	"sync/atomic"

	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
)

func (a *Application) register(topic string, fn Controller) (application.Handler, error) {
	if ch, err := a.Services.PubSubMessaging.Subscribe(topic); nil != err {
		return nil, err
	} else {
		return &handler{
			wg:    a.wg,
			fn:    fn,
			chmsg: ch,
		}, nil
	}
}

type handler struct {
	wg     sync.WaitGroup
	count  int32
	chstop chan bool

	fn    Controller
	chmsg <-chan infrastructure.Message
}

func (h *handler) Increase(count int) error {
	if nil == h.chstop {
		h.chstop = make(chan bool)
	}

	for i := 0; i < count; i++ {
		h.wg.Add(1)
		go func() {
			defer atomic.AddInt32(&h.count, -1)
			defer h.wg.Done()

			atomic.AddInt32(&h.count, 1)

			for {
				select {
				case <-h.chstop:
					return

				case msg, ok := <-h.chmsg:
					if !ok {
						return
					}

					// TODO: Return response io.Writer or anything else
					if err := h.fn(bytes.NewReader(msg.Data), nil); nil != err {
						// TODO: Handle error
					}
				}
			}
		}()
	}

	return nil
}

func (h *handler) Decrease(count int) error {
	for i := 0; i < count; i++ {
		h.chstop <- true
	}

	return nil
}

func (h *handler) Unregister() error {
	close(h.chstop)

	return nil
}
