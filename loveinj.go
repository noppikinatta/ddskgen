package ddskgen

import (
	"fmt"
	"io"
)

// Love injects love.
type Love struct {
	writer io.Writer
}

// NewLove returns a new Love instance.
func NewLove(w io.Writer) *Love {
	return &Love{w}
}

// Inject outputs the last part of the line that Mikihisa Azuma might say.
func (l *Love) Inject() {
	fmt.Fprint(l.writer, "ラブ注入♡")
}
