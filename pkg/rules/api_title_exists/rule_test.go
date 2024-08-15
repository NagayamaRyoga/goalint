package api_title_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/api_title_exists"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		apiWithTitle = dsl.API("calc1", func() {
			dsl.Title("Calculation API")
		})

		apiWithoutTitle = dsl.API("calc2", func() {
		})

		apiWithEmptyTitle = dsl.API("calc3", func() {
			dsl.Title("")
		})
	)

	// given
	testCases := []struct {
		description string
		api         *expr.APIExpr
		wantReports int
	}{
		{
			description: "success",
			api:         apiWithTitle,
			wantReports: 0,
		},
		{
			description: "failed/no title",
			api:         apiWithoutTitle,
			wantReports: 1,
		},
		{
			description: "failed/empty",
			api:         apiWithEmptyTitle,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			eval.Execute(tc.api.DSL(), tc.api)

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := api_title_exists.NewConfig()
			rule := api_title_exists.NewRule(logger, cfg)

			// when
			got := rule.WalkMethodExpr(tc.api)
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
