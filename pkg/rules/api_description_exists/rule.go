package api_description_exists

import (
	"fmt"
	"log"

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
	return "ApiDescriptionExists"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Expression(roots, r.WalkMethodExpr)
}

func (r *Rule) WalkMethodExpr(e eval.Expression) (rl reports.ReportList) {
	if e, ok := e.(*expr.APIExpr); ok {
		if len(e.Description) == 0 {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				fmt.Sprintf("API(%q)", e.Name),
				"API should have non-empty description",
			))
		}
	}

	return
}
