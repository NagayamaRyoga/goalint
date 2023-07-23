package lint

import (
	"fmt"

	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/rules"
	"go.uber.org/multierr"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("github.com/NagayamaRyoga/goa-lint-plugin", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	var merr error

	fmt.Println("genpkg:", genpkg)

	for _, rule := range rules.Rules {
		if err := rule.Apply(cfg, roots); err != nil {
			merr = multierr.Append(merr, err)
		}
	}

	return files, merr
}
