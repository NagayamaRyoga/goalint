package runner

import (
	"io"
	"log"
	"os"

	"github.com/NagayamaRyoga/goalint/inner/config"
	"github.com/NagayamaRyoga/goalint/inner/rules"
	"github.com/NagayamaRyoga/goalint/inner/rules/http_path_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/http_path_segment_validation"
	"github.com/NagayamaRyoga/goalint/inner/rules/method_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_description_exists"
	"github.com/hashicorp/go-multierror"
	"goa.design/goa/v3/eval"
)

var ruleSet = []rules.NewRule{
	method_casing_convention.NewRule,
	type_casing_convention.NewRule,
	type_description_exists.NewRule,
	http_path_casing_convention.NewRule,
	http_path_segment_validation.NewRule,
}

func Run(cfg *config.Config, genpkg string, roots []eval.Root) error {
	var merr error

	out := io.Discard
	if cfg.Debug {
		out = os.Stderr
	}

	logger := log.New(out, "[goa-lint] ", log.Ltime)

	logger.Println("genpkg:", genpkg)

	for _, rule := range ruleSet {
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