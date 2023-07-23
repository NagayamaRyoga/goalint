package config

import (
	"os"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
)

type Config struct {
	Disabled bool
	Debug    bool

	MethodCasingConvention *MethodCasingConventionConfig
	TypeCasingConvention   *TypeCasingConventionConfig
}

type caseConfig struct {
	Disabled    bool
	WordCase    casing.WordCase
	Initialisms casing.Initialisms
}

type MethodCasingConventionConfig struct {
	caseConfig
}

type TypeCasingConventionConfig struct {
	caseConfig
}

func NewConfig() *Config {
	debug := false
	if v, ok := os.LookupEnv("GOA_LINT_DEBUG"); ok && len(v) > 0 {
		debug = true
	}

	return &Config{
		Disabled: false,
		Debug:    debug,
		MethodCasingConvention: &MethodCasingConventionConfig{
			caseConfig: caseConfig{
				Disabled:    false,
				WordCase:    casing.SnakeCase,
				Initialisms: nil,
			},
		},
		TypeCasingConvention: &TypeCasingConventionConfig{
			caseConfig: caseConfig{
				Disabled:    false,
				WordCase:    casing.GoPascalCase,
				Initialisms: nil,
			},
		},
	}
}
