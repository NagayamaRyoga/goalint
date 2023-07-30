package type_casing_convention

import (
	"github.com/NagayamaRyoga/goalint/inner/common/casing"
)

type Config struct {
	Disabled    bool
	WordCase    casing.WordCase
	Initialisms casing.Initialisms
}

func NewConfig() *Config {
	return &Config{
		Disabled:    false,
		WordCase:    casing.GoPascalCase,
		Initialisms: nil,
	}
}
