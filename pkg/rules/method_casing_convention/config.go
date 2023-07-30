package method_casing_convention

import (
	"github.com/NagayamaRyoga/goalint/pkg/common/casing"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
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
		WordCase:    casing.SnakeCase,
		Initialisms: nil,
	}
}
