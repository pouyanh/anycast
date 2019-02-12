package butler

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/infrastructure"
)

type Butler interface {
	RequestForHelp() error
}

type butler struct {
	logger infrastructure.LevelledLogger
}

func NewButler(logger infrastructure.LevelledLogger) Butler {
	return &butler{
		logger: logger,
	}
}

func (app butler) RequestForHelp() error {
	app.logger.Log(infrastructure.DEBUG, "%s command called", "RequestForHelp")

	return fmt.Errorf("not implemented")
}
