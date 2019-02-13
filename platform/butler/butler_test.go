package butler

import (
	"testing"

	"github.com/pouyanh/anycast/platform/prosecution"
	"github.com/pouyanh/anycast/platform/services/logger"
	"github.com/pouyanh/anycast/platform/services/repository"
)

func TestButler_RequestForHelp(t *testing.T) {
	mll := logger.MockLevelledLogger{}
	servants := make(repository.InMemoryServantRepository, 0)

	app := NewButler(&mll, servants)
	app.RequestForHelp(prosecution.Petition{})
}
