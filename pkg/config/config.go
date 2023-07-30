package config

import (
	"os"

	"github.com/NagayamaRyoga/goalint/pkg/rules/http_path_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/http_path_segment_validation"
	"github.com/NagayamaRyoga/goalint/pkg/rules/method_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/method_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/no_unnamed_method_payload_type"
	"github.com/NagayamaRyoga/goalint/pkg/rules/result_type_identifier_naming_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/service_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_description_exists"
)

type Config struct {
	Disabled bool
	Debug    bool

	ServiceDescriptionExists             *service_description_exists.Config
	MethodCasingConvention               *method_casing_convention.Config
	MethodDescriptionExists              *method_description_exists.Config
	NoUnnamedMethodPayloadType           *no_unnamed_method_payload_type.Config
	TypeCasingConvention                 *type_casing_convention.Config
	TypeDescriptionExists                *type_description_exists.Config
	TypeAttributeCasingConvention        *type_attribute_casing_convention.Config
	TypeAttributeDescriptionExists       *type_attribute_description_exists.Config
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

		ServiceDescriptionExists:             service_description_exists.NewConfig(),
		MethodCasingConvention:               method_casing_convention.NewConfig(),
		MethodDescriptionExists:              method_description_exists.NewConfig(),
		NoUnnamedMethodPayloadType:           no_unnamed_method_payload_type.NewConfig(),
		TypeCasingConvention:                 type_casing_convention.NewConfig(),
		TypeDescriptionExists:                type_description_exists.NewConfig(),
		TypeAttributeCasingConvention:        type_attribute_casing_convention.NewConfig(),
		TypeAttributeDescriptionExists:       type_attribute_description_exists.NewConfig(),
		ResultTypeIdentifierNamingConvention: result_type_identifier_naming_convention.NewConfig(),
		HTTPPathCasingConvention:             http_path_casing_convention.NewConfig(),
		HTTPPathSegmentValidation:            http_path_segment_validation.NewConfig(),
	}
}
