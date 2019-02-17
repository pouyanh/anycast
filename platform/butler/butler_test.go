package butler

import (
	"testing"

	"github.com/pouyanh/anycast/platform/driven/logger"
	"github.com/pouyanh/anycast/platform/driven/repository"
	"github.com/pouyanh/anycast/platform/prosecution"
)

func TestButler_RequestForHelp(t *testing.T) {
	mll := logger.MockLevelledLogger{}
	servants := make(repository.InMemoryServantRepository, 0)

	app := NewButler(&mll, servants)
	app.RequestForHelp(prosecution.Petition{})
}
