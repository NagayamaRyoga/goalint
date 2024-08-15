package http_path_naming_convention_test

import (
	"log"
	"os"
	"testing"

	"github.com/NagayamaRyoga/goalint/pkg/rules/http_path_naming_convention"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
)

func TestRule(t *testing.T) {
	t.Parallel()

	service := dsl.Service("calc", func() {
		dsl.Method("good_root", func() {
			dsl.HTTP(func() {
				dsl.GET("/")
			})
		})
		dsl.Method("good_nested", func() {
			dsl.HTTP(func() {
				dsl.GET("/a/b")
			})
		})
		dsl.Method("bad_empty", func() {
			dsl.HTTP(func() {
				dsl.GET("")
			})
		})
		dsl.Method("bad_nested", func() {
			dsl.HTTP(func() {
				dsl.GET("a/b")
			})
		})
	})

	err := eval.RunDSL()
	require.NoError(t, err)

	testCases := []struct {
		description string
		expr        eval.Expression
		path        string
		wantReports int
	}{
		{
			description: "success/root",
			expr:        service.Method("good_root"),
			path:        "/",
			wantReports: 0,
		},
		{
			description: "success/nested",
			expr:        service.Method("good_nested"),
			path:        "/a/b",
			wantReports: 0,
		},
		{
			description: "failed/empty",
			expr:        service.Method("bad_empty"),
			path:        "",
			wantReports: 1,
		},
		{
			description: "failed/nested",
			expr:        service.Method("bad_nested"),
			path:        "a/b",
			wantReports: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			snaps := snaps.WithConfig(snaps.Dir("testdata"))

			logger := log.New(os.Stdout, "", 0)
			cfg := http_path_naming_convention.NewConfig()
			rule := http_path_naming_convention.NewRule(logger, cfg)

			// when
			got := rule.WalkPath(tc.expr, tc.path)
			assert.Len(t, got, tc.wantReports)
			snaps.MatchSnapshot(t, got.String())
		})
	}
}
