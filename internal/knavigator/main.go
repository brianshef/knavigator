package knavigator

import "github.com/brianshef/knavigator/internal/character"

// Hello simply prints "hello"
func GenerateCharacter() {
	c := character.NewCharacter("John Smith")
	c.Print()
}