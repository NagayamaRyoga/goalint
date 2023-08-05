package method_array_result_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/method_array_result"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		Response = dsl.ResultType("application/vnd.response", func() {
			dsl.Required("titles")
			dsl.Attribute("titles", dsl.ArrayOf(dsl.String))
		})

		methodReturnsNonArray = dsl.Service("calc", func() {
			dsl.Method("list_titles", func() {
				dsl.Result(Response)
			})
		})

		methodReturnsArray = dsl.Service("calc2", func() {
			dsl.Method("list_titles2", func() {
				dsl.Result(dsl.ArrayOf(dsl.String))
			})
		})

		Item = dsl.ResultType("application/vnd.item", func() {
			dsl.Required("title")
			dsl.Attribute("title", dsl.String)
		})

		methodReturnsCollection = dsl.Service("calc3", func() {
			dsl.Method("list_titles3", func() {
				dsl.Result(dsl.CollectionOf(Item))
			})
		})
	)

	err := eval.RunDSL()
	assert.NoError(t, err)

	testCases := []struct {
		description string
		service     *expr.ServiceExpr
		wantReports int
	}{
		{
			description: "success",
			service:     methodReturnsNonArray,
			wantReports: 0,
		},
		{
			description: "failed/ArrayOf",
			service:     methodReturnsArray,
			wantReports: 1,
		},
		{
			description: "failed/CollectionOf",
			service:     methodReturnsCollection,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := method_array_result.NewConfig()
			rule := method_array_result.NewRule(logger, cfg)

			// when
			assert.Equal(t, 1, len(tc.service.Methods))

			got := rule.WalkMethodExpr(tc.service.Methods[0])
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
