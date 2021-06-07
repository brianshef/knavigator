package character

import (
	"fmt"

	d "github.com/brianshef/knavigator/internal/dice"
)

const (
	abilityRolls             = 3
	abilityDefenseScalar     = 10
	explorationSpeed     int = 120
	combatSpeed          int = 40
)

var dice = d.NewDice()

type ability struct {
	name    string
	defense int
	bonus   int
}

// Sprint is a method for returning the string representation of an ability
func (a *ability) Sprint() string {
	return fmt.Sprintf("%s: %d/+%d", a.name, a.defense, a.bonus)
}

func newAbility(name string) *ability {
	lowest := dice.D6.RollKeepLowest(abilityRolls)
	return &ability{
		name:    name,
		bonus:   lowest,
		defense: lowest + abilityDefenseScalar,
	}
}

type abilities struct {
	Strength     *ability
	Dexterity    *ability
	Constitution *ability
	Intelligence *ability
	Wisdom       *ability
	Charisma     *ability
}

// Sprint is a method for returning the string representation of the ability scores
func (a *abilities) Sprint() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		a.Strength.Sprint(),
		a.Dexterity.Sprint(),
		a.Constitution.Sprint(),
		a.Intelligence.Sprint(),
		a.Wisdom.Sprint(),
		a.Charisma.Sprint(),
	)
}

func generateAbilities() *abilities {
	return &abilities{
		Strength:     newAbility("Strength"),
		Dexterity:    newAbility("Dexterity"),
		Constitution: newAbility("Constitution"),
		Intelligence: newAbility("Intelligence"),
		Wisdom:       newAbility("Wisdom"),
		Charisma:     newAbility("Charisma"),
	}
}

type hitpoints struct {
	maximum     int
	healingRate int // 1d8 + CON bonus
}

func generateHitPoints(conBonus int) *hitpoints {
	max := dice.D8.Roll(1)[0]
	hr := dice.D8.Roll(1)[0] + conBonus

	return &hitpoints{maximum: max, healingRate: hr}
}

type speed struct {
	exploration int
	combat      int
}

type traits struct {
	physique    string
	face        string
	skin        string
	hair        string
	clothing    string
	virtue      string
	dice        string
	speech      string
	background  string
	misfortunes string
	alignment   string
}

func generateTraits() *traits {
	return &traits{}
}

type armor struct {
	name    string
	helmet  bool
	shield  bool
	defense int
}

func generateArmor() *armor {
	return &armor{
		name:    "none",
		helmet:  false,
		shield:  false,
		defense: 11,
	}
}

type inventory struct {
	armor             *armor
	dungeoneeringGear []string
	generalGear       []string
}

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
		speed:     &speed{exploration: explorationSpeed, combat: combatSpeed},
		traits:    generateTraits(),
		inventory: &inventory{
			armor:             armor,
			dungeoneeringGear: []string{"thing", "stuff"},
			generalGear:       []string{"knickknacks", "tchotchkes"},
		},
	}
	return &c
}

// Print is a method which prints a string representation of a Character
func (c *Character) Print() {
	fmt.Printf(
		"%s\n%s\nHitpoints: %+v\nTraits: %+v\nGear: %+v\n",
		c.name,
		c.abilities.Sprint(),
		*c.hitpoints,
		*c.traits,
		*c.inventory,
	)
}
