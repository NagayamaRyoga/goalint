//go:build tools

package examples

import (
	_ "goa.design/goa/v3/cmd/goa"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
