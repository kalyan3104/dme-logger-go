//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf  --gogoslick_out=proto testdata.proto
package marshal

import "github.com/kalyan3104/dme-logger-go/marshal/proto"

// TestData -
type TestData struct {
	proto.TestData
}
