# goalint

Goa lint plugin/CLI for Goa v3

## 導入方法

```go
// design/goalint.go
package design

import (
    "github.com/NagayamaRyoga/goalint"
    _ "github.com/NagayamaRyoga/goalint/plugin" // goa gen 時にlintを実行する
)

var _ = goalint.Configure(func(c *goalint.Config) {
    // ...
})
```

## Rules

### APITitleExists

`API` に `Title` が存在することを確認するルール。

```go
// Bad
var _ = API("api", func() {
})
// Good
var _ = API("api", func() {
	Title("API Title")
})
```

### APIDescriptionExists

`API` に `Description` が存在することを確認するルール。

```go
// Bad
var _ = API("api", func() {
})
// Good
var _ = API("api", func() {
	Description("Description about api")
})
```

### ServiceDescriptionExists

`Service` に `Description` が存在することを確認するルール。

```go
// Bad
var _ = Service("bad_service", func() {
})
// Good
var _ = Service("good_service", func() {
	Description("Description about good_service")
})
```

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

### MethodDescriptionExists

`Method` に `Description` が存在することを確認するルール。

```go
var _ = Service("service", func() {
	// Bad
	Method("getBadExample", func() {})
	// Good
	Method("get_good_example", func() {
		Description("Description about get_good_example")
	})
})
```

### MethodArrayResult

`Method` の `Result` が直接配列を返すことを禁止するルール。

ref: https://github.com/IBM/openapi-validator/blob/main/docs/ibm-cloud-rules.md#ibm-no-array-responses

```go
var _ = Service("service", func() {
	Method("list_titles", func() {
		// Bad
		Result(ArrayOf(String))
		// Good
		Result(ListTitlesResponse)
	})
})
var ListTitlesResponse = Type("ListTitlesResponse", func() {
	Required("titles")
	Attribute("titles", ArrayOf(String))
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

### NoUnnamedMethodResultType

`Method` の `Result` に無名型を使用することを禁止するルール。

```go
var _ = Service("service", func() {
	// Bad
	Method("bad", func() {
		Result(func() {
			Attribute("a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
	})
	// Good
	Method("good", func() {
		Result(GoodResponse)
	})
})

var GoodResponse = Type("GoodResponse", ...)
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

### TypeAttributeDescriptionExists

`Type`, `ResultType` の `Attribute` に `Description` が存在することを確認するルール

```go
var _ = Type("Type", func() {
	// Bad
	Attribute("bad_attr", Int)
	// Good
	Attribute("good_attr", Int, func() {
		Description("Description about good_attr")
	})
})
```

### TypeAttributeExampleExists

`Type`, `ResultType` の `Attribute` に `Example` が存在することを確認するルール

```go
var _ = Type("Type", func() {
	// Bad
	Attribute("bad_attr", Int)
	// Good
	Attribute("good_attr", Int, func() {
		Example(42)
	})
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

### HTTPErrorDescriptionExists

HTTPエラーの `Response` に `Description` がが存在することを確認するルール

```go
var _ = Service("service", func() {
	Method("method", func() {
		HTTP(func() {
			GET("/")
            Response(StatusOK)
			// Bad
            Response("not_found", StatusNotFound)
			// Good
            Response("not_found", StatusNotFound, func() {
                Description("not found")
            })
		})
	})
    Error("not_found")
})
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

### HTTPPathNamingConvention

HTTPメソッドのパスの命名に関するルール。

```go
var _ = Service("service", func() {
	Method("method", func() {
		HTTP(func() {
			// Bad
			GET("bad/path")
			// Good
			GET("/good/path")
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
