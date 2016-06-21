package rules

type Rules interface {
	GetMinStat() int
	GetMaxStat() int
	GetMinEnd() int
	GetMaxEnd() int
	StatPointFromCoinRatio() float64
	StatPointForCoinRatio() float64
	MaxStatPointsFromCoin() int
	MaxStatPointsForCoin() int
	StartingCoin() int
	InitialStatPoints() int
	MaxStatPoints() int
	InitialBaseSkills() int
	InitialAdvancedSkills() int
	GetOptionalCanTradeCoinsForStats() bool
	SetOptionalCanTradeCoinsForStats(bool)
	ArmorGroups() []string
	VehicleGroups() []string
	WeaponGroups() []string
}

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
		minStat:                -1,
		maxStat:                5,
		minEnd:                 1,
		maxEnd:                 13,
		startingCoin:           30,
		initialStatPoints:      6,
		initialBaseSkills:      2,
		initialAdvancedSkills:  0,
		maxStatPointsFromCoin:  1,
		maxStatPointsForCoin:   1,
		statPointFromCoinRatio: 0.2,
		statPointForCoinRatio:  2.0,
		armorGroups:            []string{"Unarmored", "Light", "Medium", "Heavy", "Power"},
		vehicleGroups:          []string{},
		weaponGroups:           []string{"Blades", "Bows and Slings", "Clubs and Axes", "Polearms", "Power", "Unarmed"},
	}

	revisedEdition = &rulesGuide{
		minStat:               -1,
		maxStat:               6,
		minEnd:                1,
		maxEnd:                15,
		startingCoin:          30,
		initialStatPoints:     6,
		initialBaseSkills:     2,
		initialAdvancedSkills: 1,
		armorGroups:           []string{"Unarmored", "Light", "Medium", "Heavy", "Power"},
		vehicleGroups:         []string{"Aircraft", "Ground Vehicles", "Mounts", "Spaceships", "Watercraft"},
		weaponGroups:          []string{"Blades", "Bows and Slings", "Clubs and Axes", "Firearms", "Polearms", "Power", "Unarmed"},
	}
}

type rulesGuide struct {
	minStat                       int
	maxStat                       int
	minEnd                        int
	maxEnd                        int
	startingCoin                  int
	initialStatPoints             int
	initialBaseSkills             int
	initialAdvancedSkills         int
	optionalCanTradeCoinsForStats bool
	maxStatPointsFromCoin         int
	maxStatPointsForCoin          int
	statPointFromCoinRatio        float64
	statPointForCoinRatio         float64
	armorGroups                   []string
	vehicleGroups                 []string
	weaponGroups                  []string
}

func (guide *rulesGuide) GetMinStat() int { return guide.minStat }

func (guide *rulesGuide) GetMaxStat() int { return guide.maxStat }

func (guide *rulesGuide) GetMinEnd() int { return guide.minEnd }

func (guide *rulesGuide) GetMaxEnd() int { return guide.maxEnd }

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

func (guide *rulesGuide) GetOptionalCanTradeCoinsForStats() bool {
	return guide.optionalCanTradeCoinsForStats
}

func (guide *rulesGuide) SetOptionalCanTradeCoinsForStats(v bool) {
	guide.optionalCanTradeCoinsForStats = v
}

func (guide *rulesGuide) ArmorGroups() []string { return guide.armorGroups }

func (guide *rulesGuide) VehicleGroups() []string { return guide.vehicleGroups }

func (guide *rulesGuide) WeaponGroups() []string { return guide.weaponGroups }
