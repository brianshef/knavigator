package data

import "encoding/json"

// GearConfig defines the data structure holding all the gear options
type GearConfig struct {
	DungeoneeringGear []string `json:"dungeoneering_gear"`
	GeneralGear1      []string `json:"general_gear_1"`
	GeneralGear2      []string `json:"general_gear_2"`
}

func (c *GearConfig) mapFrom(m Mapping) (err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonString, &c)

	return
}
