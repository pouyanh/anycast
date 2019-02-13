package logger

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/pouyanh/anycast/lib/port"
)

func TestMockLevelledLogger_Read(t *testing.T) {
	logger := &MockLevelledLogger{}
	buf := new(bytes.Buffer)

	var length int64
	var format, compiled string
	var args []interface{}

	// Test suite
	format = "testing mock levelled logger: %s"
	args = []interface{}{"read"}
	length = 9
	compiled = fmt.Sprintf(format, args...)

	go logger.Log(port.FATAL, format, args...)
	for i := int64(0); i < int64(len(compiled)) / length; i++ {
		if n, err := io.CopyN(buf, logger, length); 0 == n || nil != err {
			t.Errorf("Read error: %s", err)
		} else if text := buf.String(); compiled[:length * (i + 1)] != text {
			t.Errorf("Expected `%s` got `%s`", compiled[:length * (i + 1)], text)
		}
	}
}
