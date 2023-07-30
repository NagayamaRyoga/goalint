package type_casing_convention

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
	return "TypeCasingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Type(roots, r.walkType)
}

func (r *Rule) walkType(t expr.UserType) (rl reports.ReportList) {
	if !r.caser.Check(t.Name()) {
		kind := kind.DSLName(t.Kind())

		rl = append(rl, reports.NewReport(
			r.cfg.Level,
			fmt.Sprintf("%s(%q)", kind, t.ID()),
			"%s name %q should be %s (%q)", kind, t.Name(), r.cfg.WordCase, r.caser.To(t.Name()),
		))
	}

	return
}
