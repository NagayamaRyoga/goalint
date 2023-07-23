package walk

import (
	"log"
	"reflect"

	"github.com/hashicorp/go-multierror"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type ExpressionWalker interface {
	WalkAPIExpr(*expr.APIExpr) error
	WalkServerExpr(*expr.ServerExpr) error
	WalkServiceExpr(*expr.ServiceExpr) error
	WalkMethodExpr(*expr.MethodExpr) error
	WalkHTTPExpr(*expr.HTTPExpr) error
	WalkHTTPServiceExpr(*expr.HTTPServiceExpr) error
	WalkHTTPEndpointExpr(*expr.HTTPEndpointExpr) error
	WalkHTTPFileServerExpr(*expr.HTTPFileServerExpr) error
	WalkGRPCExpr(*expr.GRPCExpr) error
	WalkGRPCServiceExpr(*expr.GRPCServiceExpr) error
	WalkGRPCEndpointExpr(*expr.GRPCEndpointExpr) error
	WalkAttributeExpr(*expr.AttributeExpr) error
	WalkResultTypeExpr(*expr.ResultTypeExpr) error
	WalkUnknownExpr(eval.Expression) error
}

func WalkExpression(roots []eval.Root, walker ExpressionWalker) error {
	var merr error

	for _, root := range roots {
		root.WalkSets(func(s eval.ExpressionSet) error {
			for _, e := range s {
				if err := walkExpr(e, walker); err != nil {
					merr = multierror.Append(merr, err)
				}
			}
			return nil
		})
	}

	return merr
}

func walkExpr(e eval.Expression, walker ExpressionWalker) error {
	switch e := e.(type) {
	case *expr.APIExpr:
		return walker.WalkAPIExpr(e)
	case *expr.ServerExpr:
		return walker.WalkServerExpr(e)
	case *expr.ServiceExpr:
		return walker.WalkServiceExpr(e)
	case *expr.MethodExpr:
		return walker.WalkMethodExpr(e)
	case *expr.HTTPExpr:
		return walker.WalkHTTPExpr(e)
	case *expr.HTTPServiceExpr:
		return walker.WalkHTTPServiceExpr(e)
	case *expr.HTTPEndpointExpr:
		return walker.WalkHTTPEndpointExpr(e)
	case *expr.HTTPFileServerExpr:
		return walker.WalkHTTPFileServerExpr(e)
	case *expr.GRPCExpr:
		return walker.WalkGRPCExpr(e)
	case *expr.GRPCServiceExpr:
		return walker.WalkGRPCServiceExpr(e)
	case *expr.GRPCEndpointExpr:
		return walker.WalkGRPCEndpointExpr(e)
	case *expr.AttributeExpr:
		return walker.WalkAttributeExpr(e)
	case *expr.ResultTypeExpr:
		return walker.WalkResultTypeExpr(e)
	default:
		return walker.WalkUnknownExpr(e)
	}
}

var _ ExpressionWalker = (*BaseExpressionWalker)(nil)

type BaseExpressionWalker struct {
	logger *log.Logger
}

func NewBaseExpressionWalker(logger *log.Logger) *BaseExpressionWalker {
	return &BaseExpressionWalker{
		logger: logger,
	}
}

func (w *BaseExpressionWalker) WalkAPIExpr(*expr.APIExpr) error                       { return nil }
func (w *BaseExpressionWalker) WalkServerExpr(*expr.ServerExpr) error                 { return nil }
func (w *BaseExpressionWalker) WalkServiceExpr(*expr.ServiceExpr) error               { return nil }
func (w *BaseExpressionWalker) WalkMethodExpr(*expr.MethodExpr) error                 { return nil }
func (w *BaseExpressionWalker) WalkHTTPExpr(*expr.HTTPExpr) error                     { return nil }
func (w *BaseExpressionWalker) WalkHTTPServiceExpr(*expr.HTTPServiceExpr) error       { return nil }
func (w *BaseExpressionWalker) WalkHTTPEndpointExpr(*expr.HTTPEndpointExpr) error     { return nil }
func (w *BaseExpressionWalker) WalkHTTPFileServerExpr(*expr.HTTPFileServerExpr) error { return nil }
func (w *BaseExpressionWalker) WalkGRPCExpr(*expr.GRPCExpr) error                     { return nil }
func (w *BaseExpressionWalker) WalkGRPCServiceExpr(*expr.GRPCServiceExpr) error       { return nil }
func (w *BaseExpressionWalker) WalkGRPCEndpointExpr(*expr.GRPCEndpointExpr) error     { return nil }
func (w *BaseExpressionWalker) WalkAttributeExpr(*expr.AttributeExpr) error           { return nil }
func (w *BaseExpressionWalker) WalkResultTypeExpr(*expr.ResultTypeExpr) error         { return nil }

func (w *BaseExpressionWalker) WalkUnknownExpr(e eval.Expression) error {
	w.logger.Printf("WalkUnknownExpr(%s): %s\n", reflect.TypeOf(e), e.EvalName())

	return nil
}
