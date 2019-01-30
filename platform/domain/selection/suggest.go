package selection

type HelpRequest struct {
	Topic string
}

type Client interface {
	Request() HelpRequest
}

type client struct {
}

type Servant interface {
	Serve()
}

type servant struct {
}
