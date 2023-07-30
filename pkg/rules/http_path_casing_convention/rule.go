package http_path_casing_convention

import (
	"log"
	"strings"

	"github.com/NagayamaRyoga/goalint/pkg/common/casing"
	"github.com/NagayamaRyoga/goalint/pkg/common/walk"
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"github.com/NagayamaRyoga/goalint/pkg/rules"
	"goa.design/goa/v3/eval"
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
	return "HTTPPathCasingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Path(roots, r.walkPath)
}

func (r *Rule) walkPath(e eval.Expression, path string) (rl reports.ReportList) {
	for _, pathSeg := range strings.Split(path, "/") {
		if strings.HasPrefix(pathSeg, "{") && strings.HasSuffix(pathSeg, "}") {
			continue
		}

		if !r.caser.Check(pathSeg) {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				e.EvalName(),
				"Path segment %q should be %s (%q)", pathSeg, r.cfg.WordCase, r.caser.To(path),
			))
		}
	}

	return
}
