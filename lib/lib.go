package lib

import (
	"math/rand"
)

var (
	ColliderSize		= 0.1
)

func RandomChoice[T any](a, b T) T {
	if rand.Int() % 2 == 0 {
		return a
	}
	return b
}
