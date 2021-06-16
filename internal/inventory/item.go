package inventory

import "errors"

// Item represents any item a Knave character can hold, equip, or use
type Item struct {
	Name  string
	Slots int
}

func (i *Item) addTo(inv *Inventory) error {
	if i.Slots > inv.AvailableSlots {
		return errors.New("no slots avaiable")
	}

	inv.GeneralGear = append(inv.GeneralGear, i)
	inv.AvailableSlots -= i.Slots

	return nil
}

// String is a method which returns a string representation of an item
func (i *Item) String() string {
	return i.Name
}
