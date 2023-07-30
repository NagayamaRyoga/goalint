package type_casing_convention

import (
	"github.com/NagayamaRyoga/goalint/inner/common/casing"
	"github.com/NagayamaRyoga/goalint/inner/reports"
)

type Config struct {
	Disabled    bool
	Level       reports.ErrorLevel
	WordCase    casing.WordCase
	Initialisms casing.Initialisms
}

func NewConfig() *Config {
	return &Config{
		Disabled:    false,
		Level:       reports.ErrorLevelError,
		WordCase:    casing.GoPascalCase,
		Initialisms: nil,
	}
}
