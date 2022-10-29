package entities

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type StructureBuilt struct {
	gorm.Model
	LocationID       uint       `gorm:"not null"`
	StructureID      uint       `gorm:"not null"`
	Structure        Structure  `gorm:"foreignkey:StructureID"`
	Location         Location   `gorm:"foreignkey:LocationID"`
	Level            uint       `gorm:"not null;default:0"`
	UpgradeOngoing   bool       `gorm:"not null"`
	UpgradeStartedAt time.Time  `gorm:"not null"`
	UpgradeEndedAt   *time.Time `gorm:"null"`
}

func (sb *StructureBuilt) UpgradeCost() map[string]float64 {
	costMap := map[string]float64{}
	for _, prereq := range sb.Structure.Prerequisites {
		preqConstraints := strings.Split(prereq.PrerequisiteFormula, ":")
		increaseFunc := preqConstraints[2]
		levelMultiplier, _ := strconv.ParseFloat(preqConstraints[1], 64)
		baseQuantity, _ := strconv.ParseFloat(preqConstraints[len(preqConstraints)-1], 64)
		fmt.Println(increaseFunc, levelMultiplier, baseQuantity, preqConstraints)
		var requiredQuantity float64
		switch increaseFunc {
		case "linear":
			requiredQuantity = baseQuantity * (levelMultiplier * float64(sb.Level))
		case "exp":
			requiredQuantity = baseQuantity * math.Pow(levelMultiplier, float64(sb.Level))
		}
		fmt.Println("requiredQuantity", requiredQuantity)
		costMap[preqConstraints[0]] = requiredQuantity
	}
	return costMap
}
