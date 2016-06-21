package adventurer

import (
	"adventurers_tools/math"
)

type Adventurer interface {
	GetAgility() int
	GetBaseDef() int
	GetCreationPoints() int
	GetEndurance() int
	GetHeroism() int
	GetMind() int
	GetStrength() int
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

func (adv *adventurer) GetAgility() int {
	return adv.agility
}

func (adv *adventurer) GetBaseDef() int {
	return adv.agility
}

func (adv *adventurer) GetCreationPoints() int {
	return adv.creationPoints
}

func (adv *adventurer) GetEndurance() int {
	return adv.strength + adv.agility + 3
}

func (adv *adventurer) GetHeroism() int {
	return math.MinInt(adv.strength, adv.agility, adv.mind) + 1
}

func (adv *adventurer) GetMind() int {
	return adv.mind
}

func (adv *adventurer) GetStrength() int {
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
