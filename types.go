package lint

import (
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
)

type WordCase = casing.WordCase

const (
	CamelCase    = casing.CamelCase
	PascalCase   = casing.PascalCase
	SnakeCase    = casing.SnakeCase
	KebabCase    = casing.KebabCase
	GoCamelCase  = casing.GoCamelCase
	GoPascalCase = casing.GoPascalCase
	GoSnakeCase  = casing.GoSnakeCase
	GoKebabCase  = casing.GoKebabCase
)

type Initialisms = casing.Initialisms
