package lint

import (
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/runner"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("github.com/NagayamaRyoga/goa-lint-plugin", "gen", Prepare, Generate)
}

func Prepare(genpkg string, roots []eval.Root) error {
	cfg := config.NewConfig()

	if Configurator != nil {
		Configurator(cfg)
	}

	if cfg.Disabled {
		return nil
	}

	return runner.Run(cfg, genpkg, roots)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	return files, nil
}
