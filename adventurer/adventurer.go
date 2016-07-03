package adventurer

import (
	"adventurers_tools/math"
)

type Adventurer interface {
	Agility() int
	BaseDef() int
	CreationPoints() int
	Endurance() int
	Heroism() int
	Mind() int
	Strength() int
	SetAgility(int)
	SetCreationPoints(int)
	SetMind(int)
	SetStrength(int)
}

type adventurer struct {
	agility        int
	creationPoints int
	mind           int
	strength       int
}

func NewAdventurer() Adventurer {
	return &adventurer{creationPoints: 6}
}

func (adv *adventurer) Agility() int {
	return adv.agility
}

func (adv *adventurer) BaseDef() int {
	return adv.agility
}

func (adv *adventurer) CreationPoints() int {
	return adv.creationPoints
}

func (adv *adventurer) Endurance() int {
	return adv.strength + adv.agility + 3
}

func (adv *adventurer) Heroism() int {
	return math.MinInt(adv.strength, adv.agility, adv.mind) + 1
}

func (adv *adventurer) Mind() int {
	return adv.mind
}

func (adv *adventurer) Strength() int {
	return adv.strength
}

func (adv *adventurer) SetAgility(v int) {
	adv.agility = v
}

func (adv *adventurer) SetCreationPoints(v int) {
	adv.creationPoints = v
}

func (adv *adventurer) SetMind(v int) {
	adv.mind = v
}

func (adv *adventurer) SetStrength(v int) {
	adv.strength = v
}
