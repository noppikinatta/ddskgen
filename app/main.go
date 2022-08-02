package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/noppikinatta/ddskgen"
)

func main() {
	w := os.Stdout
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	line := ddskgen.NewLine(w, rnd)
	love := ddskgen.NewLove(w)

	var count int

	for {
		if line.TryDodosuko3() {
			count++
		} else {
			count = 0
		}

		if count == 3 {
			love.Inject()
			break
		}
	}
}
