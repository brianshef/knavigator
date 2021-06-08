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

// Sprintf is a method to return a string representation of the speed
func (s *speed) Sprintf() string {
	return fmt.Sprintf(
		"Speed: Exploration %d ft. / Combat %d ft.",
		s.exploration,
		s.combat,
	)
}

func generateSpeed() *speed {
	return &speed{exploration: explorationSpeed, combat: combatSpeed}
}
