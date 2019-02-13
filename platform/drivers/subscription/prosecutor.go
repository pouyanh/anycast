package subscription

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/pouyanh/anycast/lib/port"
	"github.com/pouyanh/anycast/platform/prosecution"
)

const (
	REQUEST_FOR_HELP string = "request_for_help"
)

func BindProsecutor(subscriber port.Subscriber, prosecutor prosecution.Prosecutor) (Subscription, error) {
	mq := &subscription{
		subscriber: subscriber,
	}

	if wp, err := mq.HireWorkers(REQUEST_FOR_HELP, prosecutor); nil != err {
		return nil, fmt.Errorf("worker assignment failed: %s", err)
	} else if err := wp.Increase(100); nil != err {
		return nil, fmt.Errorf("worker allocation failed: %s", err)
	}

	return mq, nil
}

// Request for Help Worker Pool
type rfhwp struct {
	app prosecution.Prosecutor

	count uint64 // Number of currently working workers

	// Receive only channel of incoming messages
	chmsg <-chan port.Message

	// Stop channel
	chstop chan bool
}

func (wp *rfhwp) Count() int {
	return int(atomic.LoadUint64(&wp.count))
}

func (wp *rfhwp) Increase(count int) error {
	if nil == wp.chstop {
		wp.chstop = make(chan bool)
	}

	go wp.work()

	return nil
}

func (wp *rfhwp) Decrease(count int) error {
	for i := 0; i < count; i++ {
		wp.chstop <- true
	}

	return nil
}

func (wp *rfhwp) Unregister() error {
	close(wp.chstop)

	return nil
}

// Workers main task
func (wp *rfhwp) work() {
	defer atomic.AddUint64(&wp.count, -1)
	atomic.AddUint64(&wp.count, 1)

	for {
		select {
		case <-wp.chstop:
			return

		case msg, ok := <-wp.chmsg:
			if !ok {
				return
			}

			var v prosecution.Petition
			if err := json.Unmarshal(msg.Data, &v); nil != err {
				// TODO: Inform about input error
			} else if err := wp.app.RequestForHelp(v); nil != err {
				// TODO: Inform about job error
			}
		}
	}
}
