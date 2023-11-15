package type_attribute_description_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_description_exists"
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
		typeWithDescription = dsl.Type("User", func() {
			dsl.Attribute("id", dsl.Int, "User ID")
			dsl.Attribute("name", dsl.String, func() {
				dsl.Description("User name")
			})
		})

		resultTypeWithDescription = dsl.ResultType("application/vnd.user", func() {
			dsl.Attribute("id", dsl.Int, "User ID")
			dsl.Attribute("name", dsl.String, func() {
				dsl.Description("User name")
			})
		})

		typeWithoutDescription = dsl.Type("Book", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("title", dsl.String, "")
		})

		resultTypeWithoutDescription = dsl.ResultType("application/vnd.book", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("title", dsl.String, func() {
				dsl.Description("")
			})
		})

		payloadWithDescription = dsl.Service("calc", func() {
			dsl.Method("add", func() {
				dsl.Payload(func() {
					dsl.Attribute("lhs", dsl.Int, "Left hand side operand")
					dsl.Attribute("rhs", dsl.Int, func() {
						dsl.Description("Right hand side operand")
					})
				})
			})
		})

		payloadWithoutDescription = dsl.Service("calc2", func() {
			dsl.Method("add2", func() {
				dsl.Payload(func() {
					dsl.Attribute("lhs", dsl.Int)
					dsl.Attribute("rhs", dsl.Int, func() {
					})
				})
			})
		})
	)

	// given
	err := eval.RunDSL()
	require.NoError(t, err)

	testCases := []struct {
		description string
		dataType    expr.DataType
		wantReports int
	}{
		{
			description: "success/Type",
			dataType:    typeWithDescription,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			dataType:    resultTypeWithDescription,
			wantReports: 0,
		},
		{
			description: "success/Payload",
			dataType:    payloadWithDescription.Methods[0].Payload.Type,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			dataType:    typeWithoutDescription,
			wantReports: 2,
		},
		{
			description: "failed/ResultType",
			dataType:    resultTypeWithoutDescription,
			wantReports: 2,
		},
		{
			description: "failed/Payload",
			dataType:    payloadWithoutDescription.Methods[0].Payload.Type,
			wantReports: 2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := type_attribute_description_exists.NewConfig()
			rule := type_attribute_description_exists.NewRule(logger, cfg)

			// when
			got := rule.WalkType(tc.dataType)
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
