package type_attribute_example_exists

import (
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"goa.design/goa/v3/expr"
)

type Config struct {
	Disabled     bool
	Level        reports.ErrorLevel
	AllowedTypes []expr.Kind
}

func NewConfig() *Config {
	return &Config{
		Disabled: false,
		Level:    reports.ErrorLevelError,
		AllowedTypes: []expr.Kind{
			expr.ArrayKind,
			expr.ObjectKind,
			expr.UserTypeKind,
			expr.ResultTypeKind,
		},
	}
}
