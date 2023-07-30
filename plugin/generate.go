package lintplugin

import (
	lint "github.com/NagayamaRyoga/goalint"
	"github.com/NagayamaRyoga/goalint/pkg/config"
	"github.com/NagayamaRyoga/goalint/pkg/runner"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("github.com/NagayamaRyoga/goalint", "gen", Prepare, Generate)
}

func Prepare(genpkg string, roots []eval.Root) error {
	cfg := config.NewConfig()

	if lint.Configurator != nil {
		lint.Configurator(cfg)
	}

	if cfg.Disabled {
		return nil
	}

	return runner.Run(cfg, genpkg, roots)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	return files, nil
}
