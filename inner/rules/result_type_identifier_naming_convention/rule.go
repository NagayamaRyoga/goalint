package result_type_identifier_naming_convention

import (
	"fmt"
	"log"
	"strings"

	"github.com/NagayamaRyoga/goalint/inner/common/walk"
	"github.com/NagayamaRyoga/goalint/inner/reports"
	"github.com/NagayamaRyoga/goalint/inner/rules"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	logger *log.Logger
	cfg    *Config
}

func NewRule(logger *log.Logger, cfg *Config) *Rule {
	return &Rule{
		logger: logger,
		cfg:    cfg,
	}
}

func (r *Rule) Name() string {
	return "ResultTypeIdentifierNamingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.WalkResultType)
}

func (r *Rule) WalkResultType(t expr.UserType) (rl reports.ReportList) {
	const resultTypeIDPrefix = "application/vnd."

	if t.Kind() == expr.ResultTypeKind && !strings.HasPrefix(t.ID(), resultTypeIDPrefix) {
		rl = append(rl, reports.NewReport(
			r.cfg.Level,
			r.Name(),
			fmt.Sprintf("ResultType(%q)", t.ID()),
			"ResultType identifier should have prefix '%s'", resultTypeIDPrefix,
		))
	}

	return
}
