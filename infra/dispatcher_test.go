package infra_test

import (
	"fmt"
	"testing"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/common"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestDispatcher(t *testing.T) {
	should := require.New(t)

	ch := make(chan *common.PacketID)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("first chan closed")
	}()

	ch2 := make(chan *common.PacketID)
	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
		fmt.Println("second chan closed")
	}()

	hub := infra.NewDispatcherHub()
	hub.Register(1001, 1, infra.NewProtobufChan(ch))
	hub.Register(1001, 1, infra.NewProtobufChan(ch2)) // replace the old chan
	hub.Register(1001, 3, infra.NewProtobufChan(ch2))
	err := hub.Dispatch(1001, 1, []byte{1, 2, 3})
	should.Error(err)

	p := &common.PacketID{ConnID: proto.Uint64(333), SerialNo: proto.Uint32(333)}
	b, err := proto.Marshal(p)
	should.NoError(err)
	err = hub.Dispatch(1001, 1, b)
	should.NoError(err)
	err = hub.Dispatch(1001, 2, b) // channel not found
	should.Error(err)
	err = hub.Dispatch(1002, 1, b) // dispatcher not found
	should.Error(err)

	hub.Register(1001, 0, infra.NewProtobufChan(ch2))
	err = hub.Dispatch(1001, 3, b) // fallback to the default channel (sn = 0)
	should.NoError(err)

	close(ch2)
	err = hub.Dispatch(1001, 3, b)
	should.NoError(err) // the panic is caught

	p2 := &common.PacketID{ConnID: proto.Uint64(444), SerialNo: proto.Uint32(444)}
	b2, err := proto.Marshal(p2)
	should.NoError(err)

	ch3 := make(chan *common.PacketID)
	hub.Register(1001, 0, infra.NewProtobufChan(ch3))

	err = hub.Dispatch(1001, 4, b2)
	should.Error(err) // blocked and timeout

	hub.Close()
}
