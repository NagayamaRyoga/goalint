package result_type_identifier_naming_convention

import (
	"fmt"
	"log"
	"strings"

	"github.com/NagayamaRyoga/goalint/pkg/common/walk"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"github.com/NagayamaRyoga/goalint/pkg/rules"
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
	return walk.Type(roots, r.WalkType)
}

func (r *Rule) WalkType(t expr.DataType) (rl reports.ReportList) {
	const resultTypeIDPrefix = "application/vnd."

	if t, ok := t.(*expr.ResultTypeExpr); ok {
		if !strings.HasPrefix(t.ID(), resultTypeIDPrefix) {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				fmt.Sprintf("ResultType(%q)", t.ID()),
				"ResultType identifier should have prefix '%s'", resultTypeIDPrefix,
			))
		}
	}

	return
}
