package design

import (
	lint "github.com/NagayamaRyoga/goa-lint-plugin"
)

var _ = lint.Configure(func(c *lint.Config) {
	// ...
	c.TypeDescriptionExists.Disabled = true
	c.HTTPPathCasingConvention.WordCase = lint.SnakeCase
})
