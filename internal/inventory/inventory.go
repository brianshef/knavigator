package inventory

import (
	"fmt"

	"github.com/brianshef/knavigator/internal/data"
	d "github.com/brianshef/knavigator/internal/dice"
)

const noSlotsErrMsg = "no slots available"

var dice = d.NewDice()

// Inventory is a specific collection of items carried, held, worn, or equipped by the character
type Inventory struct {
	TotalSlots        int
	AvailableSlots    int
	Armor             *Armor
	Weapon            *Weapon
	DungeoneeringGear []*Gear
	GeneralGear       []*Gear
	StartingMoney     int
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
		"Armor: %s\nDungeoneering Gear: %v\nGeneral Gear: %v\nSlots: %v / %v\nOptional Starting Money: %v copper pieces",
		i.Armor.String(),
		i.DungeoneeringGear,
		i.GeneralGear,
		i.AvailableSlots,
		i.TotalSlots,
		i.StartingMoney,
	)
}

// GenerateInventory generates a new character inventory set
func GenerateInventory(slots int, config *data.Config) *Inventory {
	inv := &Inventory{
		TotalSlots:     slots,
		AvailableSlots: slots,
		StartingMoney:  generateStartingMoney(),
	}

	inv.Add(generateArmor(config.Armor.Armor, config.Armor.HelmetsAndShields))
	inv.Add(generateWeapon(config.Weapons.Weapons))
	for i := 0; i < 2; i++ {
		inv.Add(&Gear{
			Item: Item{
				Name:    "Rations",
				Slots:   1,
				Quality: 3,
			},
			GearType: general,
		})
	}
	inv.Add(generateGear(config.Gear.DungeoneeringGear, dungeoneering))
	inv.Add(generateGear(config.Gear.GeneralGear1, general))
	inv.Add(generateGear(config.Gear.GeneralGear2, general))

	return inv
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
