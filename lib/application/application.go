package application

type Application interface {
	Start() error
	Stop() error
}

type WorkerPool interface {
	Count() int
	Increase(count int) error
	Decrease(count int) error
	Unregister() error
}

type Command interface {
	Run(b []byte) error
}
