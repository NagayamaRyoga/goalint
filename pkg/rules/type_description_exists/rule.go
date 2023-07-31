package type_description_exists

import (
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/datatype"
	"github.com/NagayamaRyoga/goalint/pkg/common/walk"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"github.com/NagayamaRyoga/goalint/pkg/rules"
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
	return "TypeDescriptionExists"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.WalkType)
}

func (r *Rule) WalkType(t expr.DataType) (rl reports.ReportList) {
	if t, ok := t.(expr.UserType); ok {
		if len(t.Attribute().Description) == 0 {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				datatype.TypeName(t),
				"%s should have non-empty description", datatype.DSLName(t),
			))
		}
	}

	return
}
