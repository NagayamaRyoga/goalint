package http_path_segment_validation

import (
	"log"
	"regexp"
	"strings"

	"github.com/NagayamaRyoga/goalint/inner/common/walk"
	"github.com/NagayamaRyoga/goalint/inner/reports"
	"github.com/NagayamaRyoga/goalint/inner/rules"
	"goa.design/goa/v3/eval"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	logger         *log.Logger
	cfg            *Config
	pathSegPattern *regexp.Regexp
}

func NewRule(logger *log.Logger, cfg *Config) *Rule {
	return &Rule{
		logger:         logger,
		cfg:            cfg,
		pathSegPattern: regexp.MustCompile(`^\{\*?[^{*}]+\}$`),
	}
}

func (r *Rule) Name() string {
	return "HTTPPathSegmentValidation"
}

func (r *Rule) IsDisabled() bool {
	return r.cfg.Disabled
}

func (r *Rule) Apply(roots []eval.Root) reports.ReportList {
	return walk.Path(roots, r.walkPath)
}

func (r *Rule) walkPath(e eval.Expression, path string) (rl reports.ReportList) {
	for _, pathSeg := range strings.Split(path, "/") {
		if !r.isValidPathSegment(pathSeg) {
			rl = append(rl, reports.NewReport(
				r.cfg.Level,
				r.Name(),
				e.EvalName(),
				"Invalid path parameter reference %q", pathSeg,
			))
		}
	}

	return
}

func (r *Rule) isValidPathSegment(pathSeg string) bool {
	if !strings.ContainsAny(pathSeg, "{*}") {
		return true
	}

	return r.pathSegPattern.MatchString(pathSeg)
}
