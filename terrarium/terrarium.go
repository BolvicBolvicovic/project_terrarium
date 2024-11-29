package terrarium

import (
	"log"

	"github.com/BolvicBolvicovic/project_terrarium/lib"
)

type Food interface {
	AsFood(rate float64) float64
	GetPosition() *lib.Position
	AsSelf() interface{}
}

type Terrarium struct {
	CurrentBeastNumber	int		`json:"current_beast_number"`
	CurrentPlantNumber	int 		`json:"current_plant_number"`
	CurrentIteration	int 		`json:"current_iteration"`
	MaxIteration		int 		`json:"max_iteration"`
	MaxPlants		int		`json:"max_plants"`
	Beasts			[]*Beast	`json:"beasts"`
	Plants			[]*Plant	`json:"plants"`
	Width			float64		`json:"w"`
	Height			float64		`json:"h"`
}

func (t *Terrarium) RunOneTurn() {
	t.CurrentIteration += 1
	for i, beast := range t.Beasts {
		beast.Hungrier()
		if !beast.Alive { continue }
		if beast.Age >= 10 && beast.Age <= 20 {
			for j, b := range t.Beasts {
				if i == j { continue }
				if b.Alive && beast.Position.Collide(b.Position) && beast.CanMate(b) {
					beast.Mate(b)
					break
				}
			}
			continue
		}
		violenceState := beast.CanAttack()
		switch violenceState {
		case All:
			for j, b := range t.Beasts {
				if i == j { continue }
				if b.Alive && beast.Position.Collide(b.Position) {
					beast.Attack(b)
					if !b.Alive {
						beast.Eat(b.AsFood(beast.Genom.carnivor))
						break
					}
				}
			}
		case OtherSpecies:
			for j, b := range t.Beasts {
				if i == j { continue }
				if b.Alive && beast.Position.Collide(b.Position) && !beast.Genom.IsSameSpecies(b.Genom) {
					beast.Attack(b)
					if !b.Alive {
						beast.Eat(b.AsFood(beast.Genom.carnivor))
						break
					}
				} else if b.Alive && beast.Position.Collide(b.Position) && beast.CanMate(b) {
					beast.Mate(b)
					break
				}
			}
		case None:
			mate := false
			for j, b := range t.Beasts {
				if i == j { continue }
				if b.Alive && beast.Position.Collide(b.Position) && beast.CanMate(b) {
					beast.Mate(b)
					mate = true
					break
				}
			}
			if mate { continue }
			for _, p := range t.Plants {
				if p.Alive {
					p.Alive = false
					beast.Eat(p.AsFood(beast.Genom.herbivor))
					break
				}
			}
		}
	}
	newBeasts := make([]*Beast, 0)
	newPlants := make([]*Plant, 0)
	for _, beast := range t.Beasts {
		if beast.CanBirth() {
			newBeasts = append(newBeasts, beast.Birth())
		} else if beast.Genom.IsGestating {
			beast.GestationCycle += 1
		}
		if beast.Alive {
			newBeasts = append(newBeasts, beast)
		}
	}
	for _, plant := range t.Plants {
		if plant.Alive {
			newPlants = append(newPlants, plant)
		}
	}
	t.Beasts = newBeasts
	t.Plants = newPlants
	t.CurrentBeastNumber = len(t.Beasts)
	t.CurrentPlantNumber = len(t.Plants)
	for _, beast := range t.Beasts {
		targets := make([]Food, 0)
		for _, b := range t.Beasts {
        	    targets = append(targets, b)
        	}
        	for _, p := range t.Plants {
        	    targets = append(targets, p)
        	}
		potentialTargets := beast.GetTargetsInRange(targets, t.Width, t.Height)
		beast.LockTarget(potentialTargets)
		if beast.CurrentTarget != nil {
			beast.Position.MoveTowardPosition(beast.CurrentTarget.GetPosition(), t.Width, t.Height, beast.Genom.speed)
		}
	}
	if t.CurrentIteration % 2 == 0 {
		for _, plant := range t.Plants {
			prop := plant.Propagate()
			for _, p := range prop {
				if p != nil {
					t.Plants = append(t.Plants, p)
				}
			}
			plant.Alive = false
		}
	}
	if t.CurrentIteration % 15 == 0 {
		newPlants := make([]*Plant, 0)
		for i, plant := range t.Plants {
			if i > t.CurrentPlantNumber / 4 || i > t.MaxPlants { break }
			plant.Alive = true
			newPlants = append(newPlants, plant)
		}
		t.Plants = newPlants
		t.CurrentPlantNumber = len(t.Plants)
		log.Println("newPlants", t.CurrentPlantNumber)
	}
}

func NewTerrarium(maxRandomBeast, maxBeastsPerSpeciesAtStart, maxRandomPlantAtStart, maxIteration, maxPlants int, w, h float64, beastTypes ...*Beast) *Terrarium {
	totalBeastNumber := len(beastTypes) * maxBeastsPerSpeciesAtStart + maxRandomBeast * maxBeastsPerSpeciesAtStart
	beasts := make([]*Beast, 0)
	plants := make([]*Plant, 0)
	
	for _, beast := range beastTypes {
		if beast == nil { continue }
		for j := 0; j < maxBeastsPerSpeciesAtStart; j++ {
			beasts = append(beasts, beast.CopyRandomGenre())
		}
	}

	for i := range maxRandomBeast {
		beast := NewRandomBeast("RandomBeast" + string(i))
		for j := 0; j < maxBeastsPerSpeciesAtStart; j++ {
			beasts = append(beasts, beast.CopyRandomGenre())
		}
	}

	for range maxRandomPlantAtStart {
		plant := NewRandomPlant()
		plants = append(plants, plant)
	}
	
	return &Terrarium{
		CurrentBeastNumber: totalBeastNumber,
		CurrentPlantNumber: maxRandomPlantAtStart,
		CurrentIteration: 0,
		MaxIteration: maxIteration,
		MaxPlants: maxPlants,
		Beasts: beasts,
		Plants: plants,
		Width: w,
		Height: h,
	}
}
