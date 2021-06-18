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
	name      string
	abilities *abilities
	hitpoints *hitpoints
	armor     *inventory.Armor
	weapon    *inventory.Weapon
	speed     *speed
	traits    *traits
	inventory *inventory.Inventory
}

// NewCharacter generates a new character
func NewCharacter(config *data.Config) *Character {
	abs := generateAbilities()

	inv := inventory.GenerateInventory(
		abs.Constitution.defense,
		config.Armor.Armor,
		config.Armor.HelmetsAndShields,
		config.Weapons.Weapons,
	)

	c := Character{
		name: names.GenerateName(
			config.Names.FirstNames,
			config.Names.Surnames,
		),
		abilities: abs,
		hitpoints: generateHitPoints(abs.Constitution.bonus),
		armor:     inv.Armor,
		weapon:    inv.Weapon,
		speed:     generateSpeed(),
		traits:    generateTraits(config.Traits),
		inventory: inv,
	}
	return &c
}

// Print is a method which prints a string representation of a Character
func (c *Character) Print() {
	fmt.Printf(
		"\n%s\n%s\n%s\nDefense: %d / +%d\nWeapon: %s\n%+v\nTraits: %s",
		c.name,
		c.abilities.String(),
		c.hitpoints.String(),
		c.armor.Defense,
		c.armor.Defense-10,
		c.weapon.String(),
		c.inventory.String(),
		c.traits.DescriptiveString(),
	)
}
