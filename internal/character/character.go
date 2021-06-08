package character

import (
	"fmt"

	d "github.com/brianshef/knavigator/internal/dice"
)

var dice = d.NewDice()

// Character defines all the attributes of a Knave character
type Character struct {
	name      string
	abilities *abilities
	hitpoints *hitpoints
	armor     *armor
	speed     *speed
	traits    *traits
	inventory *inventory
}

// NewCharacter generates a new character
func NewCharacter(name string) *Character {
	abs := generateAbilities()
	armor := generateArmor()
	c := Character{
		name:      name,
		abilities: abs,
		hitpoints: generateHitPoints(abs.Constitution.bonus),
		armor:     armor,
		speed:     generateSpeed(),
		traits:    generateTraits(),
		inventory: &inventory{
			armor:             armor,
			dungeoneeringGear: []*item{},
			generalGear:       []*item{},
		},
	}
	return &c
}

// Print is a method which prints a string representation of a Character
func (c *Character) Print() {
	fmt.Printf(
		"\n%s\n%s\n%+v\nDefense: %d / +%d\nTraits: %+v\n%+v\n",
		c.name,
		c.abilities.String(),
		c.hitpoints.String(),
		c.armor.defense,
		c.armor.defense - 10,
		c.traits.String(),
		c.inventory.String(),
	)
}
