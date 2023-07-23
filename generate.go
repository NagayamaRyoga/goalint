package lint

import (
	"io"
	"log"
	"os"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/rules"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/rules/method_casing_convention"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/rules/type_casing_convention"
	"github.com/hashicorp/go-multierror"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

var Rules = []rules.NewRule{
	method_casing_convention.NewRule,
	type_casing_convention.NewRule,
}

func init() {
	codegen.RegisterPlugin("github.com/NagayamaRyoga/goa-lint-plugin", "gen", Prepare, Generate)
}

func Prepare(genpkg string, roots []eval.Root) error {
	if cfg.Disabled {
		return nil
	}

	var merr error

	out := io.Discard
	if cfg.Debug {
		out = os.Stderr
	}

	logger := log.New(out, "[goa-lint] ", log.Ltime)

	logger.Println("genpkg:", genpkg)

	for _, rule := range Rules {
		r := rule(logger, cfg)
		if r.IsDisabled() {
			continue
		}

		logger.Println("rule:", r.Name())

		if err := r.Apply(roots); err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	return merr
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	return files, nil
}
