package type_attribute_example_exists

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/datatype"
	"github.com/NagayamaRyoga/goalint/pkg/common/walk"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"github.com/NagayamaRyoga/goalint/pkg/rules"
	"github.com/samber/lo"
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
	return "TypeAttributeExampleExists"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.WalkType)
}

func (r *Rule) WalkType(t expr.DataType) (rl reports.ReportList) {
	if obj := expr.AsObject(t); obj != nil {
		for _, attr := range *obj {
			if len(attr.Attribute.UserExamples) > 0 {
				continue
			}

			if r.isRequiredExamples(attr.Attribute.Type) {
				rl = append(rl, reports.NewReport(
					r.cfg.Level,
					r.Name(),
					fmt.Sprintf("attribute %q in %s", attr.Name, datatype.TypeName(t)),
					"Attribute of type %s should have examples", datatype.TypeName(attr.Attribute.Type),
				))
			}
		}
	}

	return
}

func (r *Rule) isRequiredExamples(t expr.DataType) bool {
	if lo.Contains(r.cfg.RequiredTypes, t) {
		return true
	} else if lo.Contains(r.cfg.ExcludeTypes, t) {
		return false
	}

	if t := expr.AsArray(t); t != nil {
		return r.isRequiredExamples(t.ElemType.Type)
	}

	return false
}
