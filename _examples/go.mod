module goa.design/examples

go 1.20

replace (
	github.com/NagayamaRyoga/goalint => ../

	// https://github.com/goadesign/goa/issues/3309
	github.com/smartystreets/assertions v1.15.1 => github.com/smarty/assertions v1.13.0
)

require (
	github.com/NagayamaRyoga/goalint v0.0.0-00010101000000-000000000000
	goa.design/goa/v3 v3.12.4
	google.golang.org/grpc v1.57.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/AnatolyRugalev/goregen v0.1.0 // indirect
	github.com/dimfeld/httppath v0.0.0-20170720192232-ee938bf73598 // indirect
	github.com/dimfeld/httptreemux/v5 v5.5.0 // indirect
	github.com/ettle/strcase v0.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/manveru/faker v0.0.0-20171103152722-9fbc68a78c4d // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	golang.org/x/tools v0.12.1-0.20230815132531-74c255bcf846 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
