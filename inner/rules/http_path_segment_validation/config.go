package http_path_segment_validation

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
