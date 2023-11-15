package http_error_description_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/http_error_description_exists"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	_ = dsl.Service("calc", func() {
		dsl.Method("good_ok", func() {
			dsl.HTTP(func() {
				dsl.GET("/good_ok")
				dsl.Response(dsl.StatusOK)
			})
		})
		dsl.Method("good_error", func() {
			dsl.HTTP(func() {
				dsl.GET("/good_error")
				dsl.Response(dsl.StatusOK)
				dsl.Response("not_found", dsl.StatusNotFound, func() {
					dsl.Description("entity is not found")
				})
			})
		})
		dsl.Method("bad_error", func() {
			dsl.HTTP(func() {
				dsl.GET("/bad_error")
				dsl.Response("bad_request", dsl.StatusBadRequest)
			})
		})
		dsl.Error("not_found")
		dsl.Error("bad_request")
	})

	err := eval.RunDSL()
	assert.NoError(t, err)

	roots, err := eval.Context.Roots()
	assert.NoError(t, err)

	httpSvc := roots[0].(*expr.RootExpr).HTTPService("calc")

	testCases := []struct {
		description string
		expr        *expr.HTTPEndpointExpr
		wantReports int
	}{
		{
			description: "success/ok only",
			expr:        httpSvc.Endpoint("good_ok"),
			wantReports: 0,
		},
		{
			description: "success/error with description",
			expr:        httpSvc.Endpoint("good_error"),
			wantReports: 0,
		},
		{
			description: "failed/no description",
			expr:        httpSvc.Endpoint("bad_error"),
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := http_error_description_exists.NewConfig()
			rule := http_error_description_exists.NewRule(logger, cfg)

			// when

			got := rule.WalkHTTPEndpointExpr(tc.expr)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
