package method_description_exists

import (
	"github.com/NagayamaRyoga/goalint/inner/reports"
)

type Config struct {
	Disabled bool
	Level    reports.ErrorLevel
}

func NewConfig() *Config {
	return &Config{
		Disabled: false,
		Level:    reports.ErrorLevelError,
	}
}
