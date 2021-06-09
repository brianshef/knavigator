package knavigator

import (
	"log"

	"github.com/brianshef/knavigator/internal/character"
	"github.com/brianshef/knavigator/internal/data"
)

// Hello simply prints "hello"
func GenerateCharacter() {
	config, err := data.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	c := character.NewCharacter("John Smith", config)
	c.Print()
}
