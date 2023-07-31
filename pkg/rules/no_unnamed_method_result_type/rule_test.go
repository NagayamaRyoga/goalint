package no_unnamed_method_result_type_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/no_unnamed_method_result_type"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		AddNumbersResult = dsl.Type("AddNumbersResult", func() {
			dsl.Attribute("sum", expr.Int)
			dsl.Required("sum")
		})

		methodWithNamedResult = dsl.Service("calc", func() {
			dsl.Method("add_numbers", func() {
				dsl.Result(AddNumbersResult)
			})
		})

		methodWithUnnamedResult = dsl.Service("calc2", func() {
			dsl.Method("add_numbers2", func() {
				dsl.Result(func() {
					dsl.Attribute("sum", expr.Int)
					dsl.Required("sum")
				})
			})
		})

		methodWithIntResult = dsl.Service("calc3", func() {
			dsl.Method("add_numbers3", func() {
				dsl.Result(dsl.Int)
			})
		})
	)

	// given
	err := eval.RunDSL()
	assert.NoError(t, err)

	testCases := []struct {
		description string
		service     *expr.ServiceExpr
		wantReports int
	}{
		{
			description: "success",
			service:     methodWithNamedResult,
			wantReports: 0,
		},
		{
			description: "failed/object",
			service:     methodWithUnnamedResult,
			wantReports: 1,
		},
		{
			description: "failed/int",
			service:     methodWithIntResult,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := no_unnamed_method_result_type.NewConfig()
			rule := no_unnamed_method_result_type.NewRule(logger, cfg)

			// when
			assert.Equal(t, 1, len(tc.service.Methods))

			got := rule.WalkMethodExpr(tc.service.Methods[0])
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
