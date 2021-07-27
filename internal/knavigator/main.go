package knavigator

import (
	"github.com/brianshef/knavigator/internal/character"
	"github.com/brianshef/knavigator/internal/data"
)

// GenerateCharacter generates a new Knave character
func GenerateCharacter() (*character.Character, error) {
	config, err := data.NewConfig()
	if err != nil {
		return nil, err
	}

	c := character.NewCharacter(config)
	return c, nil
}
