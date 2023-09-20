package walk

import (
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type (
	ExpressionWalkerFunc func(e eval.Expression) reports.ReportList
	PathWalkerFunc       func(e eval.Expression, path string) reports.ReportList
	TypeWalkerFunc       func(t expr.DataType) reports.ReportList
)

func Expression(roots []eval.Root, walker ExpressionWalkerFunc) (rl reports.ReportList) {
	for _, root := range roots {
		root.WalkSets(func(s eval.ExpressionSet) {
			for _, e := range s {
				rl = append(rl, walker(e)...)
			}
		})
	}

	return
}

func Path(roots []eval.Root, walker PathWalkerFunc) reports.ReportList {
	return Expression(roots, func(e eval.Expression) (rl reports.ReportList) {
		switch e := e.(type) {
		case *expr.APIExpr:
			http := e.HTTP
			rl = append(rl, walker(http, http.Path)...)

		case *expr.HTTPServiceExpr:
			for _, path := range e.Paths {
				rl = append(rl, walker(e, path)...)
			}

		case *expr.HTTPEndpointExpr:
			for _, route := range e.Routes {
				rl = append(rl, walker(route, route.Path)...)
			}
		}

		return
	})
}

func Type(roots []eval.Root, walker TypeWalkerFunc) (rl reports.ReportList) {
	for _, root := range roots {
		if root, ok := root.(*expr.RootExpr); ok {
			generatedTypes := make(map[expr.UserType]struct{})
			if root.GeneratedTypes != nil {
				for _, gt := range *root.GeneratedTypes {
					generatedTypes[gt] = struct{}{}
				}
			}

			for _, t := range root.Types {
				if _, ok := generatedTypes[t]; !ok {
					rl = append(rl, walker(t)...)
				}
			}

			for _, t := range root.ResultTypes {
				if _, ok := generatedTypes[t]; !ok {
					rl = append(rl, walker(t)...)
				}
			}

			for _, service := range root.Services {
				for _, method := range service.Methods {
					rl = append(rl, walker(method.Payload.Type)...)
					rl = append(rl, walker(method.Result.Type)...)
				}
			}
		}
	}

	return
}
