package method_array_result

import (
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
	return "MethodArrayResult"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Expression(roots, r.WalkMethodExpr)
}

func (r *Rule) WalkMethodExpr(e eval.Expression) (rl reports.ReportList) {
	if e, ok := e.(*expr.MethodExpr); ok {
		switch t := e.Result.Type.(type) {
		case *expr.Array:
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				e.EvalName(),
				"Method should not return an array as the top-level structure of a response",
			))
		case *expr.ResultTypeExpr:
			if _, ok := t.AttributeExpr.Type.(*expr.Array); ok {
				rl = append(rl, reports.NewReport(
					r.cfg.Level,
					r.Name(),
					e.EvalName(),
					"Method should not return an array as the top-level structure of a response",
				))
			}
		}
	}

	return
}
