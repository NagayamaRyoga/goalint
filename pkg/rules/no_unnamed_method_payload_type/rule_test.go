package no_unnamed_method_payload_type_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/no_unnamed_method_payload_type"
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

		methodWithIntPayload = dsl.Service("calc3", func() {
			dsl.Method("add_numbers3", func() {
				dsl.Payload(dsl.Int)
			})
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
			service:     methodWithNamedPayload,
			wantReports: 0,
		},
		{
			description: "failed/object",
			service:     methodWithUnnamedPayload,
			wantReports: 1,
		},
		{
			description: "failed/int",
			service:     methodWithIntPayload,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := no_unnamed_method_payload_type.NewConfig()
			rule := no_unnamed_method_payload_type.NewRule(logger, cfg)

			// when
			assert.Len(t, tc.service.Methods, 1)

			got := rule.WalkMethodExpr(tc.service.Methods[0])
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
