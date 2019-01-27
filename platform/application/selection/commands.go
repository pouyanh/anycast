package selection

import (
	"encoding/json"

	"github.com/pouyanh/anycast/lib/kernel"
)

func (a *Application) setup() error {
	if h, err := a.listen(kernel.HELP_REQUEST_RECEIVED, suggestRequest{app: *a}); nil != err {
		return err
	} else if err := h.Increase(1000); nil != err {
		return err
	} else {
		a.handlers = append(a.handlers, h)
	}

	if h, err := a.listen(kernel.SERVANT_VOLUNTEERED, determineVolunteering{app: *a}); nil != err {
		return err
	} else if err := h.Increase(1000); nil != err {
		return err
	} else {
		a.handlers = append(a.handlers, h)
	}

	return nil
}

type suggestRequest struct {
	app Application
}

func (cmd suggestRequest) Run(b []byte) error {
	var event kernel.HelpRequestReceived

	if err := json.Unmarshal(b, &event); nil != err {
		return err
	}

	if err := cmd.app.SuggestRequest(event.Location); nil != err {
		return err
	}

	return nil
}

type determineVolunteering struct {
	app Application
}

func (cmd determineVolunteering) Run(b []byte) error {
	var event kernel.ServantVolunteered

	if err := json.Unmarshal(b, &event); nil != err {
		return err
	}

	if err := cmd.app.DetermineVolunteering(event.Servant.Name); nil != err {
		return err
	}

	return nil
}
