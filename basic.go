package futu

import (
	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/protoid"
	"google.golang.org/protobuf/proto"
)

// GetGlobalState gets the global state.
func (client *Client) GetGlobalState() (*getglobalstate.S2C, error) {
	req := &getglobalstate.Request{
		C2S: &getglobalstate.C2S{
			UserID: proto.Uint64(client.userID),
		},
	}
	ch := make(chan *getglobalstate.Response)

	if err := client.Request(protoid.GetGlobalState, req, infra.NewProtobufChan(ch)); err != nil {
		return nil, err
	}

	select {
	case <-client.closed:
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		close(ch)
		return resp.GetS2C(), infra.Error(resp)
	}
}
