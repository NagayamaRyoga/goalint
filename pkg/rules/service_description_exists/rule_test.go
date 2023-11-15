package service_description_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/service_description_exists"
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
		serviceWithDescription = dsl.Service("calc", func() {
			dsl.Description("A simple calculator service")
		})

		serviceWithoutDescription = dsl.Service("calc2", func() {
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
			service:     serviceWithDescription,
			wantReports: 0,
		},
		{
			description: "failed",
			service:     serviceWithoutDescription,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := service_description_exists.NewConfig()
			rule := service_description_exists.NewRule(logger, cfg)

			// when
			got := rule.WalkServiceExpr(tc.service)
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
