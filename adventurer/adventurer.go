package adventurer

import (
	"adventurers_tools/math"
	"adventurers_tools/rules"
)

type Adventurer interface {
	Agility() int
	BaseDef() int
	StatCreationPoints() int
	BasicSkillCreationPoints() int
	Endurance() int
	Heroism() int
	Mind() int
	Strength() int
	SetAgility(int)
	SetStatCreationPoints(int)
	SetBasicSkillCreationPoints(int)
	SetMind(int)
	SetStrength(int)
}

type adventurer struct {
	agility                     int
	statCreationPoints          int
	basicSkillCreationPoints    int
	advancedSkillCreationPoints int
	mind                        int
	strength                    int
}

func NewAdventurer(theRules rules.Rules) Adventurer {
	return &adventurer{
		statCreationPoints:          theRules.InitialStatPoints,
		basicSkillCreationPoints:    theRules.InitialBasicSkills,
		advancedSkillCreationPoints: theRules.InitialAdvancedSkills,
	}
}

func (adv *adventurer) Agility() int {
	return adv.agility
}

func (adv *adventurer) BaseDef() int {
	return adv.agility
}

func (adv *adventurer) StatCreationPoints() int {
	return adv.statCreationPoints
}

func (adv *adventurer) BasicSkillCreationPoints() int {
	return adv.basicSkillCreationPoints
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

func (adv *adventurer) SetStatCreationPoints(v int) {
	adv.statCreationPoints = v
}

func (adv *adventurer) SetBasicSkillCreationPoints(v int) {
	adv.basicSkillCreationPoints = v
}

func (adv *adventurer) SetMind(v int) {
	adv.mind = v
}

func (adv *adventurer) SetStrength(v int) {
	adv.strength = v
}
