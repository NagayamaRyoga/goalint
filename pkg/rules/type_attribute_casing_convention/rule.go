package type_attribute_casing_convention

import (
	"fmt"
	"log"

	"github.com/NagayamaRyoga/goalint/pkg/common/casing"
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

func (r *Rule) WalkType(t expr.DataType) (rl reports.ReportList) {
	if obj := expr.AsObject(t); obj != nil {
		for _, attr := range *obj {
			if !r.caser.Check(attr.Name) {
				rl = append(rl, reports.NewReport(
					r.cfg.Level,
					r.Name(),
					fmt.Sprintf("attribute %q in %s", attr.Name, datatype.TypeName(t)),
					"Attribute name %q should be %s (%q)", attr.Name, r.cfg.WordCase, r.caser.To(attr.Name),
				))
			}
		}
	}

	return
}
