package rules

import (
	"goa.design/goa/v3/eval"
)

type Rule interface {
	Name() string
	IsDisabled() bool
	Apply(roots []eval.Root) error
}
