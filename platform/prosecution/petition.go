package prosecution

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Petition struct {
	Topic    string `json:"topic"`
	Location Point  `json:"location"`
}
