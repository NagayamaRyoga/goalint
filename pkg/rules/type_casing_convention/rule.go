package type_casing_convention

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/casing"
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
	return "TypeCasingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.WalkUserType)
}

func (r *Rule) WalkUserType(t expr.UserType) (rl reports.ReportList) {
	if !r.caser.Check(t.Name()) {
		kind := kind.DSLName(t.Kind())

		rl = append(rl, reports.NewReport(
			r.cfg.Level,
			r.Name(),
			fmt.Sprintf("%s(%q)", kind, t.ID()),
			"%s name %q should be %s (%q)", kind, t.Name(), r.cfg.WordCase, r.caser.To(t.Name()),
		))
	}

	return
}
