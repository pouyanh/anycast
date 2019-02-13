package restful

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/port"
)

type Reporter interface {

}

func BindReporter(mux port.HttpMux, reporter Reporter) error {
	return fmt.Errorf("not implemented")
}
