package inventory

import (
	"errors"
	"fmt"
)

type gearType int

const (
	misc gearType = iota
	dungeoneering
	general
)

// A Gear is a specific kind of item which can be held, used, or carried in inventory
type Gear struct {
	Item
	GearType gearType
}

func (g *Gear) addTo(inv *Inventory) error {
	if g.Slots > inv.AvailableSlots {
		return errors.New(noSlotsErrMsg)
	}

	inv.AvailableSlots -= g.Slots

	switch g.GearType {
	case dungeoneering:
		inv.DungeoneeringGear = append(inv.DungeoneeringGear, g)
	case general:
		inv.GeneralGear = append(inv.GeneralGear, g)
	case misc:
		inv.GeneralGear = append(inv.GeneralGear, g)
	default:
		inv.GeneralGear = append(inv.GeneralGear, g)
	}
	return nil
}

// String returns a string representation of a Gear
func (g *Gear) String() string {
	return fmt.Sprintf("%s (%v Slots)", g.Name, g.Slots)
}

func generateGear(gear []string, gt gearType) *Gear {
	return &Gear{
		Item: Item{
			Name:    gear[dice.D20.RollOnce()-1],
			Slots:   1,
			Quality: 3,
		},
		GearType: gt,
	}
}
