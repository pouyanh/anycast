package restful

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/actor"
)

type Reporter interface {

}

func BindReporter(mux actor.HttpMux, reporter Reporter) error {
	return fmt.Errorf("not implemented")
}
