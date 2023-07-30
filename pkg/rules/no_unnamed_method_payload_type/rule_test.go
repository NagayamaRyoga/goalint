package no_unnamed_method_payload_type_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/no_unnamed_method_payload_type"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		AddNumbersPayload = dsl.Type("AddNumbersPayload", func() {
			dsl.Attribute("a", expr.Int)
			dsl.Attribute("b", expr.Int)
			dsl.Required("a", "b")
		})

		methodWithNamedPayload = dsl.Service("calc", func() {
			dsl.Method("add_numbers", func() {
				dsl.Payload(AddNumbersPayload)
			})
		})

		methodWithUnnamedPayload = dsl.Service("calc2", func() {
			dsl.Method("add_numbers2", func() {
				dsl.Payload(func() {
					dsl.Attribute("a", expr.Int)
					dsl.Attribute("b", expr.Int)
					dsl.Required("a", "b")
				})
			})
		})
	)

	// given
	err := eval.RunDSL()
	assert.NoError(t, err)

	testCases := []struct {
		description string
		service     *expr.ServiceExpr
		wantReports int
	}{
		{
			description: "success",
			service:     methodWithNamedPayload,
			wantReports: 0,
		},
		{
			description: "failed",
			service:     methodWithUnnamedPayload,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := no_unnamed_method_payload_type.NewConfig()
			rule := no_unnamed_method_payload_type.NewRule(logger, cfg)

			// when
			assert.Equal(t, 1, len(tc.service.Methods))

			got := rule.WalkMethodExpr(tc.service.Methods[0])
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
