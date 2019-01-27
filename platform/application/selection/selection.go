package selection

import (
	"fmt"
	"sync"

	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/infrastructure"
)

type Application struct {
	wg       sync.WaitGroup
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
	a.Services.LevelledLogger.Log(infrastructure.INFO, "SOS is going to be stopped")

	for _, h := range a.handlers {
		if err := h.Unregister(); nil != err {
			return err
		}
	}

	a.wg.Wait()

	a.Services.LevelledLogger.Log(infrastructure.INFO, "SOS has been stopped")

	return nil
}

func (a Application) check() error {
	if nil == a.Services.LevelledLogger {
		return fmt.Errorf("no levelled logger service has been registered")
	}

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
