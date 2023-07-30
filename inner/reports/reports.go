package reports

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type ErrorLevel int

const (
	ErrorLevelInfo ErrorLevel = iota + 1
	ErrorLevelWarning
	ErrorLevelError
)

func (el ErrorLevel) String() string {
	switch el {
	case ErrorLevelInfo:
		return "info"
	case ErrorLevelWarning:
		return "warning"
	case ErrorLevelError:
		return "error"
	default:
		return "unknown"
	}
}

type ReportList []*Report

func (rl ReportList) String() string {
	var sb strings.Builder

	ce := rl.CountErrors()
	cw := rl.CountWarnings()
	if ce > 0 || cw > 0 {
		fmt.Fprintf(&sb, "Found %d errors and %d warnings\n", ce, cw)
	}

	for _, r := range rl {
		fmt.Fprintf(&sb, "%s\n", r.String())
	}

	return sb.String()
}

func (rl ReportList) CountWarnings() int {
	return lo.CountBy(rl, func(r *Report) bool {
		return r.IsWarning()
	})
}

func (rl ReportList) CountErrors() int {
	return lo.CountBy(rl, func(r *Report) bool {
		return r.IsError()
	})
}

type Report struct {
	Level    ErrorLevel
	Rule     string
	Location string
	Message  string
}

func NewReport(
	level ErrorLevel,
	rule string,
	location string,
	format string,
	args ...any,
) *Report {
	return &Report{
		Level:    level,
		Rule:     rule,
		Location: location,
		Message:  fmt.Sprintf(format, args...),
	}
}

func (r *Report) String() string {
	return fmt.Sprintf("[%s]: %s: %s in %s", r.Rule, r.Level, r.Message, r.Location)
}

func (r *Report) IsWarning() bool {
	return r.Level == ErrorLevelWarning
}

func (r *Report) IsError() bool {
	return r.Level == ErrorLevelError
}
