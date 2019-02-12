package butler

import (
	"testing"

	"github.com/pouyanh/anycast/lib/infrastructure"
)

func TestButler_RequestForHelp(t *testing.T) {
	logger := infrastructure.MockLevelledLogger{}

	app := NewButler(&logger)
	app.RequestForHelp()
}
