package character

import "fmt"

type item struct {
	name  string
	slots int
}

// String is a method which returns a string representation of an item
func (i *item) String() string {
	return i.name
}

type inventory struct {
	armor             *armor
	dungeoneeringGear []*item
	generalGear       []*item
}

// String is a method which returns the string representation of the inventory
func (i *inventory) String() string {
	return fmt.Sprintf(
		"Armor: %s\nDungeoneering Gear: %v\nGeneral Gear: %v",
		i.armor.name,
		i.dungeoneeringGear,
		i.generalGear,
	)
}
