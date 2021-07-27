package character

import "fmt"

const (
	abilityRolls         = 3
	abilityDefenseScalar = 10
)

type Ability struct {
	Name    string
	Defense int
	Bonus   int
}

// String is a method for returning the string representation of an ability
func (a *Ability) String() string {
	return fmt.Sprintf("%s: %d / +%d", a.Name, a.Defense, a.Bonus)
}

func newAbility(name string) *Ability {
	lowest := dice.D6.RollKeepLowest(abilityRolls)
	return &Ability{
		Name:    name,
		Bonus:   lowest,
		Defense: lowest + abilityDefenseScalar,
	}
}

type abilities struct {
	Strength     *Ability
	Dexterity    *Ability
	Constitution *Ability
	Intelligence *Ability
	Wisdom       *Ability
	Charisma     *Ability
}

// String is a method for returning the string representation of the ability scores
func (a *abilities) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		a.Strength.String(),
		a.Dexterity.String(),
		a.Constitution.String(),
		a.Intelligence.String(),
		a.Wisdom.String(),
		a.Charisma.String(),
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
