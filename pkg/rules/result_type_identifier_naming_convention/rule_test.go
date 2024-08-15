package result_type_identifier_naming_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/result_type_identifier_naming_convention"
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
		resultTypeWithValidID   = dsl.ResultType("application/vnd.user", func() {})
		resultTypeWithInvalidID = dsl.ResultType("book", func() {})
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
			description: "success",
			dataType:    resultTypeWithValidID,
			wantReports: 0,
		},
		{
			description: "failed",
			dataType:    resultTypeWithInvalidID,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := result_type_identifier_naming_convention.NewConfig()
			rule := result_type_identifier_naming_convention.NewRule(logger, cfg)

			// when
			got := rule.WalkType(tc.dataType)
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
