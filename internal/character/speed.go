package character

import "fmt"

const (
	explorationSpeed int = 120
	combatSpeed      int = 40
)

type speed struct {
	exploration int
	combat      int
}

// String is a method to return a string representation of the speed
func (s *speed) String() string {
	return fmt.Sprintf(
		"Exploration %d ft. / Combat %d ft.",
		s.exploration,
		s.combat,
	)
}

func generateSpeed() *speed {
	return &speed{exploration: explorationSpeed, combat: combatSpeed}
}
