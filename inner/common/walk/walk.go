package walk

import (
	"github.com/hashicorp/go-multierror"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type (
	ExpressionWalkerFunc func(e eval.Expression) error
	TypeWalkerFunc       func(t expr.UserType) error
)

func Expression(roots []eval.Root, walker ExpressionWalkerFunc) error {
	var merr error

	for _, root := range roots {
		root.WalkSets(func(s eval.ExpressionSet) error {
			for _, e := range s {
				if err := walker(e); err != nil {
					merr = multierror.Append(merr, err)
				}
			}

			return nil
		})
	}

	return merr
}

func Type(roots []eval.Root, walker TypeWalkerFunc) error {
	var merr error

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
					if err := walker(t); err != nil {
						merr = multierror.Append(merr, err)
					}
				}
			}

			for _, t := range root.ResultTypes {
				if _, ok := generatedTypes[t]; !ok {
					if err := walker(t); err != nil {
						merr = multierror.Append(merr, err)
					}
				}
			}
		}
	}

	return merr
}
