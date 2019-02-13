package selection

import (
	"fmt"

	"github.com/pouyanh/anycast/lib/infrastructure"
)

func (a Application) SuggestRequest(location int) error {
	a.Services.LevelledLogger.Log(
		port.DEBUG,
		"Selection: Suggest Request called",
	)

	// TODO: Find appropriate servants and give request header to 'em

	// TODO: Dispatch events

	return fmt.Errorf("not implemented")
}

func (a Application) DetermineVolunteering(name string) error {
	a.Services.LevelledLogger.Log(
		port.DEBUG,
		"Selection: Determine Volunteering called",
	)

	// TODO: Choose one of the volunteer servants, accept it and reject the others

	// TODO: Dispatch events

	return fmt.Errorf("not implemented")
}
