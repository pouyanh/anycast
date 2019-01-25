package application

type Application interface {
	Start() error
	Stop() error
}

type Handler interface {
	Increase(count int) error
	Decrease(count int) error
	Unregister() error
}
