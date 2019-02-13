package prosecution

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Petition struct {
	Topic    string `json:"topic"`
	Location Point  `json:"location"`
}

type Prosecutor interface {
	RequestForHelp(Petition) error
}

type Client interface {
	Request() Petition
}

type Servant interface {
	Serve(Petition) error
}

type ServantRepository interface {
	Add(servant Servant) error
	Remove(sid int) error

	GetByID(sid int) (Servant, error)
	GetAll() ([]Servant, error) // TODO: Pagination

	FindByService(topic string) ([]Servant, error) // TODO: Pagination
	FindByLocation(location Point) ([]Servant, error) // TODO: Pagination
}
