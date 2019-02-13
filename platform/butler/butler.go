package butler

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/port"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Butler interface {
	prosecution.Prosecutor
}

type butler struct {
	logger   port.LevelledLogger
	servants ServantRepository
}

func NewButler(
	logger port.LevelledLogger,
	srvrepo ServantRepository,
) Butler {
	return &butler{
		logger:   logger,
		servants: srvrepo,
	}
}

func (app butler) RequestForHelp(petition prosecution.Petition) error {
	app.logger.Log(port.DEBUG, "%s command called", "RequestForHelp")

	// Find servants located in a specific distance from petition's location
	// which are able to serve the petition's topic

	// Suggest the petition to all of them

	return fmt.Errorf("not implemented")
}
