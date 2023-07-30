package main

import (
	"fmt"
	"os"

	lint "github.com/NagayamaRyoga/goalint"
	"github.com/NagayamaRyoga/goalint/inner/config"
	"github.com/NagayamaRyoga/goalint/inner/runner"
	_ "goa.design/examples/cellar/design"
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
