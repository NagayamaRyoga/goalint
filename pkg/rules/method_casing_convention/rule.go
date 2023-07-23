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
	logger *log.Logger
	cfg    *config.MethodCasingConvention
	caser  *casing.Caser
}

func NewRule(logger *log.Logger, c *config.Config) rules.Rule {
	cfg := c.MethodCasingConvention

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

func (r *Rule) Apply(roots []eval.Root) error {
	return walk.Expression(roots, r.walkExpr)
}

func (r *Rule) walkExpr(e eval.Expression) error {
	if e, ok := e.(*expr.MethodExpr); ok {
		if !r.caser.Check(e.Name) {
			return fmt.Errorf("goa-lint[%s]: Method name %q should be %s (%q) in %s", r.Name(), e.Name, r.cfg.WordCase, r.caser.To(e.Name), e.EvalName())
		}
	}

	return nil
}
