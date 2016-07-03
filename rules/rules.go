package rules

type Rules interface {
	MinStat() int
	MaxStat() int
	MinEnd() int
	MaxEnd() int
  RulesEdition() RulesEdition
	StatPointFromCoinRatio() float64
	StatPointForCoinRatio() float64
	MaxStatPointsFromCoin() int
	MaxStatPointsForCoin() int
	StartingCoin() int
	InitialStatPoints() int
	MaxStatPoints() int
	InitialBaseSkills() int
	InitialAdvancedSkills() int
	OptionalCanTradeCoinsForStats() bool
	SetOptionalCanTradeCoinsForStats(bool)
	ArmorGroups() []string
	VehicleGroups() []string
	WeaponGroups() []string
}

type RulesEdition int

const (
  UNDEFINED_EDITION = RulesEdition(0)
  FIRST_EDITION = RulesEdition(1)
  REVISED_EDITION = RulesEdition(2)
)

func GetAdventurersFirstEditionRules() Rules { return firstEdition }

func GetAdventurersRevisedRules() Rules { return revisedEdition }

var firstEdition *rulesGuide
var revisedEdition *rulesGuide

// 1e
// Weapon Groups. Blades, Bows and Slings, Clubs and Axes, Polearms, Power, Unarmed.
// Armor Groups. Unarmored, Light, Medium, Heavy, Power
//
// revised
// Weapon Groups. Blades, Bows and Slings, Clubs and Axes, Firearms, Polearms, Power, Unarmed.
// Armor Groups. Unarmored, Light, Medium, Heavy, Power
// Vehicle Groups. Aircraft, Ground Vehicles, Mounts, Spaceships, Watercraft.

func init() {

	firstEdition = &rulesGuide{
    armorGroups:            []string{"Unarmored", "Light", "Medium", "Heavy", "Power"},
    initialAdvancedSkills:  0,
    initialBaseSkills:      2,
    initialStatPoints:      6,
    maxEnd:                 13,
    maxStat:                5,
    maxStatPointsForCoin:   1,
    maxStatPointsFromCoin:  1,
    minEnd:                 1,
    minStat:                -1,
    rulesEdition:           FIRST_EDITION,
    startingCoin:           30,
    statPointForCoinRatio:  2.0,
    statPointFromCoinRatio: 0.2,
    vehicleGroups:          []string{},
    weaponGroups:           []string{"Blades", "Bows and Slings", "Clubs and Axes", "Polearms", "Power", "Unarmed"},
  }

	revisedEdition = &rulesGuide{
    armorGroups:           []string{"Unarmored", "Light", "Medium", "Heavy", "Power"},
    initialAdvancedSkills: 1,
    initialBaseSkills:     2,
    initialStatPoints:     6,
    maxEnd:                15,
    maxStat:               6,
    minEnd:                1,
    minStat:               -1,
    rulesEdition:          REVISED_EDITION,
    startingCoin:          30,
    vehicleGroups:         []string{"Aircraft", "Ground Vehicles", "Mounts", "Spaceships", "Watercraft"},
    weaponGroups:          []string{"Blades", "Bows and Slings", "Clubs and Axes", "Firearms", "Polearms", "Power", "Unarmed"},
  }
}

type rulesGuide struct {
  armorGroups                   []string
  initialAdvancedSkills         int
  initialBaseSkills             int
  initialStatPoints             int
  maxEnd                        int
  maxStat                       int
  maxStatPointsForCoin          int
  maxStatPointsFromCoin         int
  minEnd                        int
  minStat                       int
  optionalCanTradeCoinsForStats bool
  rulesEdition                  RulesEdition
  startingCoin                  int
  statPointForCoinRatio         float64
  statPointFromCoinRatio        float64
  vehicleGroups                 []string
  weaponGroups                  []string
}

func (guide *rulesGuide) MinStat() int { return guide.minStat }

func (guide *rulesGuide) MaxStat() int { return guide.maxStat }

func (guide *rulesGuide) MinEnd() int { return guide.minEnd }

func (guide *rulesGuide) MaxEnd() int { return guide.maxEnd }

func (guide *rulesGuide) StatPointFromCoinRatio() float64 { return guide.statPointFromCoinRatio }

func (guide *rulesGuide) StatPointForCoinRatio() float64 { return guide.statPointForCoinRatio }

func (guide *rulesGuide) MaxStatPointsForCoin() int { return guide.maxStatPointsForCoin }

func (guide *rulesGuide) MaxStatPointsFromCoin() int { return guide.maxStatPointsFromCoin }

func (guide *rulesGuide) StartingCoin() int { return guide.startingCoin }

func (guide *rulesGuide) InitialStatPoints() int { return guide.initialStatPoints }

func (guide *rulesGuide) MaxStatPoints() int {
	var possibleBonusStatPoints int
	if guide.optionalCanTradeCoinsForStats {
		possibleBonusStatPoints += guide.maxStatPointsFromCoin
	}
	return guide.initialStatPoints + possibleBonusStatPoints
}

func (guide *rulesGuide) InitialBaseSkills() int { return guide.initialBaseSkills }

func (guide *rulesGuide) InitialAdvancedSkills() int { return guide.initialAdvancedSkills }

func (guide *rulesGuide) OptionalCanTradeCoinsForStats() bool {
	return guide.optionalCanTradeCoinsForStats
}

func (guide *rulesGuide) SetOptionalCanTradeCoinsForStats(v bool) {
	guide.optionalCanTradeCoinsForStats = v
}

func (guide *rulesGuide) RulesEdition() RulesEdition { return guide.rulesEdition }

func (guide *rulesGuide) ArmorGroups() []string { return guide.armorGroups }

func (guide *rulesGuide) VehicleGroups() []string { return guide.vehicleGroups }

func (guide *rulesGuide) WeaponGroups() []string { return guide.weaponGroups }
