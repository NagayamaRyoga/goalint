package method_description_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/inner/rules/method_description_exists"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		methodWithDescription = dsl.Service("calc", func() {
			dsl.Method("add_numbers", func() {
				dsl.Description("Adds up two numbers")
			})
		})

		methodWithoutDescription = dsl.Service("calc2", func() {
			dsl.Method("add_numbers2", func() {
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
			service:     methodWithDescription,
			wantReports: 0,
		},
		{
			description: "failed",
			service:     methodWithoutDescription,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := method_description_exists.NewConfig()
			rule := method_description_exists.NewRule(logger, cfg)

			// when
			assert.Equal(t, 1, len(tc.service.Methods))

			got := rule.WalkMethodExpr(tc.service.Methods[0])
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
