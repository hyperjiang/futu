package futu

import (
	"context"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/common"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
	"github.com/hyperjiang/futu/pb/trdgetorderlist"
	"github.com/hyperjiang/futu/pb/trdmodifyorder"
	"github.com/hyperjiang/futu/pb/trdplaceorder"
	"github.com/hyperjiang/futu/pb/trdunlocktrade"
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

// TrdUnlockTrade 2005 - 解锁或锁定交易
func (client *Client) TrdUnlockTrade(ctx context.Context, c2s *trdunlocktrade.C2S) (*trdunlocktrade.S2C, error) {
	req := &trdunlocktrade.Request{
		C2S: c2s,
	}

	ch := make(chan *trdunlocktrade.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdUnlockTrade, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetOrderList 2101 - 查询未完成订单
func (client *Client) TrdGetOrderList(ctx context.Context, c2s *trdgetorderlist.C2S) (*trdgetorderlist.S2C, error) {
	req := &trdgetorderlist.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetorderlist.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetOrderList, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdPlaceOrder 2202 - 下单
func (client *Client) TrdPlaceOrder(ctx context.Context, c2s *trdplaceorder.C2S) (*trdplaceorder.S2C, error) {
	if c2s.GetPacketID() == nil {
		c2s.PacketID = &common.PacketID{
			ConnID:   proto.Uint64(client.GetConnID()),
			SerialNo: proto.Uint32(client.nextSN()),
		}
	}

	req := &trdplaceorder.Request{
		C2S: c2s,
	}

	ch := make(chan *trdplaceorder.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdPlaceOrder, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdModifyOrder 2205 - 修改订单
func (client *Client) TrdModifyOrder(ctx context.Context, c2s *trdmodifyorder.C2S) (*trdmodifyorder.S2C, error) {
	if c2s.GetPacketID() == nil {
		c2s.PacketID = &common.PacketID{
			ConnID:   proto.Uint64(client.GetConnID()),
			SerialNo: proto.Uint32(client.nextSN()),
		}
	}

	req := &trdmodifyorder.Request{
		C2S: c2s,
	}

	ch := make(chan *trdmodifyorder.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdModifyOrder, req, infra.NewProtobufChan(ch)); err != nil {
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
