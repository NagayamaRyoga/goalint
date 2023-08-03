package type_attribute_example_exists_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/type_attribute_example_exists"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestRule(t *testing.T) {
	t.Parallel()

	var (
		typeWithExample = dsl.Type("User", func() {
			dsl.Attribute("id", dsl.Int, func() {
				dsl.Example(1)
			})
			dsl.Attribute("name", dsl.String, func() {
				dsl.Example("User name")
			})
		})

		resultTypeWithExample = dsl.ResultType("application/vnd.user", func() {
			dsl.Attribute("id", dsl.Int, func() {
				dsl.Example(1)
			})
			dsl.Attribute("name", dsl.String, func() {
				dsl.Example("User name")
			})
			dsl.Attribute("parent", typeWithExample)
		})

		typeWithoutExample = dsl.Type("Book", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("title", dsl.String)
		})

		resultTypeWithoutExample = dsl.ResultType("application/vnd.book", func() {
			dsl.Attribute("id", dsl.Int)
			dsl.Attribute("title", dsl.String, func() {
				dsl.Example("")
			})
		})

		payloadWithExample = dsl.Service("calc", func() {
			dsl.Method("add", func() {
				dsl.Payload(func() {
					dsl.Attribute("lhs", dsl.Int, func() {
						dsl.Example(10)
					})
					dsl.Attribute("rhs", dsl.Int, func() {
						dsl.Example(20)
					})
				})
			})
		})

		payloadWithoutExample = dsl.Service("calc2", func() {
			dsl.Method("add2", func() {
				dsl.Payload(func() {
					dsl.Attribute("lhs", dsl.Int)
					dsl.Attribute("rhs", dsl.Int)
				})
			})
		})

		arrayTypeWithExample = dsl.Type("Thesis", func() {
			dsl.Attribute("authors", dsl.ArrayOf(dsl.String), func() {
				dsl.Example([]string{"Author 1", "Author 2"})
			})
		})

		arrayTypeWithoutExample = dsl.Type("Thesis2", func() {
			dsl.Attribute("authors", dsl.ArrayOf(dsl.String))
		})
	)

	// given
	err := eval.RunDSL()
	assert.NoError(t, err)

	testCases := []struct {
		description string
		dataType    expr.DataType
		wantReports int
	}{
		{
			description: "success/Type",
			dataType:    typeWithExample,
			wantReports: 0,
		},
		{
			description: "success/ResultType",
			dataType:    resultTypeWithExample,
			wantReports: 0,
		},
		{
			description: "success/Payload",
			dataType:    payloadWithExample.Methods[0].Payload.Type,
			wantReports: 0,
		},
		{
			description: "failed/Type",
			dataType:    typeWithoutExample,
			wantReports: 2,
		},
		{
			description: "failed/ResultType",
			dataType:    resultTypeWithoutExample,
			wantReports: 1,
		},
		{
			description: "failed/Payload",
			dataType:    payloadWithoutExample.Methods[0].Payload.Type,
			wantReports: 2,
		},
		{
			description: "success/Array",
			dataType:    arrayTypeWithExample,
			wantReports: 0,
		},
		{
			description: "failed/Array",
			dataType:    arrayTypeWithoutExample,
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := type_attribute_example_exists.NewConfig()
			rule := type_attribute_example_exists.NewRule(logger, cfg)

			// when
			got := rule.WalkType(tc.dataType)
			assert.Equal(t, tc.wantReports, len(got))
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
