package type_attribute_casing_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_casing_convention"
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
		typeWithValidSnakeCasedAttributes = dsl.Type("User", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("given_name", dsl.String)
			dsl.Attribute("family_name", dsl.String)
			dsl.Attribute("phone_number", dsl.String)
		})

		resultTypeWithValidSnakeCasedAttributes = dsl.ResultType("application/vnd.user2", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("given_name", dsl.String)
			dsl.Attribute("family_name", dsl.String)
			dsl.Attribute("phone_number", dsl.String)
		})

		typeWithValidCamelCasedAttributes = dsl.Type("User3", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("givenName", dsl.String)
			dsl.Attribute("familyName", dsl.String)
			dsl.Attribute("phoneNumber", dsl.String)
		})

		resultTypeWithValidCamelCasedAttributes = dsl.ResultType("application/vnd.user4", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("givenName", dsl.String)
			dsl.Attribute("familyName", dsl.String)
			dsl.Attribute("phoneNumber", dsl.String)
		})

		payloadWithSnakeCasedAttributes = dsl.Service("calc", func() {
			dsl.Method("add", func() {
				dsl.Payload(func() {
					dsl.Attribute("left_hand_side", dsl.Int)
					dsl.Attribute("right_hand_side", dsl.Int)
				})
			})
		})

		payloadWithPascalCasedAttributes = dsl.Service("calc2", func() {
			dsl.Method("add2", func() {
				dsl.Payload(func() {
					dsl.Attribute("LeftHandSide", dsl.Int)
					dsl.Attribute("RightHandSide", dsl.Int)
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
			dataType:    typeWithValidSnakeCasedAttributes,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			dataType:    resultTypeWithValidSnakeCasedAttributes,
			wantReports: 0,
		},
		{
			description: "success/Payload",
			dataType:    payloadWithSnakeCasedAttributes.Methods[0].Payload.Type,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			dataType:    typeWithValidCamelCasedAttributes,
			wantReports: 3,
		},
		{
			description: "failed/ResultType",
			dataType:    resultTypeWithValidCamelCasedAttributes,
			wantReports: 3,
		},
		{
			description: "failed/Payload",
			dataType:    payloadWithPascalCasedAttributes.Methods[0].Payload.Type,
			wantReports: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := type_attribute_casing_convention.NewConfig()
			rule := type_attribute_casing_convention.NewRule(logger, cfg)

			// when
			got := rule.WalkType(tc.dataType)
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
