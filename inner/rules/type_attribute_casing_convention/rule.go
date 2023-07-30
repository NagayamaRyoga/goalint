package type_attribute_casing_convention

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/inner/common/casing"
	"github.com/NagayamaRyoga/goalint/inner/common/kind"
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
	return "TypeAttributeCasingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.WalkType)
}

func (r *Rule) WalkType(t expr.UserType) (rl reports.ReportList) {
	kind := kind.DSLName(t.Kind())

	if obj, ok := t.Attribute().Type.(*expr.Object); ok {
		for _, attr := range *obj {
			if !r.caser.Check(attr.Name) {
				rl = append(rl, reports.NewReport(
					r.cfg.Level,
					r.Name(),
					fmt.Sprintf("attribute %q in %s(%q)", attr.Name, kind, t.ID()),
					"Attribute name %q should be %s (%q)", attr.Name, r.cfg.WordCase, r.caser.To(attr.Name),
				))
			}
		}
	}

	return
}
