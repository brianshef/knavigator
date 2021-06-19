package data

import (
	"encoding/json"
)

// WeaponsConfig defines the data structure holding all the weapons options
type WeaponsConfig struct {
	Weapons []string `json:"weapons"`
}

func (c *WeaponsConfig) mapFrom(m Mapping) (err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonString, &c)

	return
}
