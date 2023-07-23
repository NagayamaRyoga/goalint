package type_casing_convention

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/casing"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/kind"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/common/walk"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/rules"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	logger *log.Logger
	cfg    *config.TypeCasingConventionConfig
	caser  *casing.Caser
}

func NewRule(logger *log.Logger, c *config.Config) rules.Rule {
	cfg := c.TypeCasingConvention

	return &Rule{
		logger: logger,
		cfg:    cfg,
		caser:  casing.NewCaser(cfg.WordCase, cfg.Initialisms),
	}
}

func (r *Rule) Name() string {
	return "TypeCasingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) error {
	return walk.Type(roots, r.walkType)
}

func (r *Rule) walkType(t expr.UserType) error {
	if !r.caser.Check(t.Name()) {
		kind := kind.DSLName(t.Kind())

		return fmt.Errorf("goa-lint[%s]: %s names should be %s (%#v) in %[2]s(%#v)", r.Name(), kind, r.cfg.WordCase, r.caser.To(t.Name()), t.Name())
	}

	return nil
}
