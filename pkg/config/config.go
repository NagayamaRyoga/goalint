package config

import (
	"os"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
)

type Config struct {
	Disabled bool
	Debug    bool

	MethodCasingConvention *MethodCasingConvention
	TypeCasingConvention   *TypeCasingConvention
	TypeDescriptionExists  *TypeDescriptionExists
}

type casingConfig struct {
	Disabled    bool
	WordCase    casing.WordCase
	Initialisms casing.Initialisms
}

type MethodCasingConvention struct {
	casingConfig
}

type TypeCasingConvention struct {
	casingConfig
}

type TypeDescriptionExists struct {
	Disabled bool
}

func NewConfig() *Config {
	disabled := false
	if v, ok := os.LookupEnv("GOA_LINT_DISABLED"); ok && len(v) > 0 {
		disabled = true
	}

	debug := false
	if v, ok := os.LookupEnv("GOA_LINT_DEBUG"); ok && len(v) > 0 {
		debug = true
	}

	return &Config{
		Disabled: disabled,
		Debug:    debug,
		MethodCasingConvention: &MethodCasingConvention{
			casingConfig: casingConfig{
				Disabled:    false,
				WordCase:    casing.SnakeCase,
				Initialisms: nil,
			},
		},
		TypeCasingConvention: &TypeCasingConvention{
			casingConfig: casingConfig{
				Disabled:    false,
				WordCase:    casing.GoPascalCase,
				Initialisms: nil,
			},
		},
		TypeDescriptionExists: &TypeDescriptionExists{
			Disabled: false,
		},
	}
}
