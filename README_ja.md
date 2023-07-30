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

```go
var _ = Service("service", func() {
	// Bad
	Method("getBadExample", ...)
	// Good
	Method("get_good_example", ...)
})
```

### NoUnnamedMethodPayloadType

`Method` の `Payload` に無名型を使用することを禁止するルール。

```go
var _ = Service("service", func() {
	// Bad
	Method("bad", func() {
		Payload(func() {
			Attribute("a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
	})
	// Good
	Method("good", func() {
		Payload(GoodPayload)
	})
})

var GoodPayload = Type("GoodPayload", ...)
```

### TypeCasingConvention

`Type`, `ResultType` 名のケーシングに関するルール。

```go
// Bad
var BadType = Type("bad_type", ...)
// Good
var GoodType = Type("GoodType", ...)
```

### TypeDescriptionExists

`Type`, `ResultType` の `Description` が存在することを確認するルール。

```go
// Bad
var BadType = Type("BadType", func() {
	Attribute("a", Int)
})
// Good
var GoodType = Type("GoodType", func() {
	Description("Description about GoodType")
	Attribute("a", Int)
})
```

### TypeAttributeCasingConvention

`Type`, `ResultType` のアトリビュート名に関するルール。

```go
var _ = Type("Something", func() {
	// Bad
	Attribute("badAttribute", Int)
	// Good
	Attribute("good_attribute", Int)
})
```

### ResultTypeIdentifierNamingConvention

`ResultType` のIDに関するルール。

```go
// Bad
var BadResultType = Type("bad-result-type", ...)
// Good
var GoodResultType = Type("application/vnd.good-result-type", ...)
```

### HTTPPathCasingConvention

HTTPメソッドのパスのケーシングに関するルール。

```go
var _ = Service("service", func() {
	Method("method", func() {
		HTTP(func() {
			// Bad
			GET("/bad_path")
			// Good
			GET("/good-path")
		})
	})
})
```

### HTTPPathSegmentValidation

HTTPメソッドのパスに関するルール。

```go
var _ = Service("service", func() {
	Method("method", func() {
		HTTP(func() {
			// Bad
			GET("/b{ad}_path")
		})
	})
})
```
