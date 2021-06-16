package knavigator

import (
	"log"

	"github.com/brianshef/knavigator/internal/character"
	"github.com/brianshef/knavigator/internal/data"
)

// GenerateCharacter generates a new Knave character
func GenerateCharacter() {
	config, err := new(data.Config).Load()
	if err != nil {
		log.Fatal(err)
	}

	c := character.NewCharacter("John Smith", config)
	c.Print()
}
