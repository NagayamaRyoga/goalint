package design

import (
	lint "github.com/NagayamaRyoga/goalint"
	_ "github.com/NagayamaRyoga/goalint/plugin"
)

var _ = lint.Configure(func(c *lint.Config) {
	// ...
	c.TypeDescriptionExists.Disabled = true
	c.HTTPPathCasingConvention.WordCase = lint.SnakeCase
})
