package futu

import (
	"context"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
	"github.com/hyperjiang/futu/protoid"
	"google.golang.org/protobuf/proto"
)

// TrdGetAccList 2001 - 获取交易业务账户列表
func (client *Client) TrdGetAccList(ctx context.Context, c2s *trdgetacclist.C2S) (*trdgetacclist.S2C, error) {
	c2s.UserID = proto.Uint64(client.GetUserID())

	req := &trdgetacclist.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetacclist.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetAccList, req, infra.NewProtobufChan(ch)); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-client.closed:
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return resp.GetS2C(), infra.Error(resp)
	}
}

// TrdGetFunds 2101 - 查询账户资金
func (client *Client) TrdGetFunds(ctx context.Context, c2s *trdgetfunds.C2S) (*trdgetfunds.S2C, error) {
	req := &trdgetfunds.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetfunds.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetFunds, req, infra.NewProtobufChan(ch)); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-client.closed:
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return resp.GetS2C(), infra.Error(resp)
	}
}
