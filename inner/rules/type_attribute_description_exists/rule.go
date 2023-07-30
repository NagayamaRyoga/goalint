package type_attribute_description_exists

import (
	"fmt"
	"log"

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
}

func NewRule(logger *log.Logger, cfg *Config) *Rule {
	return &Rule{
		logger: logger,
		cfg:    cfg,
	}
}

func (r *Rule) Name() string {
	return "TypeAttributeDescriptionExists"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.walkType)
}

func (r *Rule) walkType(t expr.UserType) (rl reports.ReportList) {
	kind := kind.DSLName(t.Kind())

	if obj, ok := t.Attribute().Type.(*expr.Object); ok {
		for _, attr := range *obj {
			if len(attr.Attribute.Description) == 0 {
				rl = append(rl, reports.NewReport(
					r.cfg.Level,
					r.Name(),
					fmt.Sprintf("attribute %q in %s(%q)", attr.Name, kind, t.ID()),
					"Attribute should have non-empty description",
				))
			}
		}
	}

	return
}
