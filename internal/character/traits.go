package character

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/brianshef/knavigator/internal/data"
)

type traits struct {
	Physique    string `json:"physique"`
	Face        string `json:"face"`
	Skin        string `json:"skin"`
	Hair        string `json:"hair"`
	Clothing    string `json:"clothing"`
	Virtue      string `json:"virtue"`
	Vice        string `json:"vice"`
	Speech      string `json:"speech"`
	Background  string `json:"background"`
	Misfortunes string `json:"misfortunes"`
	Alignment   string `json:"alignment"`
}

func generateTraits(config *data.TraitsConfig) *traits {
	return &traits{
		Physique:    strings.ToUpper(config.Physique[dice.D20.RollOnce()-1]),
		Face:        strings.ToUpper(config.Face[dice.D20.RollOnce()-1]),
		Skin:        strings.ToUpper(config.Skin[dice.D20.RollOnce()-1]),
		Hair:        strings.ToUpper(config.Hair[dice.D20.RollOnce()-1]),
		Clothing:    strings.ToUpper(config.Clothing[dice.D20.RollOnce()-1]),
		Virtue:      strings.ToUpper(config.Virtue[dice.D20.RollOnce()-1]),
		Vice:        strings.ToUpper(config.Vice[dice.D20.RollOnce()-1]),
		Speech:      strings.ToUpper(config.Speech[dice.D20.RollOnce()-1]),
		Background:  strings.ToUpper(config.Background[dice.D20.RollOnce()-1]),
		Misfortunes: strings.ToUpper(config.Misfortunes[dice.D20.RollOnce()-1]),
		Alignment:   strings.ToUpper(config.Alignment[dice.D20.RollOnce()-1]),
	}
}

// String returns the string representation of the traits
func (t *traits) String() string {
	j, _ := json.Marshal(t)
	return string(j)
}

// DescriptiveString returns a human-friendly description based on all the traits
func (t *traits) DescriptiveString() string {
	return fmt.Sprintf(
		"Has a %s physique with a %s face, %s skin, and %s hair. "+
			"Wears %s clothing and speaks in a %s voice. "+
			"Is %s, yet %s. Aligned towards %s. "+
			"Was once a %s, and has had the misfortune of being %s.",
		t.Physique, t.Face, t.Skin, t.Hair,
		t.Clothing, t.Speech,
		t.Virtue, t.Vice, t.Alignment,
		t.Background, t.Misfortunes,
	)
}
