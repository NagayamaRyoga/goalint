package type_attribute_example_exists

import (
	"github.com/NagayamaRyoga/goalint/pkg/common/types"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"goa.design/goa/v3/expr"
)

type Config struct {
	Disabled      bool
	Level         reports.ErrorLevel
	RequiredTypes types.DataTypeList
}

func NewConfig() *Config {
	return &Config{
		Disabled: false,
		Level:    reports.ErrorLevelError,
		RequiredTypes: types.DataTypeList{
			expr.Boolean,
			expr.Int,
			expr.Int32,
			expr.Int64,
			expr.UInt,
			expr.UInt32,
			expr.UInt64,
			expr.Float32,
			expr.Float64,
			expr.String,
			expr.Bytes,
			expr.Any,
		},
	}
}
