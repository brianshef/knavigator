package data

import (
	"encoding/json"
)

// TraitsConfig defines the data structure holding all the traits options
type TraitsConfig struct {
	Physique    []string `json:"physique"`
	Face        []string `json:"face"`
	Skin        []string `json:"skin"`
	Hair        []string `json:"hair"`
	Clothing    []string `json:"clothing"`
	Virtue      []string `json:"virtue"`
	Vice        []string `json:"vice"`
	Speech      []string `json:"speech"`
	Background  []string `json:"background"`
	Misfortunes []string `json:"misfortunes"`
	Alignment   []string `json:"alignment"`
}

func (c *TraitsConfig) mapFrom(m Mapping) (err error) {
	jsonString, err := json.Marshal(m)

	if err != nil {
		return
	}

	err = json.Unmarshal(jsonString, &c)

	return
}
