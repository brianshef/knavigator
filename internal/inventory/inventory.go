package inventory

import "fmt"

// Inventory is a specific collection of items carried, held, worn, or equipped by the character
type Inventory struct {
	Armor             *Armor
	DungeoneeringGear []*Item
	GeneralGear       []*Item
}

// GenerateInventory generates a new character inventory set
func GenerateInventory() *Inventory {
	return &Inventory{
		Armor:             generateArmor(),
		DungeoneeringGear: []*Item{},
		GeneralGear:       []*Item{},
	}
}

// String is a method which returns the string representation of the inventory
func (i *Inventory) String() string {
	return fmt.Sprintf(
		"Armor: %s\nDungeoneering Gear: %v\nGeneral Gear: %v",
		i.Armor.Name,
		i.DungeoneeringGear,
		i.GeneralGear,
	)
}
