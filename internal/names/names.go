package names

import (
	"fmt"
	"math/rand"
)

func random(choices []string) string {
	return choices[rand.Intn((len(choices)))]
}

// GenerateName generates a random firstname+surname combo based on separate lists of options
func GenerateName(firstNames, surnames []string) string {
	return fmt.Sprintf("%s %s", random(firstNames), random(surnames))
}
