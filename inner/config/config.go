package config

import (
	"os"

	"github.com/NagayamaRyoga/goalint/inner/rules/http_path_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/http_path_segment_validation"
	"github.com/NagayamaRyoga/goalint/inner/rules/method_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/result_type_identifier_naming_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_description_exists"
)

type Config struct {
	Disabled bool
	Debug    bool

	MethodCasingConvention               *method_casing_convention.Config
	TypeCasingConvention                 *type_casing_convention.Config
	TypeDescriptionExists                *type_description_exists.Config
	ResultTypeIdentifierNamingConvention *result_type_identifier_naming_convention.Config
	HTTPPathCasingConvention             *http_path_casing_convention.Config
	HTTPPathSegmentValidation            *http_path_segment_validation.Config
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

		MethodCasingConvention:               method_casing_convention.NewConfig(),
		TypeCasingConvention:                 type_casing_convention.NewConfig(),
		TypeDescriptionExists:                type_description_exists.NewConfig(),
		ResultTypeIdentifierNamingConvention: result_type_identifier_naming_convention.NewConfig(),
		HTTPPathCasingConvention:             http_path_casing_convention.NewConfig(),
		HTTPPathSegmentValidation:            http_path_segment_validation.NewConfig(),
	}
}
