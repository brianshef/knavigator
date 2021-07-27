package api

import (
	"fmt"

	"github.com/brianshef/knavigator/internal/knavigator"
)

// Character defines all the attributes of a Knave character
// in a web-friendly output for html template engines
type Character struct {
	Name      string                 `json:"Name"`
	Abilities map[string]interface{} `json:"Abilities"`
	Hitpoints string                 `json:"Hitpoints"`
	Armor     string                 `json:"Armor"`
	Weapon    string                 `json:"Weapon"`
	Speed     string                 `json:"Speed"`
	Traits    string                 `json:"Traits"`
	Inventory map[string]interface{} `json:"Inventory"`
}

// GenerateCharacter() randomly generates a new character for rendering
func GenerateCharacter() (c *Character, err error) {
	res, err := knavigator.GenerateCharacter()
	if err != nil {
		return
	}

	c = &Character{
		Name:      res.Name,
		Hitpoints: res.Hitpoints.String(),
		Speed:     res.Speed.String(),
		Traits:    res.Traits.DescriptiveString(),
		Abilities: map[string]interface{}{
			res.Abilities.Strength.Name: map[string]interface{}{
				"Defense": fmt.Sprintf("%v", res.Abilities.Strength.Defense),
				"Bonus":   fmt.Sprintf("%v", res.Abilities.Strength.Bonus),
			},
			res.Abilities.Dexterity.Name: map[string]interface{}{
				"Defense": fmt.Sprintf("%v", res.Abilities.Dexterity.Defense),
				"Bonus":   fmt.Sprintf("%v", res.Abilities.Dexterity.Bonus),
			},
			res.Abilities.Constitution.Name: map[string]interface{}{
				"Defense": fmt.Sprintf("%v", res.Abilities.Constitution.Defense),
				"Bonus":   fmt.Sprintf("%v", res.Abilities.Constitution.Bonus),
			},
			res.Abilities.Intelligence.Name: map[string]interface{}{
				"Defense": fmt.Sprintf("%v", res.Abilities.Intelligence.Defense),
				"Bonus":   fmt.Sprintf("%v", res.Abilities.Intelligence.Bonus),
			},
			res.Abilities.Wisdom.Name: map[string]interface{}{
				"Defense": fmt.Sprintf("%v", res.Abilities.Wisdom.Defense),
				"Bonus":   fmt.Sprintf("%v", res.Abilities.Wisdom.Bonus),
			},
			res.Abilities.Charisma.Name: map[string]interface{}{
				"Defense": fmt.Sprintf("%v", res.Abilities.Charisma.Defense),
				"Bonus":   fmt.Sprintf("%v", res.Abilities.Charisma.Bonus),
			},
		},
		Inventory: map[string]interface{}{
			"Weapon":            res.Inventory.Weapon.String(),
			"Armor":             res.Inventory.Armor.String(),
			"Money":             fmt.Sprintf("%v copper pieces", res.Inventory.StartingMoney),
			"DungeoneeringGear": res.Inventory.PPrint("DungeoneeringGear"),
			"GeneralGear":       res.Inventory.PPrint("GeneralGear"),
			"Slots":             fmt.Sprintf("%v / %v", res.Inventory.AvailableSlots, res.Inventory.TotalSlots),
		},
	}

	return
}
