package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

const configPath = "./configs/"

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

// Config is the configuration superstructure for Knave, holding all the options for the game
type Config struct {
	Traits *TraitsConfig
}

func loadTraits() (tc *TraitsConfig, err error) {
	absPath, err := filepath.Abs(configPath + "traits.json")
	if err != nil {
		return
	}

	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		return
	}

	tc = &TraitsConfig{
		Physique:    []string{},
		Face:        []string{},
		Skin:        []string{},
		Hair:        []string{},
		Clothing:    []string{},
		Virtue:      []string{},
		Vice:        []string{},
		Speech:      []string{},
		Background:  []string{},
		Misfortunes: []string{},
		Alignment:   []string{},
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

// LoadConfig reads in the configuration files and returns the config superstructure
func LoadConfig() (c *Config, err error) {
	c = &Config{}

	tc, err := loadTraits()
	if err != nil {
		return
	}
	c.Traits = tc

	return
}
