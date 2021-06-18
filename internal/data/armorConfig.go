package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

// ArmorConfig defines the data structure holding all the armor options
type ArmorConfig struct {
	Armor             []string `json:"armor"`
	HelmetsAndShields []string `json:"helmets_and_shields"`
}

func loadArmor(path string) (tc *map[string]interface{}, err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}

	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &tc)
	if err != nil {
		return
	}
	if err != nil {
		log.Fatalf("%s: %s", absPath, err)
		return
	}

	return
}

func mapToArmorConfig(m map[string]interface{}) (c *ArmorConfig, err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	c = new(ArmorConfig)
	err = json.Unmarshal(jsonString, &c)

	return
}
