package selection

import (
	"encoding/json"
	"github.com/pouyanh/anycast/lib/kernel"
)

func (a *Application) setup() error {
	if wp, err := a.createWorkers(kernel.HELP, HELP_REQUEST_RECEIVED); nil != err {
		return err
	} else if err := wp.Increase(1); nil != err {
		return err
	} else {
		a.wps = append(a.wps, wp)
	}

	if h, err := a.listen(HELP_REQUEST_RECEIVED, suggestRequest{app: *a}); nil != err {
		return err
	} else if err := h.Increase(1); nil != err {
		return err
	} else {
		a.wps = append(a.wps, wp)
	}

	if h, err := a.listen(SERVANT_VOLUNTEERED, determineVolunteering{app: *a}); nil != err {
		return err
	} else if err := h.Increase(1); nil != err {
		return err
	} else {
		a.wps = append(a.wps, wp)
	}

	return nil
}

type suggestRequest struct {
	app Application
}

func (cmd suggestRequest) Run(b []byte) error {
	var event HelpRequestReceived

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
	var event ServantVolunteered

	if err := json.Unmarshal(b, &event); nil != err {
		return err
	}

	if err := cmd.app.DetermineVolunteering(event.Servant.Name); nil != err {
		return err
	}

	return nil
}
