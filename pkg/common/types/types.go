package types

import (
	"goa.design/goa/v3/expr"
)

type DataTypeList []expr.DataType

func (dt *DataTypeList) Add(t ...expr.DataType) *DataTypeList {
	*dt = append(*dt, t...)

	return dt
}
