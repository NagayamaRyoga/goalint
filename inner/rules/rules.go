package rules

import (
	"log"

	"github.com/NagayamaRyoga/goalint/inner/config"
	"goa.design/goa/v3/eval"
)

type NewRule func(logger *log.Logger, cfg *config.Config) Rule

type Rule interface {
	Name() string
	IsDisabled() bool
	Apply(roots []eval.Root) error
}
