package method_casing_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/method_casing_convention"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		methodWithSnakeName = dsl.Service("calc", func() {
			dsl.Method("add_numbers", func() {
				dsl.Description("Adds up two numbers")
			})
		})

		methodWithPascalName = dsl.Service("calc2", func() {
			dsl.Method("AddNumbers", func() {
				dsl.Description("Adds up two numbers")
			})
		})
	)

	// given
	err := eval.RunDSL()
	require.NoError(t, err)

	testCases := []struct {
		description string
		service     *expr.ServiceExpr
		wantReports int
	}{
		{
			description: "success",
			service:     methodWithSnakeName,
			wantReports: 0,
		},
		{
			description: "failed",
			service:     methodWithPascalName,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := method_casing_convention.NewConfig()
			rule := method_casing_convention.NewRule(logger, cfg)

			// when
			assert.Len(t, tc.service.Methods, 1)

			got := rule.WalkMethodExpr(tc.service.Methods[0])
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
