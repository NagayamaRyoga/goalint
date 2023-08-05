package http_path_naming_convention

import (
	"log"
	"strings"

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
	return "HTTPPathNamingConvention"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Path(roots, r.WalkPath)
}

func (r *Rule) WalkPath(e eval.Expression, path string) (rl reports.ReportList) {
	if _, ok := e.(*expr.HTTPExpr); ok && path == "" {
		return
	}
	if _, ok := e.(*expr.HTTPServiceExpr); ok && path == "" {
		return
	}

	if !strings.HasPrefix(path, "/") {
		rl = append(rl, reports.NewReport(
			r.cfg.Level,
			r.Name(),
			e.EvalName(),
			"Path should starts with / (%q)", "/"+path,
		))
	}

	return
}
