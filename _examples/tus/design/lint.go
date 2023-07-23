package design

import (
	lint "github.com/NagayamaRyoga/goa-lint-plugin"
)

var _ = lint.Configure(func(c *lint.Config) {
	// ...
	c.TypeCasingConvention.Initialisms = lint.Initialisms{"TUS"}
})
