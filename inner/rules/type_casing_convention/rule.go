package type_casing_convention

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/inner/common/casing"
	"github.com/NagayamaRyoga/goalint/inner/common/kind"
	"github.com/NagayamaRyoga/goalint/inner/common/walk"
	"github.com/NagayamaRyoga/goalint/inner/rules"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	logger *log.Logger
	cfg    *Config
	caser  *casing.Caser
}

func NewRule(logger *log.Logger, cfg *Config) rules.Rule {
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

		return fmt.Errorf("goa-lint[%s]: %s name %q should be %s (%q) in %s(%q)", r.Name(), kind, t.Name(), r.cfg.WordCase, r.caser.To(t.Name()), kind, t.ID())
	}

	return nil
}
