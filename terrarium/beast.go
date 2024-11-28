package terrarium

import (
	"math"

	"github.com/BolvicBolvicovic/project_terrarium/lib"
)

var (
	HungerRate float64	= 0.02
	HungerDamageThreshold	= 0.6
	HungerDamage		= 0.05

	BirthDamage		= 0.05
)

type Beast struct {
	Name		string
	Generation	int
	Genom		*Genom
	Position	*lib.Position

	health		float64
	hunger		float64
	alive		bool
	gestationCycle	int
	embryon		*Beast
}

func NewRandomBeast(name string) *Beast {
	return &Beast{
		Name: name,
		Generation: 0,
		Genom: NewRandomGenom(),
		Position: &lib.Position{ X:0.0, Y:0.0 },
		health: 1,
		hunger: 0,
		alive: true,
		gestationCycle: 0,
		embryon: nil,
	}
}

func (b *Beast) Vision() float64 { return b.Genom.vision }

func (b *Beast) AsFood(carnivorRate float64) float64 {
	return b.Genom.stamina * carnivorRate
}

func (b *Beast) Eat(food float64) {
	b.health += food
	if b.health > 1 { b.health = 1 }
}

func (b *Beast) Hungrier() bool {
	searchForFood := false
	b.hunger *= (1 + b.Genom.metabolism) * HungerRate
	if b.hunger > b.Genom.HungerThreshold { searchForFood = true }
	if b.hunger > HungerDamageThreshold { b.health -= HungerDamage * (1 - b.Genom.stamina) }
	if b.health <= 0 { b.alive = false }
	return searchForFood
}

func (b *Beast) CanMate(mate *Beast) bool {
	return b.Genom.CanGestate != mate.Genom.CanGestate && !b.Genom.IsGestating && !mate.Genom.IsGestating && b.Genom.IsSameSpecies(mate.Genom)
}

func (b *Beast) Mate(mate *Beast) {
	name := func() string {
		if b.Name == mate.Name { return b.Name }
		return lib.RandomChoice(b.Name, mate.Name)
	}()
	// Since females are the ones gestating, they are the base for knowing what generation is the new born from
	generation := func() int {
		if b.Genom.CanGestate { return b.Generation + 1 }
		return mate.Generation + 1
	}()
	genom := NewRepoductionGenom(b.Genom, mate.Genom)
	b.embryon = &Beast{
		Name: name,
		Generation: generation,
		Genom: genom,
		Position: nil,
		health: 1,
		hunger: 0,
		alive: true,
		gestationCycle: 0,
		embryon: nil,
	}
}

func (b *Beast) CanBirth() bool {
	return b.Genom.IsGestating && b.gestationCycle >= b.Genom.GestationPeriod
}

func (b *Beast) Birth() *Beast {
	newBorn := b.embryon
	b.embryon = nil
	b.Genom.IsGestating = false
	newBorn.Position.X = b.Position.X
	newBorn.Position.Y = b.Position.Y
	b.health -= BirthDamage
	if b.health < 0 {
		b.health = 0
		b.alive = false
	}
	return newBorn
}

func (b *Beast) moveByBorder(targetX, targetY, width, height float64) {
	if math.Abs(targetX-b.Position.X) > (width / 2) {
		if b.Position.X < targetX {
			b.Position.X -= b.Genom.speed
			if b.Position.X < 0 {
				b.Position.X = width - 1
			}
		} else {
			b.Position.X += b.Genom.speed
			if b.Position.X >= width {
				b.Position.X = 0
			}
		}
	} else {
		b.moveByCenter(targetX, b.Position.Y)
	}

	if math.Abs(targetY-b.Position.Y) > (height / 2) {
		if b.Position.Y < targetY {
			b.Position.Y -= b.Genom.speed
			if b.Position.Y < 0 {
				b.Position.Y = height - 1
			}
		} else {
			b.Position.Y += b.Genom.speed
			if b.Position.Y >= height {
				b.Position.Y = 0
			}
		}
	} else {
		b.moveByCenter(b.Position.X, targetY)
	}
}

func (b *Beast) moveByCenter(x, y float64) {
	if b.Position.X < x {
		b.Position.X += b.Genom.speed
	} else if b.Position.X > x {
		b.Position.X -= b.Genom.speed
	}

	if b.Position.Y < y {
		b.Position.Y += b.Genom.speed
	} else if b.Position.Y > y {
		b.Position.Y -= b.Genom.speed
	}
}

func (b *Beast) MoveTowardXY(x, y, mapWidth, mapHeight float64) {
	costByCenter := math.Abs(x - b.Position.X) + math.Abs(y - b.Position.Y)
	
	deltaX := math.Abs(x - b.Position.X)
	deltaY := math.Abs(y - b.Position.Y)

	wrapX := math.Min(deltaX, mapWidth-deltaX)
	wrapY := math.Min(deltaY, mapHeight-deltaY)

	costByBorder := wrapX + wrapY
	if costByCenter < costByBorder {
		b.moveByCenter(x, y)
	} else {
		b.moveByBorder(x, y, mapWidth, mapHeight)
	}
}
