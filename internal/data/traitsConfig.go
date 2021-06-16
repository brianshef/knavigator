package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
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

func loadTraits(path string) (tc *map[string]interface{}, err error) {
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

func mapToTraitsConfig(m map[string]interface{}) (c *TraitsConfig, err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	c = new(TraitsConfig)
	err = json.Unmarshal(jsonString, &c)

	return
}
