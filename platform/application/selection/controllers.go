package selection

import (
	"encoding/json"
	"fmt"

	"github.com/pouyanh/anycast/lib/application"
	"github.com/pouyanh/anycast/lib/kernel"
)

func (a *Application) setup() error {
	if h, err := a.listen(kernel.HELP_REQUEST_RECEIVED, SuggestRequest{app: *a}); nil != err {
		return err
	} else if err := h.Increase(1000); nil != err {
		return err
	} else {
		a.handlers = append(a.handlers, h)
	}

	if h, err := a.listen(kernel.SERVANT_VOLUNTEERED, DetermineVolunteering{app: *a}); nil != err {
		return err
	} else if err := h.Increase(1000); nil != err {
		return err
	} else {
		a.handlers = append(a.handlers, h)
	}

	return nil
}

type SuggestRequest struct {
	app Application
}

func (cmd SuggestRequest) Run(b []byte) error {
	var event kernel.HelpRequestReceived

	if err := json.Unmarshal(b, &event); nil != err {
		return err
	}

	if _, err := cmd.app.SuggestRequest(event.Location); nil != err {
		return err
	}

	return nil
}

func (a Application) SuggestRequest(location int) ([]application.Event, error) {
	// TODO: Find appropriate servants and give request header to 'em

	return nil, fmt.Errorf("not implemented")
}

type DetermineVolunteering struct {
	app Application
}

func (cmd DetermineVolunteering) Run(b []byte) error {
	var event kernel.ServantVolunteered

	if err := json.Unmarshal(b, &event); nil != err {
		return err
	}

	if _, err := cmd.app.DetermineVolunteering(event.Servant.Name); nil != err {
		return err
	}

	return nil
}

func (a Application) DetermineVolunteering(name string) ([]application.Event, error) {
	// TODO: Choose one of the volunteer servants, accept it and reject the others

	return nil, fmt.Errorf("not implemented")
}
