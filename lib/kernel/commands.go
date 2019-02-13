package kernel

const (
	HELP      string = "help"
	VOLUNTEER string = "volunteer"
)

type Help struct {
	Location int `json:"location"`
}

type Volunteer struct {
	Location int `json:"location"`
}
