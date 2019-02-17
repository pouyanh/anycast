package butler

import (
	"github.com/pouyanh/anycast/platform/prosecution"
)

type ServantRepository interface {
	FindByService(topic string) ([]prosecution.Servant, error)
}

type ServantQualifier interface {
	QualifyByStatus(servants []prosecution.Servant, status prosecution.ServantStatus) ([]prosecution.Servant, error)
	QualifyByLocation(servants []prosecution.Servant, location prosecution.Point, radius int) ([]prosecution.Servant, error)
}

type Propagator interface {

}
