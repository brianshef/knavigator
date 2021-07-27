package inventory

import (
	"fmt"
	"log"

	d "github.com/brianshef/knavigator/internal/dice"
)

type archetype int

const (
	unarmed archetype = iota
	simpleMelee
	simpleRanged
	lightMelee
	lightRanged
	heavyMelee
	heavyRanged
)

var archetypeMappings = map[archetype][]string{
	unarmed: {"Unarmed"},
	simpleMelee: {
		"Dagger",
		"Cudgel",
		"Sickle",
		"Staff",
	},
	simpleRanged: {
		"Sling",
	},
	lightMelee: {
		"Spear",
		"Sword",
		"Mace",
		"Axe",
		"Flail",
	},
	lightRanged: {
		"Bow",
	},
	heavyMelee: {
		"Halberd",
		"War Hammer",
		"Long Sword",
		"Battle Axe",
	},
	heavyRanged: {
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
	// Equipped weapons do not take up inventory slots
	// if w.Item.Slots > inv.AvailableSlots {
	// 	return errors.New(noSlotsErrMsg)
	// }

	inv.Weapon = w

	// Equipped weapons do not take up inventory slots
	// inv.AvailableSlots -= w.Item.Slots

	return nil
}

func (w *Weapon) setStats(damage *d.Die, hands, slots, quality int) {
	w.Damage = damage
	w.Hands = hands
	w.Slots = slots
	w.Quality = quality
}

func (w *Weapon) getArchetype() archetype {
	for a, weapons := range archetypeMappings {
		if stringInSlice(w.Name, weapons) {
			return a
		}
	}
	return unarmed
}

// String returns a string representation of a Weapon
func (w *Weapon) String() string {
	return fmt.Sprintf("%s (%s Damage, %v Slots, %v Hands, Quality %v)", w.Name, w.Damage.String(), w.Slots, w.Hands, w.Quality)
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
	case unarmed:
		w.setStats(&dice.D4, 0, 0, 0)
	case simpleMelee:
		w.setStats(&dice.D6, 1, 1, 3)
	case simpleRanged:
		w.setStats(&dice.D4, 1, 1, 3)
	case lightMelee:
		w.setStats(&dice.D8, 1, 2, 3)
	case lightRanged:
		w.setStats(&dice.D6, 2, 2, 3)
	case heavyMelee:
		w.setStats(&dice.D10, 2, 3, 3)
	case heavyRanged:
		w.setStats(&dice.D8, 2, 3, 3)
	default:
		log.Fatalf("unrecognized weapon archetype %v for %s", w.getArchetype(), w.Name)
	}

	return w
}
