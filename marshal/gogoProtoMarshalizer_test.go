package marshal_test

import (
	"fmt"
	"testing"

	"github.com/kalyan3104/dme-logger-go/check"
	"github.com/kalyan3104/dme-logger-go/marshal"
	"github.com/kalyan3104/dme-logger-go/marshal/proto"
	"github.com/stretchr/testify/assert"
)

var gogoMarsh = &marshal.GogoProtoMarshalizer{}

func recovedMarshal(obj interface{}) (buf []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			if panicError, ok := p.(error); ok {
				err = panicError
			} else {
				err = fmt.Errorf("%#v", p)
			}
			buf = nil
		}
	}()
	buf, err = gogoMarsh.Marshal(obj)
	return
}

func recovedUnmarshal(obj interface{}, buf []byte) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if panicError, ok := p.(error); ok {
				err = panicError
			} else {
				err = fmt.Errorf("%#v", p)
			}
		}
	}()
	err = gogoMarsh.Unmarshal(obj, buf)
	return
}

func TestGogoProtoMarshalizer_Marshal(t *testing.T) {
	t.Parallel()

	encData, err := recovedMarshal(createDummyTestData())

	assert.Nil(t, err)
	assert.NotNil(t, encData)
}

func TestGogoProtoMarshalizer_MarshalWrongObj(t *testing.T) {
	t.Parallel()

	obj := "kalyan3104"
	encData, err := recovedMarshal(obj)

	assert.Nil(t, encData)
	assert.NotNil(t, err)
}

func TestGogoProtoMarshalizer_Unmarshal(t *testing.T) {
	t.Parallel()

	data := createDummyTestData()
	encData, _ := gogoMarsh.Marshal(data)
	recoveredData := &marshal.TestData{}

	err := recovedUnmarshal(&recoveredData.TestData, encData)
	assert.False(t, check.IfNil(gogoMarsh))
	assert.Nil(t, err)
	assert.Equal(t, recoveredData, data)
}

func TestGogoProtoMarshalizer_UnmarshalWrongObj(t *testing.T) {
	t.Parallel()

	encNode, _ := gogoMarsh.Marshal(createDummyTestData())
	err := recovedUnmarshal([]byte{}, encNode)

	assert.NotNil(t, err)
}

func createDummyTestData() *marshal.TestData {
	return &marshal.TestData{
		TestData: proto.TestData{
			Hash:    []byte("hash"),
			ShardID: 2,
			Nonce:   1002,
			Hashes: [][]byte{
				[]byte("hash1"),
				[]byte("hash2"),
			},
			Message: "a message",
		},
	}

}
