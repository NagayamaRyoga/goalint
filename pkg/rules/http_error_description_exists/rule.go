package http_error_description_exists

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
	return "HTTPErrorDescriptionExists"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Expression(roots, r.WalkHTTPEndpointExpr)
}

func (r *Rule) WalkHTTPEndpointExpr(e eval.Expression) (rl reports.ReportList) {
	if e, ok := e.(*expr.HTTPEndpointExpr); ok {
		for _, httpError := range e.HTTPErrors {
			if len(httpError.Response.Description) == 0 {
				rl = append(rl, reports.NewReport(
					r.cfg.Level,
					r.Name(),
					httpError.EvalName(),
					"Error response %q should have non-empty description", httpError.Name,
				))
			}
		}
	}

	return
}
