package runner

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/NagayamaRyoga/goalint/inner/config"
	"github.com/NagayamaRyoga/goalint/inner/reports"
	"github.com/NagayamaRyoga/goalint/inner/rules"
	"github.com/NagayamaRyoga/goalint/inner/rules/http_path_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/http_path_segment_validation"
	"github.com/NagayamaRyoga/goalint/inner/rules/method_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/result_type_identifier_naming_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_attribute_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_casing_convention"
	"github.com/NagayamaRyoga/goalint/inner/rules/type_description_exists"
	"goa.design/goa/v3/eval"
)

var ErrFailed error = errors.New("goalint failed")

func newRules(logger *log.Logger, cfg *config.Config) []rules.Rule {
	return []rules.Rule{
		method_casing_convention.NewRule(logger, cfg.MethodCasingConvention),
		type_casing_convention.NewRule(logger, cfg.TypeCasingConvention),
		type_description_exists.NewRule(logger, cfg.TypeDescriptionExists),
		type_attribute_casing_convention.NewRule(logger, cfg.TypeAttributeCasingConvention),
		result_type_identifier_naming_convention.NewRule(logger, cfg.ResultTypeIdentifierNamingConvention),
		http_path_casing_convention.NewRule(logger, cfg.HTTPPathCasingConvention),
		http_path_segment_validation.NewRule(logger, cfg.HTTPPathSegmentValidation),
	}
}

func Run(cfg *config.Config, genpkg string, roots []eval.Root) error {
	out := io.Discard
	if cfg.Debug {
		out = os.Stderr
	}

	logger := log.New(out, "[goa-lint] ", log.Ltime)

	logger.Println("genpkg:", genpkg)

	var reports reports.ReportList

	for _, rule := range newRules(logger, cfg) {
		if rule.IsDisabled() {
			continue
		}

		logger.Println("rule:", rule.Name())

		reports = append(reports, rule.Apply(roots)...)
	}

	fmt.Fprint(os.Stderr, reports)

	if reports.CountErrors() > 0 {
		return ErrFailed
	}

	return nil
}
