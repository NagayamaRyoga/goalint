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

func newRules(logger *log.Logger, cfg *config.Config) []rules.Rule {
	return []rules.Rule{
		method_casing_convention.NewRule(logger, cfg.MethodCasingConvention),
		type_casing_convention.NewRule(logger, cfg.TypeCasingConvention),
		type_description_exists.NewRule(logger, cfg.TypeDescriptionExists),
		http_path_casing_convention.NewRule(logger, cfg.HTTPPathCasingConvention),
		http_path_segment_validation.NewRule(logger, cfg.HTTPPathSegmentValidation),
	}
}

func Run(cfg *config.Config, genpkg string, roots []eval.Root) error {
	var merr error

	out := io.Discard
	if cfg.Debug {
		out = os.Stderr
	}

	logger := log.New(out, "[goa-lint] ", log.Ltime)

	logger.Println("genpkg:", genpkg)

	for _, rule := range newRules(logger, cfg) {
		if rule.IsDisabled() {
			continue
		}

		logger.Println("rule:", rule.Name())

		if err := rule.Apply(roots); err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	return merr
}
