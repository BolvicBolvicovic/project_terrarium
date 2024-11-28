package lib

import (
	"math"
	"math/rand"
)

type Position struct {
	X	float64
	Y	float64
}

func RandomPosition() *Position {
	return &Position{
		X: rand.Float64(),
		Y: rand.Float64(),
	}
}

func (p *Position) Collide(objPos *Position) bool {
	return ColliderSize >= math.Abs(objPos.X - p.X) + math.Abs(objPos.Y - p.Y)
}

func (p *Position) Copy() *Position {
	return &Position{
		X: p.X,
		Y: p.Y,
	}
}

func (p *Position) moveByBorder(targetX, targetY, width, height, speed float64) {
	if math.Abs(targetX - p.X) > (width / 2) {
		if p.X < targetX {
			p.X -= speed
			if p.X < 0 {
				p.X = width - speed
			}
		} else {
			p.X += speed
			if p.X >= width {
				p.X = 0
			}
		}
	} else {
		p.moveByCenter(targetX, p.Y, speed)
	}

	if math.Abs(targetY - p.Y) > (height / 2) {
		if p.Y < targetY {
			p.Y -= speed
			if p.Y < 0 {
				p.Y = height - speed
			}
		} else {
			p.Y += speed
			if p.Y >= height {
				p.Y = 0
			}
		}
	} else {
		p.moveByCenter(p.X, targetY, speed)
	}
}

func (p *Position) moveByCenter(x, y, speed float64) {
	if p.X < x {
		p.X += speed
	} else if p.X > x {
		p.X -= speed
	}

	if p.Y < y {
		p.Y += speed
	} else if p.Y > y {
		p.Y -= speed
	}
}

func (p *Position) MoveTowardXY(x, y, mapWidth, mapHeight, speed float64) {
	costByCenter := math.Abs(x - p.X) + math.Abs(y - p.Y)
	
	deltaX := math.Abs(x - p.X)
	deltaY := math.Abs(y - p.Y)

	wrapX := math.Min(deltaX, mapWidth-deltaX)
	wrapY := math.Min(deltaY, mapHeight-deltaY)

	costByBorder := wrapX + wrapY
	if costByCenter <= costByBorder {
		p.moveByCenter(x, y, speed)
	} else {
		p.moveByBorder(x, y, mapWidth, mapHeight, speed)
	}
}

func (p *Position) InRange(objPos *Position, _range, mapWidth, mapHeight float64) bool {

	a := math.Abs(objPos.X - p.X) 
	b := math.Abs(objPos.Y - p.Y)
	costByCenter := a + b 
	
	deltaX := math.Abs(objPos.X - p.X)
	deltaY := math.Abs(objPos.Y - p.Y)

	wrapX := math.Min(deltaX, mapWidth-deltaX)
	wrapY := math.Min(deltaY, mapHeight-deltaY)

	costByBorder := wrapX + wrapY
	if costByCenter <= costByBorder {
		return _range >= math.Hypot(a, b)
	}
	return _range >= math.Hypot(wrapX, wrapY)
}
