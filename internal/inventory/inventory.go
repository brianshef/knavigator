package inventory

import (
	"fmt"

	d "github.com/brianshef/knavigator/internal/dice"
)

const noSlotsErrMsg = "no slots available"

var dice = d.NewDice()

// Inventory is a specific collection of items carried, held, worn, or equipped by the character
type Inventory struct {
	TotalSlots        int
	AvailableSlots    int
	Armor             *Armor
	DungeoneeringGear []*Item
	GeneralGear       []*Item
}

// Addable defines an interface for any kind of item which can be added to an inventory
type Addable interface {
	addTo(*Inventory) error
}

// Add adds an addable item to the Inventory using its addTo method
func (i *Inventory) Add(a Addable) error {
	return a.addTo(i)
}

// String is a method which returns the string representation of the inventory
func (i *Inventory) String() string {
	return fmt.Sprintf(
		"Armor: %s\nDungeoneering Gear: %v\nGeneral Gear: %v\nSlots: %v / %v",
		i.Armor.String(),
		i.DungeoneeringGear,
		i.GeneralGear,
		i.AvailableSlots,
		i.TotalSlots,
	)
}

// GenerateInventory generates a new character inventory set
func GenerateInventory(slots int, armors, helmetsAndShields []string) *Inventory {
	inv := &Inventory{
		TotalSlots:     slots,
		AvailableSlots: slots,
	}

	inv.Add(generateArmor(armors, helmetsAndShields))

	return inv
}
