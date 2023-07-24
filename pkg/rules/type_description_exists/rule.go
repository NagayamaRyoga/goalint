package type_description_exists

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/kind"
	"github.com/NagayamaRyoga/goalint/pkg/common/walk"
	"github.com/NagayamaRyoga/goalint/pkg/config"
	"github.com/NagayamaRyoga/goalint/pkg/rules"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	logger *log.Logger
	cfg    *config.TypeDescriptionExists
}

func NewRule(logger *log.Logger, c *config.Config) rules.Rule {
	cfg := c.TypeDescriptionExists

	return &Rule{
		logger: logger,
		cfg:    cfg,
	}
}

func (r *Rule) Name() string {
	return "TypeDescriptionExists"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) error {
	return walk.Type(roots, r.walkType)
}

func (r *Rule) walkType(t expr.UserType) error {
	if len(t.Attribute().Description) == 0 {
		kind := kind.DSLName(t.Kind())

		return fmt.Errorf("goa-lint[%s]: %s should have non-empty description in %s(%q)", r.Name(), kind, kind, t.ID())
	}

	return nil
}
