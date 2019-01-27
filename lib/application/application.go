package application

import (
	"fmt"
)

type Application interface {
	Start() error
	Stop() error
}

type Handler interface {
	Increase(count int) error
	Decrease(count int) error
	Unregister() error
}

type Event interface {
	fmt.Stringer
}

type Command interface {
	Run(b []byte) error
}
