package casing

import (
	"fmt"
	"os"

	"github.com/ettle/strcase"
)

type WordCase string

const (
	CamelCase    WordCase = "camelCase"
	PascalCase   WordCase = "PascalCase"
	SnakeCase    WordCase = "snake_case"
	KebabCase    WordCase = "kebab-case"
	GoCamelCase  WordCase = "goCamelCase"
	GoPascalCase WordCase = "GoPascalCase"
	GoSnakeCase  WordCase = "go_snake_case"
	GoKebabCase  WordCase = "go-kebab-case"
)

func To(s string, c WordCase) string {
	switch c {
	case CamelCase:
		return strcase.ToCamel(s)
	case PascalCase:
		return strcase.ToPascal(s)
	case SnakeCase:
		return strcase.ToSnake(s)
	case KebabCase:
		return strcase.ToKebab(s)
	case GoCamelCase:
		return strcase.ToGoCamel(s)
	case GoPascalCase:
		return strcase.ToGoPascal(s)
	case GoSnakeCase:
		return strcase.ToGoSnake(s)
	case GoKebabCase:
		return strcase.ToGoKebab(s)
	default:
		fmt.Fprintf(os.Stderr, "unknown word case: %s\n", c)

		return ""
	}
}

func Check(s string, c WordCase) bool {
	return To(s, c) == s
}
