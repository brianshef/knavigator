package data

import (
	"encoding/json"
)

// ArmorConfig defines the data structure holding all the armor options
type ArmorConfig struct {
	Armor             []string `json:"armor"`
	HelmetsAndShields []string `json:"helmets_and_shields"`
}

func (c *ArmorConfig) mapFrom(m Mapping) (err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonString, &c)

	return
}
