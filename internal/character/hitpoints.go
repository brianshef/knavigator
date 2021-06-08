package character

import "fmt"

type hitpoints struct {
	maximum     int
	healingRate int // 1d8 + CON bonus
}

// String is a method for returning a string representation of the hitpoints
func (h *hitpoints) String() string {
	return fmt.Sprintf("HP: %d / Healing Rate: %d", h.maximum, h.healingRate)
}

func generateHitPoints(conBonus int) *hitpoints {
	max := dice.D8.Roll(1)[0]
	hr := dice.D8.Roll(1)[0] + conBonus

	return &hitpoints{maximum: max, healingRate: hr}
}
