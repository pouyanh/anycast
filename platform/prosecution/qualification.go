package prosecution

type ServantQualifier interface {
	QualifyByStatus(servants []Servant, status ServantStatus) ([]Servant, error)
	QualifyByLocation(servants []Servant, location Point, radius int) ([]Servant, error)
	QualifyByMinStars(servants []Servant, stars int) ([]Servant, error)
	QualifyByMinMissions(servants []Servant, missions int) ([]Servant, error)
}
