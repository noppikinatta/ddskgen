package ddskgen

// Rand provides some of functions of math.Rand for testability.
type Rand interface {
	Intn(n int) int
}
