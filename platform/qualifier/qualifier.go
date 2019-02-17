package qualifier

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/actor"
	"github.com/pouyanh/anycast/platform/prosecution"
)

type Qualifier interface {
	prosecution.ServantQualifier
}

type qualifier struct {
	logger actor.LevelledLogger
}

func NewQualifier(
	logger actor.LevelledLogger,
) Qualifier {
	return &qualifier{
		logger:   logger,
	}
}

func (app qualifier) QualifyByStatus(servants []prosecution.Servant, status prosecution.ServantStatus) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (app qualifier) QualifyByLocation(servants []prosecution.Servant, location prosecution.Point, radius int) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (app qualifier) QualifyByMinStars(servants []prosecution.Servant, stars int) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}

func (app qualifier) QualifyByMinMissions(servants []prosecution.Servant, missions int) ([]prosecution.Servant, error) {
	return nil, fmt.Errorf("not implemented")
}
