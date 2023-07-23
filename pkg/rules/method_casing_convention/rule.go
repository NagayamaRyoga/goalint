package method_casing_convention

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/walk"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/rules"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	walk.ExpressionWalker

	logger *log.Logger
	cfg    *config.MethodCasingConventionConfig
}

func NewRule(logger *log.Logger, cfg *config.Config) rules.Rule {
	return &Rule{
		ExpressionWalker: walk.NewBaseExpressionWalker(logger),

		logger: logger,
		cfg:    cfg.MethodCasingConvention,
	}
}

func (r *Rule) Name() string {
	return "MethodCasingConvention"
}

func (r *Rule) IsEnabled() bool {
	return r.cfg.Enabled
}

func (r *Rule) Apply(roots []eval.Root) error {
	return walk.WalkExpression(roots, r)
}

func (r *Rule) WalkMethodExpr(e *expr.MethodExpr) error {
	if !casing.Check(e.Name, r.cfg.WordCase) {
		return fmt.Errorf("goa-lint[%s]: Method names should be %s (%#v) in %s", r.Name(), r.cfg.WordCase, casing.To(e.Name, r.cfg.WordCase), e.EvalName())
	}

	return nil
}
