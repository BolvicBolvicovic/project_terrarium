package terrarium

import (
	"errors"
	"math"
	"math/rand"

	"github.com/BolvicBolvicovic/project_terrarium/lib"
)

var SameSpeciesThreshold = 0.01

type Genom struct {
	// All values are between zero and one.

	// When a threshold is reached, the beast starts to look for food.
	HungerThreshold 	float64

	// The higher the following caracteristic are, the more a beast needs food 
	vision			float64
	stamina			float64 // The more stamina a beast has the more it gives food when it dies. It also affects how well it endures hunger.
	speed			float64
	strenght		float64 // This more strenght a beast has the less speed it gets

	// Reproduction and gestation has an impact on food needs
	CanGestate		bool
	IsGestating		bool
	GestationPeriod 	int
	SexualMaturityAge	int

	// Has an impact on how much one kind of food gives a beast what it needs 
	herbivor		float64
	carnivor		float64

	// Composite factor that is the hunger rate
	metabolism		float64
}

func (g *Genom) UpdateMetabolism() {
	gestating := func() float64 {
		if g.IsGestating { return 0.25 } // 1.0 / 4 == 0.25 as (speed + strenght) == 1.0
		return 0.0
	}()
	g.metabolism = (g.vision + g.stamina + g.speed + g.strenght + gestating)
}

func (g *Genom) UpdateVision(v float64) {
	g.vision = v
	g.UpdateMetabolism()
}

func (g *Genom) UpdateStamina(v float64) {
	g.stamina = v
	g.UpdateMetabolism()
}

func (g *Genom) UpdateSpeed(v float64) {
	g.speed = v
	g.strenght = 1.0 - v
	g.UpdateMetabolism()
}

func (g *Genom) UpdateStrenght(v float64) {
	g.strenght = v
	g.speed = 1.0 - v
	g.UpdateMetabolism()
}

func (g *Genom) UpdateCarnivor(v float64) {
	g.carnivor = v
	g.herbivor = 1.0 - v
}

func (g *Genom) UpdateHerbivor(v float64) {
	g.herbivor = v
	g.carnivor = 1.0 - v
}

func (g *Genom) IsSameSpecies(otherGenom *Genom) bool {
	result := math.Abs(g.carnivor - otherGenom.carnivor) 
	result += math.Abs(g.vision - otherGenom.vision) 
	result += math.Abs(g.stamina - otherGenom.stamina)
	result += math.Abs(g.speed - otherGenom.speed)
	result += math.Abs(float64(g.GestationPeriod - otherGenom.GestationPeriod))
	result += math.Abs(float64(g.SexualMaturityAge - otherGenom.SexualMaturityAge))
	result += math.Abs(g.HungerThreshold - otherGenom.HungerThreshold)
	return result < SameSpeciesThreshold
}

func (g *Genom) CopyRandomGenre() *Genom {
	genom := &Genom{
		HungerThreshold: g.HungerThreshold,
		vision: g.vision,
		stamina: g.stamina,
		speed: g.speed,
		strenght: g.strenght,
		carnivor: g.carnivor,
		herbivor: g.herbivor,
		GestationPeriod: g.GestationPeriod,
		SexualMaturityAge: g.SexualMaturityAge,
		IsGestating: false,
		CanGestate: lib.RandomChoice(true, false),
	}
	genom.UpdateMetabolism()
	return genom
}

func NewGenom(HungerThreshold, ThirstThreshold, vision, stamina, speed, strenght, carnivor, herbivor float64, GestationPeriod, SexualMaturityAge int, CanGestate bool) (*Genom, error) {
	genes := &Genom{
		HungerThreshold: HungerThreshold,
		vision: vision,
		stamina: stamina,
		speed: speed,
		strenght: strenght,
		carnivor: carnivor,
		herbivor: herbivor,
		GestationPeriod: GestationPeriod,
		SexualMaturityAge: SexualMaturityAge,
		IsGestating: false,
		CanGestate: CanGestate,
	}
	if speed + strenght != 1.0 || carnivor + herbivor != 1.0 { return genes, errors.New("Combined values are over 1.0") }
	genes.UpdateMetabolism()
	return genes, nil
}

func NewRepoductionGenom(parent1, parent2 *Genom) *Genom {
	hThreshold := lib.RandomChoice(parent1.HungerThreshold, parent2.HungerThreshold)
	vision := lib.RandomChoice(parent1.vision, parent2.vision)
	stamina := lib.RandomChoice(parent1.stamina, parent2.stamina)
	speed := lib.RandomChoice(parent1.speed, parent2.speed)
	carnivor := lib.RandomChoice(parent1.carnivor, parent2.carnivor)
	gPeriod := lib.RandomChoice(parent1.GestationPeriod, parent2.GestationPeriod)
	sMatAge := lib.RandomChoice(parent1.SexualMaturityAge, parent2.SexualMaturityAge)
	genom := &Genom{
		HungerThreshold: hThreshold,
		vision: vision,
		stamina: stamina,
		GestationPeriod: gPeriod,
		SexualMaturityAge: sMatAge,
		IsGestating: false,
		CanGestate: lib.RandomChoice(true, false),
	}
	genom.UpdateCarnivor(carnivor)
	genom.UpdateSpeed(speed)

	return genom
}

func NewRandomGenom() *Genom {
	genom := &Genom{
		HungerThreshold: rand.Float64(),
		vision: rand.Float64(),
		stamina: rand.Float64(),
		GestationPeriod: int(rand.Float64() * 10),
		SexualMaturityAge: int(rand.Float64() * 10),
		IsGestating: false,
		CanGestate: lib.RandomChoice(true, false),
	}
	genom.UpdateCarnivor(rand.Float64())
	genom.UpdateSpeed(rand.Float64())

	return genom
}
