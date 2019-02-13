package subscription

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Prosecutor interface {
	RequestForHelp(pt prosecution.Petition) error
}

const (
	REQUEST_FOR_HELP string = "request_for_help"
)

type Subscription interface {
	Shutdown() error
}

func BindProsecutor(subscriber infrastructure.Subscriber, app Prosecutor) (Subscription, error) {
	mq := &subscription{
		app:        app,
		subscriber: subscriber,
	}

	if wp, err := mq.HireWorkers(REQUEST_FOR_HELP); nil != err {
		return nil, fmt.Errorf("worker assignment failed: %s", err)
	} else if err := wp.Increase(100); nil != err {
		return nil, fmt.Errorf("worker allocation failed: %s", err)
	}

	return mq, nil
}

type subscription struct {
	app        Prosecutor
	subscriber infrastructure.Subscriber
	wps        map[string]application.WorkerPool
}

func (mq *subscription) Shutdown() error {
	// Wait for workers to stop
	for {
		count := 0
		for topic, wp := range mq.wps {
			if err := wp.Unregister(); nil != err {
				return err
			} else {
				delete(mq.wps, topic)
			}

			count += wp.Count()
		}

		if 0 == count {
			break
		}
	}

	return nil
}

func (mq *subscription) HireWorkers(topic string) (application.WorkerPool, error) {
	if nil == mq.wps {
		mq.wps = make(map[string]application.WorkerPool)
	}

	if _, ok := mq.wps[topic]; ok {
		return nil, fmt.Errorf("workers for `%s` has been already hired. try to increase them", topic)
	}

	var wp application.WorkerPool
	if chmsg, err := mq.subscriber.Subscribe(topic); nil != err {
		return nil, fmt.Errorf("subscription error: %s", err)
	} else {
		switch topic {
		case REQUEST_FOR_HELP:
			wp = &rfhwp{
				app:   mq.app,
				chmsg: chmsg,
			}

		default:
			mq.subscriber.Unsubscribe(topic)

			return nil, fmt.Errorf("factory unknown for `%s`", topic)
		}
	}

	mq.wps[topic] = wp

	return mq.wps[topic], nil
}

func (mq *subscription) FireWorkers(topic string) error {
	if wp, ok := mq.wps[topic]; !ok {
		return fmt.Errorf("no worker for `%s` has been hired", topic)
	} else if err := wp.Unregister(); nil != err {
		return err
	}

	return nil
}

// Request for Help Worker Pool
type rfhwp struct {
	count uint64 // Number of currently working workers

	app Prosecutor

	// Receive only channel of incoming messages
	chmsg <-chan infrastructure.Message

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
