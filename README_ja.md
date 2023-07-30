# goalint

Goa lint plugin/CLI for Goa v3

## 導入方法

```go
// design/lint.go
package design

import (
    lint "github.com/NagayamaRyoga/goalint"
    _ "github.com/NagayamaRyoga/goalint/plugin" // goa gen 時にlintを実行する
)

var _ = lint.Configure(func(c *lint.Config) {
    // ...
})
```

## Rules

### MethodCasingConvention

メソッド名のケーシングに関するルール。

### TypeCasingConvention

`Type`, `ResultType` 名のケーシングに関するルール。

### TypeAttributeCasingConvention

`Type`, `ResultType` のアトリビュート名に関するルール。

### ResultTypeIdentifierNamingConvention

`ResultType` のIDに関するルール。

### HTTPPathCasingConvention

HTTPメソッドのパスのケーシングに関するルール。

### HTTPPathSegmentValidation

HTTPメソッドのパスに関するルール。
