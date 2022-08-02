package ddskgen_test

import (
	"strings"
	"testing"

	"github.com/noppikinatta/ddskgen"
)

func TestLoveInject(t *testing.T) {
	b := strings.Builder{}
	l := ddskgen.NewLove(&b)

	l.Inject()

	expected := "ラブ注入♡"

	if b.String() != expected {
		t.Errorf("expected outputs '%s', but got '%s'", expected, b.String())
	}
}
