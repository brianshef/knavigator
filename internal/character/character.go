package character

import (
	"fmt"

	"github.com/brianshef/knavigator/internal/data"
	d "github.com/brianshef/knavigator/internal/dice"
	"github.com/brianshef/knavigator/internal/inventory"
	"github.com/brianshef/knavigator/internal/names"
)

var dice = d.NewDice()

// Character defines all the attributes of a Knave character
type Character struct {
	Name      string               `json:"Name"`
	Abilities *abilities           `json:"Abilities"`
	Hitpoints *hitpoints           `json:"Hitpoints"`
	Armor     *inventory.Armor     `json:"Armor"`
	Weapon    *inventory.Weapon    `json:"Weapon"`
	Speed     *speed               `json:"Speed"`
	Traits    *traits              `json:"Traits"`
	Inventory *inventory.Inventory `json:"Inventory"`
}

// NewCharacter generates a new character
func NewCharacter(config *data.Config) *Character {
	abs := generateAbilities()

	inv := inventory.GenerateInventory(abs.Constitution.Defense, config)

	c := Character{
		Name: names.GenerateName(
			config.Names.FirstNames,
			config.Names.Surnames,
		),
		Abilities: abs,
		Hitpoints: generateHitPoints(abs.Constitution.Bonus),
		Armor:     inv.Armor,
		Weapon:    inv.Weapon,
		Speed:     generateSpeed(),
		Traits:    generateTraits(config.Traits),
		Inventory: inv,
	}
	return &c
}

// Print is a method which prints a string representation of a Character
func (c *Character) Print() {
	fmt.Printf(
		"\n%s\n%s\nHP: %s\nDefense: %d / +%d\nWeapon: %s\n%+v\nTraits: %s",
		c.Name,
		c.Abilities.String(),
		c.Hitpoints.String(),
		c.Armor.Defense,
		c.Armor.Defense-10,
		c.Weapon.String(),
		c.Inventory.String(),
		c.Traits.DescriptiveString(),
	)
}
