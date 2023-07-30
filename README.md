# goalint

- [日本語](README_ja.md)

Goa lint plugin/CLI for Goa v3

## Usage

```go
// design/lint.go
package design

import (
    lint "github.com/NagayamaRyoga/goalint"
    _ "github.com/NagayamaRyoga/goalint/plugin"
)

var _ = lint.Configure(func(c *lint.Config) {
    // ...
})
```

## Rules

### MethodCasingConvention

### TypeCasingConvention

### TypeAttributeCasingConvention

### ResultTypeIdentifierNamingConvention

### HTTPPathCasingConvention

### HTTPPathSegmentValidation
