package kind

import (
	"goa.design/goa/v3/expr"
)

func DSLName(kind expr.Kind) string {
	switch kind {
	case expr.BooleanKind:
		return "Boolean"
	case expr.IntKind:
		return "Int"
	case expr.Int32Kind:
		return "Int32"
	case expr.Int64Kind:
		return "Int64"
	case expr.UIntKind:
		return "UInt"
	case expr.UInt32Kind:
		return "UInt32"
	case expr.UInt64Kind:
		return "UInt64"
	case expr.Float32Kind:
		return "Float32"
	case expr.Float64Kind:
		return "Float64"
	case expr.StringKind:
		return "String"
	case expr.BytesKind:
		return "Bytes"
	case expr.ArrayKind:
		return "ArrayOf"
	case expr.ObjectKind:
		return "Object"
	case expr.MapKind:
		return "MapOf"
	case expr.UnionKind:
		return "OneOf"
	case expr.UserTypeKind:
		return "Type"
	case expr.ResultTypeKind:
		return "ResultType"
	case expr.AnyKind:
		return "Any"
	default:
		return ""
	}
}
