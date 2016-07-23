package rules

import (
	"github.com/GeertJohan/go.rice"
	"github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type RulesEdition int

const (
	UNDEFINED_EDITION = RulesEdition(0)
	FIRST_EDITION     = RulesEdition(1)
	REVISED_EDITION   = RulesEdition(2)
)

func GetAdventurersFirstEditionRules() *Rules { return firstEdition }

func GetAdventurersRevisedRules() *Rules { return revisedEdition }

var firstEdition *Rules
var revisedEdition *Rules

func init() {

	firstEdition = &Rules{
		RulesEdition: FIRST_EDITION,
	}

	dataBox := rice.MustFindBox("../public/data/")
	data, err := dataBox.String("adventurers_1e.yml")
	if err != nil {
		logrus.WithField("err", err).Fatal("failed to load adventurers_1e.yml")
	}

	err = yaml.Unmarshal([]byte(data), firstEdition)
	if err != nil {
		logrus.WithField("err", err).Fatal("failed to unmarshal adventurers_1e.yml")
	}

	logrus.WithField("firstEdition", firstEdition).Info("loaded adventurers_1e.yml")

	revisedEdition = &Rules{
		ArmorGroups:           []string{"Unarmored", "Light", "Medium", "Heavy", "Power"},
		InitialAdvancedSkills: 1,
		InitialBasicSkills:    2,
		InitialStatPoints:     6,
		MaxEnd:                15,
		MaxStat:               6,
		MinEnd:                1,
		MinStat:               -1,
		RulesEdition:          REVISED_EDITION,
		InitialCoin:           30,
		VehicleGroups:         []string{"Aircraft", "Ground Vehicles", "Mounts", "Spaceships", "Watercraft"},
		WeaponGroups:          []string{"Blades", "Bows and Slings", "Clubs and Axes", "Firearms", "Polearms", "Power", "Unarmed"},
	}
}

type Rules struct {
	ArmorGroups                   []string `yaml:"armor_groups"`
	InitialAdvancedSkills         int      `yaml:"initial_advanced_skills"`
	InitialBasicSkills            int      `yaml:"initial_basic_skills"`
	InitialStatPoints             int      `yaml:"initial_stat_points"`
	InitialCoin                   int      `yaml:"initial_coin"`
	MaxEnd                        int      `yaml:"max_end"`
	MaxStat                       int      `yaml:"max_stat"`
	MaxStatPointsForCoin          int      `yaml:"max_stat_points_for_coin"`
	MaxStatPointsFromCoin         int      `yaml:"max_stat_points_from_coin"`
	MinEnd                        int      `yaml:"min_end"`
	MinStat                       int      `yaml:"min_stat"`
	OptionalCanTradeCoinsForStats bool     `yaml:"optional_can_trade_coins_for_stats"`
	RulesEdition                  RulesEdition
	StatPointForCoinRatio         float64  `yaml:"stat_point_for_coin_ratio"`
	StatPointFromCoinRatio        float64  `yaml:"stat_point_from_coin_ratio"`
	VehicleGroups                 []string `yaml:"vehicle_group"`
	WeaponGroups                  []string `yaml:"weapon_group"`
	Skills                        []Skill  `yaml:"skills"`
	Gear                          Gear     `yaml:"gear"`
}

type Skill struct {
	Name            string   `yaml:"name"`
	Basic           string   `yaml:"basic"`
	Advanced        string   `yaml:"advanced"`
	Specialized     bool     `yaml:"specialized"`
	Specializations []string `yaml:"specializations"`
}

type Gear struct {
	Armor   []Item `yaml:"armor"`
	Weapons []Item `yaml:"weapons"`
	Misc    []Item `yaml:"items"`
}

type Item struct {
	Name  string `yaml:"name"`
	Coins int    `yaml:"coins"`
	Notes string `yaml:"notes"`
	Heavy int    `yaml:"heavy"`
	Att   string `yaml:"att"`
	Def   string `yaml:"def"`
	Group string `yaml:"group"`
}

func (guide *Rules) MaxStatPoints() int {
	var possibleBonusStatPoints int
	if guide.OptionalCanTradeCoinsForStats {
		possibleBonusStatPoints += guide.MaxStatPointsFromCoin
	}
	logrus.WithField("OptionalCanTradeCoinsForStats", guide.OptionalCanTradeCoinsForStats).WithField("possibleBonusStatPoints", possibleBonusStatPoints).Debug("MaxStatPoints()")
	return guide.InitialStatPoints + possibleBonusStatPoints
}

func (guide *Rules) SetOptionalCanTradeCoinsForStats(v bool) {
	guide.OptionalCanTradeCoinsForStats = v
}
