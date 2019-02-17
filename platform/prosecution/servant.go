package prosecution

type ServantStatus int

const (
	OFF ServantStatus = iota
	READY
	BUSY
)

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

