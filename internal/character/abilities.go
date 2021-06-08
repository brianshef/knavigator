package character

import "fmt"

const (
	abilityRolls         = 3
	abilityDefenseScalar = 10
)

type ability struct {
	name    string
	defense int
	bonus   int
}

// String is a method for returning the string representation of an ability
func (a *ability) String() string {
	return fmt.Sprintf("%s: %d / +%d", a.name, a.defense, a.bonus)
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
