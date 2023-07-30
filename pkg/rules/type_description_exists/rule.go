package type_description_exists

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/kind"
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
	return walk.Type(roots, r.WalkUserType)
}

func (r *Rule) WalkUserType(t expr.UserType) (rl reports.ReportList) {
	if len(t.Attribute().Description) == 0 {
		kind := kind.DSLName(t.Kind())

		rl = append(rl, reports.NewReport(
			r.cfg.Level,
			r.Name(),
			fmt.Sprintf("%s(%q)", kind, t.ID()),
			"%s should have non-empty description", kind,
		))
	}

	return
}
