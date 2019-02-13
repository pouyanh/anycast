package butler

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/infrastructure"
)

type Butler interface {
	RequestForHelp(hr HelpRequest) error
}

type butler struct {
	logger infrastructure.LevelledLogger
}

func NewButler(logger infrastructure.LevelledLogger) Butler {
	return &butler{
		logger: logger,
	}
}

type HelpRequest struct {

}

func (app butler) RequestForHelp(hr HelpRequest) error {
	app.logger.Log(infrastructure.DEBUG, "%s command called", "RequestForHelp")

	return fmt.Errorf("not implemented")
}
