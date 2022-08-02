package ddskgen

import (
	"fmt"
	"io"
)

// Line outputs a part of lines that Mikihisa Azuma might say.
type Line struct {
	writer io.Writer
	rand   Rand
}

// NewLine returns a new Line instance.
func NewLine(w io.Writer, rand Rand) *Line {
	return &Line{w, rand}
}

// TryDodosuko3 tries to output a part of lines that Mikihisa Azuma might say, and returns true if suceeded.
func (l *Line) TryDodosuko3() bool {
	words := []string{"ドド", "スコ"}
	expectedIdxs := []int{0, 1, 1, 1}

	idxs := l.indices(len(words), len(expectedIdxs))

	succeeded := true

	for i, wi := range idxs {
		fmt.Fprint(l.writer, words[wi])
		if expectedIdxs[i] != wi {
			succeeded = false
		}
	}

	return succeeded
}

func (l *Line) indices(maxValue int, length int) []int {
	ret := make([]int, length)

	for i := 0; i < length; i++ {
		ret[i] = l.rand.Intn(maxValue)
	}

	return ret
}
