package design

import (
	"github.com/NagayamaRyoga/goalint"
	_ "github.com/NagayamaRyoga/goalint/plugin"
)

var _ = goalint.Configure(func(c *goalint.Config) {
	// ...
})
