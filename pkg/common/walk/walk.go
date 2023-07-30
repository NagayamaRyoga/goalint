package walk

import (
	"github.com/NagayamaRyoga/goalint/pkg/reports"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type (
	ExpressionWalkerFunc func(e eval.Expression) reports.ReportList
	PathWalkerFunc       func(e eval.Expression, path string) reports.ReportList
	TypeWalkerFunc       func(t expr.UserType) reports.ReportList
)

func Expression(roots []eval.Root, walker ExpressionWalkerFunc) (rl reports.ReportList) {
	for _, root := range roots {
		root.WalkSets(func(s eval.ExpressionSet) error {
			for _, e := range s {
				rl = append(rl, walker(e)...)
			}

			return nil
		})
	}

	return
}

func Path(roots []eval.Root, walker PathWalkerFunc) reports.ReportList {
	return Expression(roots, func(e eval.Expression) (rl reports.ReportList) {
		switch e := e.(type) {
		case *expr.RootExpr:
			http := e.API.HTTP
			rl = append(rl, walker(http, http.Path)...)

		case *expr.HTTPEndpointExpr:
			for _, path := range e.Service.Paths {
				rl = append(rl, walker(e.Service, path)...)
			}

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
		}
	}

	return
}
