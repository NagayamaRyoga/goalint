package config

import (
	"os"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
)

type Config struct {
	Debug bool

	MethodCasingConvention *MethodCasingConventionConfig
}

type MethodCasingConventionConfig struct {
	Enabled  bool
	WordCase casing.WordCase
}

func NewConfig() *Config {
	debug := false
	if v, ok := os.LookupEnv("GOA_LINT_DEBUG"); ok && len(v) > 0 {
		debug = true
	}

	return &Config{
		Debug: debug,
		MethodCasingConvention: &MethodCasingConventionConfig{
			Enabled:  true,
			WordCase: casing.SnakeCase,
		},
	}
}
