package inventory

import (
	"errors"
	"fmt"
	"log"
)

const defaultArmorDefense = 11

// An Armor is a specific type of item that can be worn and incurs a defensive bonus
type Armor struct {
	Item
	Helmet  bool
	Shield  bool
	Defense int
}

func (a *Armor) addTo(inv *Inventory) error {
	if a.Item.Slots > inv.AvailableSlots {
		return errors.New(noSlotsErrMsg)
	}

	inv.Armor = a
	inv.AvailableSlots -= a.Item.Slots

	return nil
}

func (a *Armor) setStats(defense, slots, quality int) {
	a.Defense = defense
	a.Slots = slots
	a.Quality = quality
}

func (a *Armor) addHelmet() {
	a.Helmet = true
	a.Defense++
	a.Slots++
	a.Name += " with helmet"
}

func (a *Armor) addShield() {
	a.Shield = true
	a.Defense++
	a.Slots++
	a.Name += " with shield"
}

// String returns a string representation of an Armor
func (a *Armor) String() string {
	return fmt.Sprintf("%s (Defense %v, %v Slots, Quality %v)", a.Name, a.Defense, a.Slots, a.Quality)
}

func generateArmor(armors, helmetsAndShields []string) *Armor {
	armorType := armors[dice.D20.RollOnce()-1]
	extrasType := helmetsAndShields[dice.D20.RollOnce()-1]

	a := &Armor{
		Item: Item{
			Name:    armorType,
			Slots:   0,
			Quality: 0,
		},
		Defense: defaultArmorDefense,
	}
	switch armorType {
	case "No armor":
		a.setStats(defaultArmorDefense, 0, 0)
	case "Gambeson":
		a.setStats(12, 1, 3)
	case "Brigandine":
		a.setStats(13, 2, 4)
	case "Chain":
		a.setStats(14, 3, 5)
	case "Half Plate":
		a.setStats(15, 4, 6)
	case "Full Plate":
		a.setStats(16, 5, 7)
	default:
		log.Fatalf("unrecognized armor %s", armorType)
	}

	switch extrasType {
	case "None":
		break
	case "Helmet":
		a.addHelmet()
	case "Shield":
		a.addShield()
	case "Helmet and Shield":
		a.addHelmet()
		a.addShield()
	default:
		log.Fatalf("unrecognized armor extras %s", extrasType)
	}

	// sanity check
	if a.Defense < defaultArmorDefense {
		log.Fatalf("armorType=%s extrasType=%s resulted in invalid armor Defense value %v", armorType, extrasType, a.Defense)
	}

	return a
}
