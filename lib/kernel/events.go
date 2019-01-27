package kernel

const HELP_REQUEST_RECEIVED string = "help_request_received"
type HelpRequestReceived struct {
	Location int `json:"location"`
}

func (e HelpRequestReceived) String() string {
	return HELP_REQUEST_RECEIVED
}

const SUGGESTION_PROPAGATED string = "suggestion_propagated"
type SuggestionPropagated int

func (e SuggestionPropagated) String() string {
	return SUGGESTION_PROPAGATED
}

const SERVANT_VOLUNTEERED string = "servant_volunteered"
type ServantVolunteered struct {
	Servant struct {
		Name string `json:"name"`
	} `json:"servant"`
}

func (e ServantVolunteered) String() string {
	return SERVANT_VOLUNTEERED
}

const VOLUNTEERING_ACCEPTED string = "volunteering_accepted"
type VolunteeringAccepted int

func (e VolunteeringAccepted) String() string {
	return VOLUNTEERING_ACCEPTED
}

const VOLUNTEERING_REJECTED string = "volunteering_rejected"
type VolunteeringRejected int

func (e VolunteeringRejected) String() string {
	return VOLUNTEERING_REJECTED
}
