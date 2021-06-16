package inventory

const defaultArmorDefense = 11

// An Armor is a specific type of item that can be worn and incurs a defensive bonus
type Armor struct {
	Item
	Helmet  bool
	Shield  bool
	Defense int
}

func generateArmor() *Armor {
	return &Armor{
		Item:    Item{Name: "none", Slots: 0},
		Helmet:  false,
		Shield:  false,
		Defense: defaultArmorDefense,
	}
}
