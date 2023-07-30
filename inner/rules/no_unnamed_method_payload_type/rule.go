package no_unnamed_method_payload_type

import (
	"log"

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
	return "NoUnnamedMethodPayloadType"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Expression(roots, r.WalkMethodExpr)
}

func (r *Rule) WalkMethodExpr(e eval.Expression) (rl reports.ReportList) {
	if e, ok := e.(*expr.MethodExpr); ok {
		if e.Payload.Type.Kind() == expr.ObjectKind {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				e.EvalName(),
				"Method payload should be an user defined type",
			))
		}
	}

	return
}
