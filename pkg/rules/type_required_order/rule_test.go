package type_required_order_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/type_required_order"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		typeWithValidRequiredOrder = dsl.Type("User", func() {
			dsl.Required("id", "name")
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("role", dsl.String)
			dsl.Attribute("name", dsl.String)
		})

		resultTypeValidRequiredOrder = dsl.ResultType("application/vnd.user2", func() {
			dsl.Required("id", "name")
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("role", dsl.String)
			dsl.Attribute("name", dsl.String)
		})

		typeWithInvalidRequiredOrder = dsl.Type("User3", func() {
			dsl.Required("name", "id")
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("role", dsl.String)
			dsl.Attribute("name", dsl.String)
		})

		resultTypeInvalidRequiredOrder = dsl.ResultType("application/vnd.user4", func() {
			dsl.Required("name", "id")
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("role", dsl.String)
			dsl.Attribute("name", dsl.String)
		})
	)

	// given
	err := eval.RunDSL()
	assert.NoError(t, err)

	testCases := []struct {
		description string
		dataType    expr.DataType
		wantReports int
	}{
		{
			description: "success/Type",
			dataType:    typeWithValidRequiredOrder,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			dataType:    resultTypeValidRequiredOrder,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			dataType:    typeWithInvalidRequiredOrder,
			wantReports: 1,
		},
		{
			description: "failed/ResultType",
			dataType:    resultTypeInvalidRequiredOrder,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := type_required_order.NewConfig()
			rule := type_required_order.NewRule(logger, cfg)

			// when
			got := rule.WalkType(tc.dataType)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
