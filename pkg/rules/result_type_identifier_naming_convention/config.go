package result_type_identifier_naming_convention

import (
	"github.com/NagayamaRyoga/goalint/pkg/reports"
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
