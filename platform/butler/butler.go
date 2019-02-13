package butler

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/infrastructure"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Butler interface {
	// Implementation of Prosecutor Port
	RequestForHelp(hr prosecution.Petition) error
}

type butler struct {
	logger infrastructure.LevelledLogger
}

func NewButler(logger infrastructure.LevelledLogger) Butler {
	return &butler{
		logger: logger,
	}
}

func (app butler) RequestForHelp(petition prosecution.Petition) error {
	app.logger.Log(infrastructure.DEBUG, "%s command called", "RequestForHelp")

	return fmt.Errorf("not implemented")
}
