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

func NewTerrarium(maxBeastsPerSpeciesAtStart, maxRandomPlantAtStart, w, h int, beastTypes ...*Beast) *Terrarium {
	totalBeastNumber := len(beastTypes) * maxBeastsPerSpeciesAtStart
	beasts := make([]*Beast, totalBeastNumber)
	plants := make([]*Plant, maxRandomPlantAtStart)
	
	//TODO: file the two slices with beasts and plants
	
	return &Terrarium{
		CurrentBeastNumber: totalBeastNumber,
		CurrentPlantNumber: maxRandomPlantAtStart,
		Beasts: beasts,
		Plants: plants,
		Width: w,
		Height: h,
	}
}
