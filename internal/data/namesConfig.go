package data

import (
	"encoding/json"
)

// NamesConfig defines the data structure holding all the names options
type NamesConfig struct {
	FirstNames []string `json:"first_names"`
	Surnames   []string `json:"surnames"`
}

func (c *NamesConfig) mapFrom(m Mapping) (err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonString, &c)

	return
}
