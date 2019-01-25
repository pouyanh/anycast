package sos

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/pouyanh/anycast/lib/kernel"
	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
)

type Application struct {
	wg sync.WaitGroup
	handlers []application.Handler

	Services infrastructure.Services
}

func (a *Application) Start() error {
	// Check for services
	if err := a.check(); nil != err {
		return err
	}

	// Register handlers
	if err := a.setup(); nil != err {
		return err
	}

	return nil
}

func (a *Application) Stop() error {
	fmt.Println("SOS is going to be stopped")

	for _, h := range a.handlers {
		if err := h.Unregister(); nil != err {
			return err
		}
	}

	a.wg.Wait()

	fmt.Println("SOS stopped")

	return nil
}

func (a Application) check() error {
	if nil == a.Services.KeyValueStorage {
		return fmt.Errorf("no key/value storage service has been registered")
	}

	if nil == a.Services.PubSubMessaging {
		return fmt.Errorf("no publish/subscribe messaging service has been registered")
	}

	if nil == a.Services.ReqRepMessaging {
		return fmt.Errorf("no request/reply messaging service has been registered")
	}

	return nil
}

func (a *Application) setup() error {
	if h, err := a.register(kernel.CMD_HELP); nil != err {
		return err
	} else if err := h.Increase(10); nil != err {
		return err
	} else {
		a.handlers = append(a.handlers, h)
	}

	return nil
}

type handler struct {
	wg     sync.WaitGroup
	count  int32
	chstop chan bool

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

				case msg := <-h.chmsg:
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
	defer close(h.chstop)

	h.Decrease(int(atomic.LoadInt32(&h.count)))

	if c := atomic.LoadInt32(&h.count); 0 != c {
		return fmt.Errorf("%d workers have not been stoped", c)
	}

	return nil
}

func (a *Application) register(topic string) (application.Handler, error) {
	if ch, err := a.Services.PubSubMessaging.Subscribe(topic); nil != err {
		return nil, err
	} else {
		return &handler{
			chmsg: ch,
			wg:    a.wg,
		}, nil
	}
}
