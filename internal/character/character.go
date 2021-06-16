package character

import (
	"fmt"

	"github.com/brianshef/knavigator/internal/data"
	d "github.com/brianshef/knavigator/internal/dice"
	"github.com/brianshef/knavigator/internal/inventory"
)

var dice = d.NewDice()

// Character defines all the attributes of a Knave character
type Character struct {
	name      string
	abilities *abilities
	hitpoints *hitpoints
	armor     *inventory.Armor
	speed     *speed
	traits    *traits
	inventory *inventory.Inventory
}

// NewCharacter generates a new character
func NewCharacter(name string, config *data.Config) *Character {
	abs := generateAbilities()
	inv := inventory.GenerateInventory()
	c := Character{
		name:      name,
		abilities: abs,
		hitpoints: generateHitPoints(abs.Constitution.bonus),
		armor:     inv.Armor,
		speed:     generateSpeed(),
		traits:    generateTraits(config.Traits),
		inventory: inv,
	}
	return &c
}

// Print is a method which prints a string representation of a Character
func (c *Character) Print() {
	fmt.Printf(
		"\n%s\n%s\n%+v\nDefense: %d / +%d\n%+v\nTraits: %s",
		c.name,
		c.abilities.String(),
		c.hitpoints.String(),
		c.armor.Defense,
		c.armor.Defense-10,
		c.inventory.String(),
		c.traits.DescriptiveString(),
	)
}
