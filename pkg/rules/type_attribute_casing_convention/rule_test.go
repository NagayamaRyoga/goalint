package type_attribute_casing_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_casing_convention"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
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
	)

	// given
	err := eval.RunDSL()
	assert.NoError(t, err)

	testCases := []struct {
		description string
		userType    expr.UserType
		wantReports int
	}{
		{
			description: "success/Type",
			userType:    typeWithValidSnakeCasedAttributes,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			userType:    resultTypeWithValidSnakeCasedAttributes,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			userType:    typeWithValidCamelCasedAttributes,
			wantReports: 3,
		},
		{
			description: "failed/ResultType",
			userType:    resultTypeWithValidCamelCasedAttributes,
			wantReports: 3,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := type_attribute_casing_convention.NewConfig()
			rule := type_attribute_casing_convention.NewRule(logger, cfg)

			// when
			got := rule.WalkType(tc.userType)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
