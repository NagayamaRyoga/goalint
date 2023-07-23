package kind

import (
	"goa.design/goa/v3/expr"
)

func DSLName(kind expr.Kind) string {
	switch kind {
	case expr.UserTypeKind:
		return "Type"
	case expr.ResultTypeKind:
		return "ResultType"
	default:
		return ""
	}
}
