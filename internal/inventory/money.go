package inventory

func generateStartingMoney() int {
	results := dice.D6.Roll(3)
	sumResult := 0
	for _, r := range results {
		sumResult += r
	}
	return sumResult * 20
}
