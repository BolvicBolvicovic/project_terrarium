package lib

import (
	"math"
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

type Position struct {
	X	float64
	Y	float64
}

func (p *Position) Collide(objPos *Position) bool {
	return ColliderSize >= math.Abs(objPos.X - p.X) + math.Abs(objPos.Y - p.Y)
}

func RandomPosition() *Position {
	return &Position{
		X: rand.Float64(),
		Y: rand.Float64(),
	}
}
