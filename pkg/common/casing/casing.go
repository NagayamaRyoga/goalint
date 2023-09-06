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

type Initialisms []string

func (i *Initialisms) Add(s ...string) *Initialisms {
	*i = append(*i, s...)

	return i
}

type Caser struct {
	impl     *strcase.Caser
	wordCase WordCase
}

func NewCaser(wordCase WordCase, initialisms Initialisms) *Caser {
	var goInitialisms bool
	switch wordCase {
	case CamelCase, PascalCase, SnakeCase, KebabCase:
		goInitialisms = false
	case GoCamelCase, GoPascalCase, GoSnakeCase, GoKebabCase:
		goInitialisms = true
	}

	initialismOverrides := make(map[string]bool, len(initialisms))
	for _, ini := range initialisms {
		initialismOverrides[ini] = true
	}

	return &Caser{
		impl:     strcase.NewCaser(goInitialisms, initialismOverrides, nil),
		wordCase: wordCase,
	}
}

func (c *Caser) To(s string) string {
	switch c.wordCase {
	case CamelCase, GoCamelCase:
		return c.impl.ToCamel(s)
	case PascalCase, GoPascalCase:
		return c.impl.ToPascal(s)
	case SnakeCase, GoSnakeCase:
		return c.impl.ToSnake(s)
	case KebabCase, GoKebabCase:
		return c.impl.ToKebab(s)
	default:
		fmt.Fprintf(os.Stderr, "unknown word case: %s\n", c.wordCase)

		return ""
	}
}

func (c *Caser) Check(s string) bool {
	return c.To(s) == s
}
