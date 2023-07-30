package service_description_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/inner/rules/service_description_exists"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
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

	testCases := []struct {
		description string
		service     *expr.ServiceExpr
	}{
		{
			description: "success",
			service:     serviceWithDescription,
		},
		{
			description: "failed",
			service:     serviceWithoutDescription,
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

			// given
			err := eval.RunDSL()
			assert.NoError(t, err)

			// when
			got := rule.WalkServiceExpr(tc.service)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
