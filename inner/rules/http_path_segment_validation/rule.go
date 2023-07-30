package http_path_segment_validation

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/NagayamaRyoga/goalint/inner/common/walk"
	"github.com/NagayamaRyoga/goalint/inner/config"
	"github.com/NagayamaRyoga/goalint/inner/rules"
	"github.com/hashicorp/go-multierror"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

var _ rules.Rule = (*Rule)(nil)

type Rule struct {
	logger         *log.Logger
	cfg            *config.HTTPPathSegmentValidation
	pathSegPattern *regexp.Regexp
}

func NewRule(logger *log.Logger, c *config.Config) rules.Rule {
	cfg := c.HTTPPathSegmentValidation

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

func (r *Rule) Apply(roots []eval.Root) error {
	return walk.Expression(roots, r.walkExpr)
}

func (r *Rule) walkExpr(e eval.Expression) error {
	var merr error

	if e, ok := e.(*expr.RootExpr); ok {
		http := e.API.HTTP

		if err := r.validatePath(http, http.Path); err != nil {
			merr = multierror.Append(merr, err)
		}
	}

	if e, ok := e.(*expr.HTTPEndpointExpr); ok {
		for _, path := range e.Service.Paths {
			if err := r.validatePath(e.Service, path); err != nil {
				merr = multierror.Append(merr, err)
			}
		}

		for _, route := range e.Routes {
			if err := r.validatePath(route, route.Path); err != nil {
				merr = multierror.Append(merr, err)
			}
		}
	}

	return merr
}

func (r *Rule) validatePath(e eval.Expression, path string) error {
	var merr error

	for _, pathSeg := range strings.Split(path, "/") {
		if !r.isValidPathSegment(pathSeg) {
			err := fmt.Errorf("goa-lint[%s]: Invalid path parameter reference %q in %s", r.Name(), pathSeg, e.EvalName())
			merr = multierror.Append(merr, err)
		}
	}

	return merr
}

func (r *Rule) isValidPathSegment(pathSeg string) bool {
	if !strings.ContainsAny(pathSeg, "{*}") {
		return true
	}

	return r.pathSegPattern.MatchString(pathSeg)
}
