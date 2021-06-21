package data

import "log"

const configPath = "./configs/"

var configFiles = []string{
	"traits.json",
	"names.json",
	"armor.json",
	"weapons.json",
	"gear.json",
}

// Mapping is a commonly used data structure representing map[string]interface{}
type Mapping map[string]interface{}

// Config is the configuration superstructure for Knave, holding all the options for the game
type Config struct {
	Traits  *TraitsConfig
	Names   *NamesConfig
	Armor   *ArmorConfig
	Weapons *WeaponsConfig
	Gear    *GearConfig
}

// Configurable defines an interface for any configuration struct that can be converted from a map
type Configurable interface {
	mapFrom(Mapping) error
}

// NewConfig generates a new configuration struct
func NewConfig() (c *Config, err error) {
	c = &Config{
		Traits:  &TraitsConfig{},
		Names:   &NamesConfig{},
		Armor:   &ArmorConfig{},
		Weapons: &WeaponsConfig{},
		Gear: &GearConfig{},
	}

	err = c.Load()

	return
}

func (c *Config) mapFrom(mappings map[string]Mapping) (err error) {
	for filename, m := range mappings {
		switch filename {
		case "traits.json":
			err = c.Traits.mapFrom(m)
		case "names.json":
			err = c.Names.mapFrom(m)
		case "armor.json":
			err = c.Armor.mapFrom(m)
		case "weapons.json":
			err = c.Weapons.mapFrom(m)
		case "gear.json":
			err = c.Gear.mapFrom(m)
		default:
			log.Printf("WARNING: skipping unknown config filename %s", filename)
			continue
		}

		if err != nil {
			log.Fatalf("error mapping %s to configuration", filename)
			return
		}
	}

	return
}
