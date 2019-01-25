package application

type Application interface {
	Setup()
	Run()
	Stop()
}
