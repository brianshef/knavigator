package dice

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once

var d4 = []int{1, 2, 3, 4}
var d6 = []int{1, 2, 3, 4, 5, 6}
var d8 = []int{1, 2, 3, 4, 5, 6, 7, 8}
var d10 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var d20 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// A Die is a specific set of possible integers
type Die struct {
	values []int
}

// Dice defines the types of dice available
type Dice struct {
	D4  Die
	D6  Die
	D8  Die
	D10 Die
	D20 Die
}

// NewDice instantiates a set of dice and their values that can be rolled
func NewDice() *Dice {
	return &Dice{
		D4:  Die{values: d4},
		D6:  Die{values: d6},
		D8:  Die{values: d8},
		D10: Die{values: d10},
		D20: Die{values: d20},
	}
}

// lowest is a utility function to find the lowest int in a set
func lowest(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// RollOnce returns a single die roll, rolling a die with d sides n times.
func (d *Die) RollOnce() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return d.values[rand.Intn(len(d.values))]
}

// Roll returns a set of dice rolls, rolling a die with d sides n times.
func (d *Die) Roll(n int) []int {
	var result []int

	for i := 0; i < n; i++ {
		result = append(result, d.RollOnce())
	}

	return result
}

// RollKeepLowest rolls n dice, and keeps the lowest
func (d *Die) RollKeepLowest(n int) int {
	return lowest(d.Roll(n))
}

// String returns a string representation (the name) of a die
func (d *Die) String() string {
	return fmt.Sprintf("d%v", len(d.values))
}
