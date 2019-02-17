package prosecution

type Prosecutor interface {
	RequestForHelp(Petition) error
}
