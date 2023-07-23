package lint

import (
	"github.com/NagayamaRyoga/goa-lint-plugin/pkg/config"
)

type Config = config.Config

var cfg = config.NewConfig()

func Configure(f func(c *Config)) struct{} {
	f(cfg)
	return struct{}{}
}
