package ddskgen_test

import (
	"strings"
	"testing"

	"github.com/noppikinatta/ddskgen"
)

func TestTryDodosuko3(t *testing.T) {
	tests := map[string]struct {
		rand     ddskgen.Rand
		expected string
		succeed  bool
	}{
		"succeeded": {
			rand:     newTestRnd(0, 1, 1, 1),
			expected: "ドドスコスコスコ",
			succeed:  true,
		},
		"failed": {
			rand:     newTestRnd(1, 0, 1, 0),
			expected: "スコドドスコドド",
			succeed:  false,
		},
	}

	for name, v := range tests {
		f := func(t *testing.T) {
			b := strings.Builder{}
			l := ddskgen.NewLine(&b, v.rand)

			result := l.TryDodosuko3()

			if b.String() != v.expected {
				t.Errorf("expected outputs '%s', but got '%s'", v.expected, b.String())
			}
			if result != v.succeed {
				t.Errorf("expected returns '%t', but got '%t'", v.succeed, result)
			}
		}

		t.Run(name, f)
	}
}

type testRnd struct {
	values []int
	idx    int
}

func newTestRnd(v ...int) ddskgen.Rand {
	return &testRnd{values: v}
}

func (r *testRnd) Intn(n int) int {
	ret := r.values[r.idx] % n
	r.idx++
	r.idx %= len(r.values)
	return ret
}
