package character

import (
	"encoding/json"
)

type traits struct {
	Physique    string `json:"physique"`
	Face        string `json:"face"`
	Skin        string `json:"skin"`
	Hair        string `json:"hair"`
	Clothing    string `json:"clothing"`
	Virtue      string `json:"virtue"`
	Dice        string `json:"dice"`
	Speech      string `json:"speech"`
	Background  string `json:"background"`
	Misfortunes string `json:"misfortunes"`
	Alignment   string `json:"alignment"`
}

func generateTraits() *traits {
	return &traits{}
}

// String returns the string representation of the traits
func (t *traits) String() string {
	j, _ := json.Marshal(t)
	return string(j)
}
