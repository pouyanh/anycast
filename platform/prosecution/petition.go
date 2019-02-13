package prosecution

type Petition struct {
	Topic string `json:"topic"`
}

type Client interface {
	Request() Petition
}

type Servant interface {
	Serve()
}
