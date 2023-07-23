# goa-lint-plugin

## Usage

```go
// design/lint.go
package design

import (
    lint "github.com/NagayamaRyoga/goa-lint-plugin"
)

var _ = lint.Configure(func(c *lint.Config) {
    // ...
})
```
