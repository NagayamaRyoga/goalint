package type_attribute_description_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/inner/rules/type_attribute_description_exists"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
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
			userType:    typeWithDescription,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			userType:    resultTypeWithDescription,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			userType:    typeWithoutDescription,
			wantReports: 2,
		},
		{
			description: "failed/ResultType",
			userType:    resultTypeWithoutDescription,
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
			got := rule.WalkUserType(tc.userType)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
