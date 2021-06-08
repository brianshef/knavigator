package character

type armor struct {
	item
	helmet  bool
	shield  bool
	defense int
}

func generateArmor() *armor {
	return &armor{
		item:    item{name: "none", slots: 0},
		helmet:  false,
		shield:  false,
		defense: 11,
	}
}
