package data

import "log"

const (
	configPath           = "./configs/"
	traitsConfigFileName = "traits.json"
	namesConfigFileName  = "names.json"
	armorConfigFileName  = "armor.json"
)

// Config is the configuration superstructure for Knave, holding all the options for the game
type Config struct {
	Traits *TraitsConfig
	Names  *NamesConfig
	Armor  *ArmorConfig
}

// Defines a function which loads in a json file and returns a generic map
type loadFn func(string) (*map[string]interface{}, error)

// Defines the config files and their respective loading functions
var configMap = map[string]loadFn{
	traitsConfigFileName: loadTraits,
	namesConfigFileName:  loadNames,
	armorConfigFileName:  loadArmor,
}

// Load loads in all the config according to the configMap
func (c *Config) Load() (*Config, error) {
	for filename, loadFunc := range configMap {
		m, err := loadFunc(configPath + filename)
		if err != nil {
			log.Fatalf("%s: %s", filename, err)
			return nil, err
		}

		switch filename {
		case "traits.json":
			c.Traits, err = mapToTraitsConfig(*m)
		case "names.json":
			c.Names, err = mapToNamesConfig(*m)
		case "armor.json":
			c.Armor, err = mapToArmorConfig(*m)
		default:
			log.Printf("WARNING: skipping unknown config filename %s", filename)
			continue
		}

		if err != nil {
			log.Fatalf("error loading %s", filename)
			return nil, err
		}
	}

	return c, nil
}
