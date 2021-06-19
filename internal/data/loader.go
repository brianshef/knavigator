package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

func loadJSON(path string) (m *Mapping, err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}

	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(f, &m)
	if err != nil {
		return
	}
	if err != nil {
		log.Fatalf("%s: %s", absPath, err)
		return
	}

	return
}

// Load loads in all the config first as a mapping, and then converts to the requisite struct
func (c *Config) Load() error {
	mappings := make(map[string]Mapping)
	for _, filename := range configFiles {
		m, err := loadJSON(configPath + filename)
		if err != nil {
			log.Fatalf("%s: %s", filename, err)
			return err
		}

		mappings[filename] = *m
	}

	err := c.mapFrom(mappings)
	if err != nil {
		return err
	}

	// Sanity checking
	if c.Traits == nil || c.Armor == nil || c.Names == nil || c.Weapons == nil {
		return errors.New("configuration is missing after loading")
	}

	return nil
}
