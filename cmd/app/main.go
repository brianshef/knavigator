package main

import (
	"log"

	"github.com/brianshef/knavigator/internal/knavigator"
)

func main() {
	c, err := knavigator.GenerateCharacter()
	if err != nil {
		log.Fatal(err)
	}
	c.Print()
}
