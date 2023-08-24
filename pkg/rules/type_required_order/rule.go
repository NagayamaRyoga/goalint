package type_required_order

import (
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
	return "TypeRequiredOrder"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.WalkType)
}

func (r *Rule) WalkType(t expr.DataType) (rl reports.ReportList) {
	if t, ok := t.(expr.UserType); ok {
		obj := expr.AsObject(t)
		if obj == nil || t.Attribute().Validation == nil {
			return
		}

		nameIndexMap := make(map[string]int)
		{
			for _, attr := range *obj {
				if _, ok := nameIndexMap[attr.Name]; !ok {
					nameIndexMap[attr.Name] = len(nameIndexMap)
				}
			}
		}

		required := t.Attribute().Validation.Required

		isSorted := lo.IsSortedByKey(required, func(name string) int {
			return lo.ValueOr(nameIndexMap, name, -1)
		})

		if !isSorted {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				datatype.TypeName(t),
				"Required attributes should be in order",
			))
		}
	}

	return
}
