package butler

import (
	"github.com/pouyanh/anycast/platform/prosecution"
)

type ServantRepository interface {
	FindByService(topic string) ([]prosecution.Servant, error)
}

type Propagator interface {

}
