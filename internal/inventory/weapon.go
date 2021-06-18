package inventory

import (
	"errors"
	"fmt"
	"log"

	d "github.com/brianshef/knavigator/internal/dice"
)

type Archetype int

const (
	Unarmed Archetype = iota
	SimpleMelee
	SimpleRanged
	LightMelee
	LightRanged
	HeavyMelee
	HeavyRanged
)

var archetypeMappings = map[Archetype][]string{
	Unarmed: {"Unarmed"},
	SimpleMelee: {
		"Dagger",
		"Cudgel",
		"Sickle",
		"Staff",
	},
	SimpleRanged: {
		"Sling",
	},
	LightMelee: {
		"Spear",
		"Sword",
		"Mace",
		"Axe",
		"Flail",
	},
	LightRanged: {
		"Bow",
	},
	HeavyMelee: {
		"Halberd",
		"War Hammer",
		"Long Sword",
		"Battle Axe",
	},
	HeavyRanged: {
		"Crossbow",
	},
}

// Weapon is a specific type of item that can be equipped and is usable in combat
type Weapon struct {
	Item
	Damage *d.Die
	Hands  int
}

func (w *Weapon) addTo(inv *Inventory) error {
	if w.Item.Slots > inv.AvailableSlots {
		return errors.New(noSlotsErrMsg)
	}

	inv.Weapon = w
	inv.AvailableSlots -= w.Item.Slots

	return nil
}

func (w *Weapon) setStats(damage *d.Die, hands, slots, quality int) {
	w.Damage = damage
	w.Hands = hands
	w.Slots = slots
	w.Quality = quality
}

func (w *Weapon) getArchetype() Archetype {
	for a, weapons := range archetypeMappings {
		if stringInSlice(w.Name, weapons) {
			return a
		}
	}
	return Unarmed
}

// String returns a string representation of a Weapon
func (w *Weapon) String() string {
	return fmt.Sprintf("%s (%s damage, %v slots, %v hands, %v quality)", w.Name, w.Damage.String(), w.Slots, w.Hands, w.Quality)
}

func generateWeapon(weapons []string) *Weapon {
	weaponType := weapons[dice.D20.RollOnce()-1]

	w := &Weapon{
		Item: Item{
			Name:    weaponType,
			Slots:   0,
			Quality: 0,
		},
	}

	switch w.getArchetype() {
	case Unarmed:
		w.setStats(&dice.D4, 0, 0, 0)
	case SimpleMelee:
		w.setStats(&dice.D6, 1, 1, 3)
	case SimpleRanged:
		w.setStats(&dice.D4, 1, 1, 3)
	case LightMelee:
		w.setStats(&dice.D8, 1, 2, 3)
	case LightRanged:
		w.setStats(&dice.D6, 2, 2, 3)
	case HeavyMelee:
		w.setStats(&dice.D10, 2, 3, 3)
	case HeavyRanged:
		w.setStats(&dice.D8, 2, 3, 3)
	default:
		log.Fatalf("unrecognized weapon archetype %v for %s", w.getArchetype(), w.Name)
	}

	return w
}
