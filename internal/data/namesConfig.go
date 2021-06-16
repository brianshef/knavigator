package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

// NamesConfig defines the data structure holding all the names options
type NamesConfig struct {
	FirstNames []string `json:"first_names"`
	Surnames   []string `json:"surnames"`
}

func loadNames(path string) (nc *map[string]interface{}, err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}

	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &nc)
	if err != nil {
		return
	}
	if err != nil {
		log.Fatalf("%s: %s", absPath, err)
		return
	}

	return
}

func mapToNamesConfig(m map[string]interface{}) (c *NamesConfig, err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return
	}

	c = new(NamesConfig)
	err = json.Unmarshal(jsonString, &c)

	return
}
