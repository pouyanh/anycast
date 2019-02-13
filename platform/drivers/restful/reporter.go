package restful

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/infrastructure"
)

type Reporter interface {

}

func BindReporter(mux infrastructure.HttpMux, reporter Reporter) error {
	return fmt.Errorf("not implemented")
}
