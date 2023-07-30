package method_casing_convention

import (
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/casing"
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
	caser  *casing.Caser
}

func NewRule(logger *log.Logger, cfg *Config) *Rule {
	return &Rule{
		logger: logger,
		cfg:    cfg,
		caser:  casing.NewCaser(cfg.WordCase, cfg.Initialisms),
	}
}

func (r *Rule) Name() string {
	return "MethodCasingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Expression(roots, r.WalkMethodExpr)
}

func (r *Rule) WalkMethodExpr(e eval.Expression) (rl reports.ReportList) {
	if e, ok := e.(*expr.MethodExpr); ok {
		if !r.caser.Check(e.Name) {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				e.EvalName(),
				"Method name %q should be %s (%q)", e.Name, r.cfg.WordCase, r.caser.To(e.Name),
			))
		}
	}

	return
}
