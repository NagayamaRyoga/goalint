package lint

import (
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
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

type Config = config.Config
type ConfiguratorFunc func(*Config)
