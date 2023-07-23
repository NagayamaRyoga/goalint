package rules

import (
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
	"goa.design/goa/v3/eval"
)

type Rule interface {
	Name() string
	Apply(c *config.Config, roots []eval.Root) error
}

var Rules = []Rule{}
