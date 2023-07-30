package type_casing_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/type_casing_convention"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		typeWithValidName       = dsl.Type("UserCredential", func() {})
		resultTypeWithValidName = dsl.ResultType("application/vnd.user-credential", func() { dsl.TypeName("UserCredential") })

		typeWithInvalidName       = dsl.Type("user_credential", func() {})
		resultTypeWithInvalidName = dsl.ResultType("application/vnd.user_credential", func() {})
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
			userType:    typeWithValidName,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			userType:    resultTypeWithValidName,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			userType:    typeWithInvalidName,
			wantReports: 1,
		},
		{
			description: "failed/ResultType",
			userType:    resultTypeWithInvalidName,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := type_casing_convention.NewConfig()
			rule := type_casing_convention.NewRule(logger, cfg)

			// when
			got := rule.WalkUserType(tc.userType)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
