package subscription

import (
	"fmt"
	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/port"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Subscription interface {
	Shutdown() error
}

type subscription struct {
	subscriber port.Subscriber
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

func (mq *subscription) HireWorkers(topic string, app interface{}) (application.WorkerPool, error) {
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
				app:   app.(prosecution.Prosecutor),
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
