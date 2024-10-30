package futu

import (
	"context"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/common"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
	"github.com/hyperjiang/futu/pb/trdgethistoryorderfilllist"
	"github.com/hyperjiang/futu/pb/trdgethistoryorderlist"
	"github.com/hyperjiang/futu/pb/trdgetmarginratio"
	"github.com/hyperjiang/futu/pb/trdgetmaxtrdqtys"
	"github.com/hyperjiang/futu/pb/trdgetorderfee"
	"github.com/hyperjiang/futu/pb/trdgetorderfilllist"
	"github.com/hyperjiang/futu/pb/trdgetorderlist"
	"github.com/hyperjiang/futu/pb/trdgetpositionlist"
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

// TrdGetPositionList 2102 - 查询持仓
func (client *Client) TrdGetPositionList(ctx context.Context, c2s *trdgetpositionlist.C2S) (*trdgetpositionlist.S2C, error) {
	req := &trdgetpositionlist.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetpositionlist.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetPositionList, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetMaxTrdQtys 2111 - 查询指定证券的最大可买可卖数量
func (client *Client) TrdGetMaxTrdQtys(ctx context.Context, c2s *trdgetmaxtrdqtys.C2S) (*trdgetmaxtrdqtys.S2C, error) {
	req := &trdgetmaxtrdqtys.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetmaxtrdqtys.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetMaxTrdQtys, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetOrderList 2201 - 查询未完成订单
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

// TrdGetOrderFillList 2211 - 获取成交列表
func (client *Client) TrdGetOrderFillList(ctx context.Context, c2s *trdgetorderfilllist.C2S) (*trdgetorderfilllist.S2C, error) {
	req := &trdgetorderfilllist.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetorderfilllist.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetOrderFillList, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetHistoryOrderList 2221 - 获取历史订单列表
func (client *Client) TrdGetHistoryOrderList(ctx context.Context, c2s *trdgethistoryorderlist.C2S) (*trdgethistoryorderlist.S2C, error) {
	req := &trdgethistoryorderlist.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgethistoryorderlist.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetHistoryOrderList, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetHistoryOrderFillList 2222 - 获取历史成交列表
func (client *Client) TrdGetHistoryOrderFillList(ctx context.Context, c2s *trdgethistoryorderfilllist.C2S) (*trdgethistoryorderfilllist.S2C, error) {
	req := &trdgethistoryorderfilllist.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgethistoryorderfilllist.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetHistoryOrderFillList, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetMarginRatio 2223 - 获取融资融券数据
func (client *Client) TrdGetMarginRatio(ctx context.Context, c2s *trdgetmarginratio.C2S) (*trdgetmarginratio.S2C, error) {
	req := &trdgetmarginratio.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetmarginratio.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetMarginRatio, req, infra.NewProtobufChan(ch)); err != nil {
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

// TrdGetOrderFee 2225 - 获取交易费用
func (client *Client) TrdGetOrderFee(ctx context.Context, c2s *trdgetorderfee.C2S) (*trdgetorderfee.S2C, error) {
	req := &trdgetorderfee.Request{
		C2S: c2s,
	}

	ch := make(chan *trdgetorderfee.Response)
	defer close(ch)
	if err := client.Request(protoid.TrdGetOrderFee, req, infra.NewProtobufChan(ch)); err != nil {
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
