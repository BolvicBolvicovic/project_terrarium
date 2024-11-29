package terrarium

import (
	"github.com/BolvicBolvicovic/project_terrarium/lib"
)

var (
	HungerRate float64	= 0.02
	HungerDamageThreshold	= 0.6
	HungerDamage		= 0.05

	BirthDamage		= 0.05
)

type ViolenceState int

const (
	All		ViolenceState	= 2
	OtherSpecies	ViolenceState	= 1
	None		ViolenceState	= 0
)

type Beast struct {
	Name		string		`json:"name"`
	Generation	int		`json:"generation"`
	Genom		*Genom		`json:"genom"`
	Position	*lib.Position	`json:"position"`

	Health		float64		`json:"health"`
	Hunger		float64		`json:"hunger"`
	Alive		bool		`json:"alive"`
	GestationCycle	int		`json:"gestation_cycle"`
	Embryon		*Beast		`json:"embryon"`
	CurrentTarget	Food		`json:"current_target"`
}

func NewBeastRandomGenre(name string, genom *Genom) *Beast {
	return &Beast{
		Name: name,
		Generation: 0,
		Genom: genom,
		Position: lib.RandomPosition(),
		Health: 1,
		Hunger: 0,
		Alive: true,
		GestationCycle: 0,
		Embryon: nil,
		CurrentTarget: nil,
	}
}

func NewRandomBeast(name string) *Beast {
	return &Beast{
		Name: name,
		Generation: 0,
		Genom: NewRandomGenom(),
		Position: lib.RandomPosition(),
		Health: 1,
		Hunger: 0,
		Alive: true,
		GestationCycle: 0,
		Embryon: nil,
		CurrentTarget: nil,
	}
}

func (b *Beast) Vision() float64 { return b.Genom.vision }

// Food interface
func (b Beast) GetPosition() *lib.Position { return b.Position }
func (b Beast) AsFood(carnivorRate float64) float64 {
	return b.Genom.stamina * carnivorRate
}

func (b *Beast) Eat(food float64) {
	b.Health += food
	if b.Health > 1 { b.Health = 1 }
}

func (b *Beast) Hungrier() bool {
	searchForFood := false
	b.Hunger *= (1 + b.Genom.metabolism) * HungerRate
	if b.Hunger > b.Genom.HungerThreshold { searchForFood = true }
	if b.Hunger > HungerDamageThreshold { b.Health -= HungerDamage * (1 - b.Genom.stamina) }
	if b.Health <= 0 { b.Alive = false }
	return searchForFood
}

func (b *Beast) GetTargetsInRange(potentialTargets []Food, mapWidth, mapHeight float64) []Food {
	targetsInRange := make([]Food, 0)
	for _, target := range potentialTargets {
		pos := target.GetPosition()
		if b.Position.InRange(pos, b.Genom.vision, mapWidth, mapHeight) {
			targetsInRange = append(targetsInRange, target)
		}
	}
	return targetsInRange
}

func (b* Beast) LockTarget(inRangeTargets []Food) {
	var targetLocked Food = nil
	foodValue := 0.0
	currentTargetFoodValue := 0.0
	for _, target := range inRangeTargets {
		switch target.(type) {
			case Beast:
				currentTargetFoodValue = target.AsFood(b.Genom.carnivor)
				if targetLocked == nil {
					foodValue = currentTargetFoodValue
					targetLocked = target
				} else if foodValue < currentTargetFoodValue {
					foodValue = currentTargetFoodValue
					targetLocked = target
				}
			case Plant:
				currentTargetFoodValue = target.AsFood(b.Genom.herbivor)
				if targetLocked == nil {
					foodValue = currentTargetFoodValue
					targetLocked = target
				} else if foodValue < currentTargetFoodValue {
					foodValue = currentTargetFoodValue
					targetLocked = target
				}
		}
	}
	b.CurrentTarget = targetLocked
}

func (b *Beast) CanAttack() ViolenceState {
	if b.Genom.carnivor >= 0.95 && b.Hunger >= 0.95 { return All }
	if b.Genom.carnivor >= 0.5  { return OtherSpecies }
	return None
}

func (b *Beast) Attack(prey *Beast) {
	prey.Health -= b.Genom.strenght * 20 //* (1 - prey.Genom.stamina)
	if prey.Health <= 0 {
		prey.Alive = false
		prey.Health = 0
	}
}

func (b *Beast) CanMate(mate *Beast) bool {
	return b.Genom.CanGestate && !mate.Genom.CanGestate && !b.Genom.IsGestating && b.Genom.IsSameSpecies(mate.Genom)
}

func (b *Beast) Mate(mate *Beast) {
	name := func() string {
		if b.Name == mate.Name { return b.Name }
		return lib.RandomChoice(b.Name, mate.Name)
	}()
	// Since females are the ones Gestating, they are the base for knowing what generation is the new born from
	generation := func() int {
		if b.Genom.CanGestate { return b.Generation + 1 }
		return mate.Generation + 1
	}()
	genom := NewRepoductionGenom(b.Genom, mate.Genom)
	b.Embryon = &Beast{
		Name: name + "_" + string(generation),
		Generation: generation,
		Genom: genom,
		Position: nil,
		Health: 1,
		Hunger: 0,
		Alive: true,
		GestationCycle: 0,
		Embryon: nil,
		CurrentTarget: nil,
	}
}

func (b *Beast) CanBirth() bool {
	return b.Genom.IsGestating && b.GestationCycle >= b.Genom.GestationPeriod
}

func (b *Beast) Birth() *Beast {
	newBorn := b.Embryon
	b.Embryon = nil
	b.Genom.IsGestating = false
	b.GestationCycle = 0
	newBorn.Position.X = b.Position.X
	newBorn.Position.Y = b.Position.Y
	b.Health -= BirthDamage
	if b.Health < 0 {
		b.Health = 0
		b.Alive = false
	}
	return newBorn
}

func (b *Beast) CopyRandomGenre() *Beast {
	return &Beast{
		Name: b.Name + "_0",
		Generation: 0,
		Genom: b.Genom.CopyRandomGenre(),
		Position: b.Position.Copy(),
		Health: 1,
		Hunger: 0,
		Alive: true,
		GestationCycle: 0,
		Embryon: nil,
		CurrentTarget: nil,
	}
}
