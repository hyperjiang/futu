package infra_test

import (
	"testing"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/keepalive"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestError(t *testing.T) {
	should := require.New(t)

	r := &keepalive.Response{
		RetType: proto.Int32(0),
		RetMsg:  proto.String("abc"),
	}
	should.Nil(infra.Error(r))

	r.RetType = proto.Int32(1)
	should.Error(infra.Error(r))
}
