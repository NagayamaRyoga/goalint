package main

import (
	"fmt"
	"os"

	_ "goa.design/examples/cellar/design"
	"github.com/NagayamaRyoga/goa-lint-plugin"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/runner"
	"goa.design/goa/v3/eval"
)

func main() {
	if err := eval.RunDSL(); err != nil {
		panic(err)
	}

	roots, err := eval.Context.Roots()
	if err != nil {
		panic(err)
	}

	cfg := config.NewConfig()
	cfg.Disabled = false

	if lint.Configurator != nil {
		lint.Configurator(cfg)
	}

	if err := runner.Run(cfg, "goa.design/examples/cellar/design", roots); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
