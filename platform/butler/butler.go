package butler

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/actor"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Butler interface {
	prosecution.Prosecutor
}

type butler struct {
	logger    actor.LevelledLogger
	servants  ServantRepository
	qualifier ServantQualifier
}

func NewButler(
	logger actor.LevelledLogger,
	servants ServantRepository,
	qualifier ServantQualifier,
) Butler {
	return &butler{
		logger:    logger,
		servants:  servants,
		qualifier: qualifier,
	}
}

func (app butler) RequestForHelp(petition prosecution.Petition) error {
	app.logger.Log(actor.DEBUG, "%s command called", "RequestForHelp")

	// Find servants located in a specific distance from petition's location
	// which are able to serve the petition's topic
	if _, err := app.servants.FindByService(petition.Topic); nil != err {
		return err
	} else {

	}

	// Suggest the petition to all of them

	return fmt.Errorf("not implemented")
}
