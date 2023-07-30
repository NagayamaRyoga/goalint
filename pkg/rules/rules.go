package rules

import (
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"goa.design/goa/v3/eval"
)

type Rule interface {
	Name() string
	IsDisabled() bool
	Apply(roots []eval.Root) reports.ReportList
}
