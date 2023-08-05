package runner

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/NagayamaRyoga/goalint/pkg/config"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"github.com/NagayamaRyoga/goalint/pkg/rules"
	"github.com/NagayamaRyoga/goalint/pkg/rules/api_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/api_title_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/http_path_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/http_path_segment_validation"
	"github.com/NagayamaRyoga/goalint/pkg/rules/method_array_result"
	"github.com/NagayamaRyoga/goalint/pkg/rules/method_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/method_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/no_unnamed_method_payload_type"
	"github.com/NagayamaRyoga/goalint/pkg/rules/no_unnamed_method_result_type"
	"github.com/NagayamaRyoga/goalint/pkg/rules/result_type_identifier_naming_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/service_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_description_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_example_exists"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_casing_convention"
	"github.com/NagayamaRyoga/goalint/pkg/rules/type_description_exists"
	"goa.design/goa/v3/eval"
)

var ErrFailed error = errors.New("goalint failed")

func newRules(logger *log.Logger, cfg *config.Config) []rules.Rule {
	return []rules.Rule{
		api_title_exists.NewRule(logger, cfg.APITitleExists),
		api_description_exists.NewRule(logger, cfg.APIDescriptionExists),
		service_description_exists.NewRule(logger, cfg.ServiceDescriptionExists),
		method_casing_convention.NewRule(logger, cfg.MethodCasingConvention),
		method_description_exists.NewRule(logger, cfg.MethodDescriptionExists),
		method_array_result.NewRule(logger, cfg.MethodArrayResult),
		no_unnamed_method_payload_type.NewRule(logger, cfg.NoUnnamedMethodPayloadType),
		no_unnamed_method_result_type.NewRule(logger, cfg.NoUnnamedMethodResultType),
		type_casing_convention.NewRule(logger, cfg.TypeCasingConvention),
		type_description_exists.NewRule(logger, cfg.TypeDescriptionExists),
		type_attribute_casing_convention.NewRule(logger, cfg.TypeAttributeCasingConvention),
		type_attribute_description_exists.NewRule(logger, cfg.TypeAttributeDescriptionExists),
		type_attribute_example_exists.NewRule(logger, cfg.TypeAttributeExampleExists),
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
