package dice

import (
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once

var d6 = []int{1, 2, 3, 4, 5, 6}
var d8 = []int{1, 2, 3, 4, 5, 6, 7, 8}
var d20 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

type die struct {
	values	[]int
}

// Dice defines the types of dice available
type Dice struct {
	// D6 is the 6-sided die
	D6	die
	// D8 is the 8-sided die
	D8	die
	// D20 is the 20-sided die
	D20	die
}

// NewDice instantiates a set of dice and their values that can be rolled
func NewDice() *Dice {
	return &Dice{
		D6: die{values: d6,},
		D8: die{values: d8,},
		D20: die{values: d20,},
	}
}

// lowest is a utility function to find the lowest int in a set
func lowest(values []int) int {
	min := values[0]
	for _, v := range values {
		if (v < min) {
			min = v
		}
	}
	return min
}

func (d *die) rollOnce() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return d.values[rand.Intn(len(d.values))]
}

// Roll returns a set of dice rolls, rolling a die with d sides n times.
func (d *die) Roll(n int) []int {
	var result []int

	for i := 0; i < n; i++ {
		result = append(result, d.rollOnce())
	}
	
	return result
}

// RollKeepLowest rolls n dice, and keeps the lowest
func (d *die) RollKeepLowest(n int) int {
	return lowest(d.Roll(n))
}