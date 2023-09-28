package marshal

import (
	gproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

// GogoProtoObj groups the necessary of a gogo protobuf marshalizeble object
type GogoProtoObj interface {
	gproto.Marshaler
	gproto.Unmarshaler
	proto.Message
}
