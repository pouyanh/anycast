package sos

import (
	"io"

	"github.com/pouyanh/anycast/lib/kernel"
)

func (a *Application) setup() error {
	if h, err := a.register(kernel.CMD_HELP, a.CommandHelp); nil != err {
		return err
	} else if err := h.Increase(1000); nil != err {
		return err
	} else {
		a.handlers = append(a.handlers, h)
	}

	return nil
}

func (a Application) CommandHelp(r io.Reader, rw io.Writer) error {
	return nil
}
