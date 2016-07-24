package adventurer

import (
	"github.com/dan-frohlich/AdventurersToolBelt/math"
	"github.com/dan-frohlich/AdventurersToolBelt/rules"
)

type Adventurer interface {
	Agility() int
	BaseDef() int
	StatCreationPoints() int
	BasicSkillCreationPoints() int
	AdvancedSkillCreationPoints() int
	Endurance() int
	Heroism() int
	Mind() int
	Strength() int
	Name() string
	Concept() string
	Rules() rules.Rules
	SourceFile() string
	BasicSkills() []rules.Skill
	AdvancedSkills() []rules.Skill
	SetAgility(int)
	SetStatCreationPoints(int)
	SetBasicSkillCreationPoints(int)
	SetAdvancedSkillCreationPoints(int)
	SetMind(int)
	SetStrength(int)
	SetName(string)
	SetConcept(string)
	SetSourceFile(string)
	AddBasicSkill(rules.Skill)
	AddAdvancedSkill(rules.Skill)
}

type adventurer struct {
	name                        string
	concept                     string
	sourceFile                  string
	theRules                    rules.Rules
	statCreationPoints          int
	basicSkillCreationPoints    int
	advancedSkillCreationPoints int
	agility                     int
	mind                        int
	strength                    int
	advancedSkills              []rules.Skill
	basicSkills                 []rules.Skill
}

type SkillKnowledge int

const (
	NO_SKILL       = SkillKnowledge(0)
	BASIC_SKILL    = SkillKnowledge(1)
	ADVANVED_SKILL = SkillKnowledge(2)
)

type KnownSkill struct {
	skill rules.Skill
	level SkillKnowledge
}

func NewAdventurer(theRules rules.Rules) Adventurer {
	return &adventurer{
		theRules:                    theRules,
		statCreationPoints:          theRules.InitialStatPoints,
		basicSkillCreationPoints:    theRules.InitialBasicSkills,
		advancedSkillCreationPoints: theRules.InitialAdvancedSkills,
		advancedSkills:              make([]rules.Skill, 0),
		basicSkills:                 make([]rules.Skill, 0),
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

func (adv *adventurer) AdvancedSkillCreationPoints() int {
	return adv.advancedSkillCreationPoints
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

func (adv *adventurer) Name() string {
	return adv.name
}

func (adv *adventurer) Concept() string {
	return adv.concept
}

func (adv *adventurer) SourceFile() string {
	return adv.sourceFile
}

func (adv *adventurer) Rules() rules.Rules {
	return adv.theRules
}

func (adv *adventurer) BasicSkills() []rules.Skill {
	return adv.basicSkills
}

func (adv *adventurer) AdvancedSkills() []rules.Skill {
	return adv.advancedSkills
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

func (adv *adventurer) SetAdvancedSkillCreationPoints(v int) {
	adv.advancedSkillCreationPoints = v
}

func (adv *adventurer) SetMind(v int) {
	adv.mind = v
}

func (adv *adventurer) SetStrength(v int) {
	adv.strength = v
}

func (adv *adventurer) SetName(v string) {
	adv.name = v
}

func (adv *adventurer) SetConcept(v string) {
	adv.concept = v
}

func (adv *adventurer) SetSourceFile(v string) {
	adv.sourceFile = v
}

func (adv *adventurer) AddBasicSkill(v rules.Skill) {
	adv.basicSkills = append(adv.basicSkills, v)
}

func (adv *adventurer) AddAdvancedSkill(v rules.Skill) {
	adv.advancedSkills = append(adv.advancedSkills, v)
}
