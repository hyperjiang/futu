package infra_test

import (
	"fmt"
	"testing"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestProtobufChan(t *testing.T) {
	should := require.New(t)

	should.Nil(infra.NewProtobufChan(nil))
	should.Nil(infra.NewProtobufChan(1))
	should.Nil(infra.NewProtobufChan("abc"))
	should.Nil(infra.NewProtobufChan(make(chan int)))
	should.Nil(infra.NewProtobufChan(make(chan common.PacketID)))

	ch := make(chan *common.PacketID)
	go func() {
		for p := range ch {
			fmt.Println(p)
		}
	}()

	pc := infra.NewProtobufChan(ch)
	should.NotNil(pc)
	should.Error(pc.Send([]byte{1, 2, 3}))

	p := &common.PacketID{ConnID: proto.Uint64(1), SerialNo: proto.Uint32(2)}
	b, err := proto.Marshal(p)
	should.NoError(err)

	should.NoError(pc.Send(b))

	pc.Close()
}
