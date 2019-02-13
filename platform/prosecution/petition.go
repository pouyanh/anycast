package prosecution

type Petition struct {
	Topic string `json:"topic"`
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
