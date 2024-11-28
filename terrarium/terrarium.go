package terrarium

import (
)

type Terrarium struct {
	CurrentBeastNumber	int
	CurrentPlantNumber	int
	Beasts			[]*Beast
	Plants			[]*Plant
	Width			int
	Height			int
}

func (t *Terrarium) RunOneTurn() {
}

func NewTerrarium(maxRandomBeast, maxBeastsPerSpeciesAtStart, maxRandomPlantAtStart, w, h int, beastTypes ...*Beast) *Terrarium {
	totalBeastNumber := len(beastTypes) * maxBeastsPerSpeciesAtStart + maxRandomBeast * maxBeastsPerSpeciesAtStart
	beasts := make([]*Beast, totalBeastNumber)
	plants := make([]*Plant, maxRandomPlantAtStart)
	
	//TODO: file the two slices with beasts and plants
	for _, beast := range beastTypes {
		for j := 0; j < maxBeastsPerSpeciesAtStart; j++ {
			beasts = append(beasts, beast.CopyRandomGenre())
		}
	}

	for i := range maxRandomBeast {
		beast := NewRandomBeast(string(i))
		for j := 0; j < maxBeastsPerSpeciesAtStart; j++ {
			beasts = append(beasts, beast.CopyRandomGenre())
		}
	}

	for range maxRandomPlantAtStart {
		plants = append(plants, NewRandomPlant())
	}
	
	return &Terrarium{
		CurrentBeastNumber: totalBeastNumber,
		CurrentPlantNumber: maxRandomPlantAtStart,
		Beasts: beasts,
		Plants: plants,
		Width: w,
		Height: h,
	}
}
