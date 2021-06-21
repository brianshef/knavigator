package inventory

import (
	"fmt"
)

// Item represents any item a Knave character can hold, equip, or use
type Item struct {
	Name    string
	Slots   int
	Quality int
}

// String is a method which returns a string representation of an item
func (i *Item) String() string {
	return fmt.Sprintf("%s (%v slots, quality %v)", i.Name, i.Slots, i.Quality)
}
