package result_type_identifier_naming_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/inner/rules/result_type_identifier_naming_convention"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		resultTypeWithValidID   = dsl.ResultType("application/vnd.user", func() {})
		resultTypeWithInvalidID = dsl.ResultType("book", func() {})
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
			description: "success",
			userType:    resultTypeWithValidID,
			wantReports: 0,
		},
		{
			description: "failed",
			userType:    resultTypeWithInvalidID,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := result_type_identifier_naming_convention.NewConfig()
			rule := result_type_identifier_naming_convention.NewRule(logger, cfg)

			// when
			got := rule.WalkResultType(tc.userType)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
