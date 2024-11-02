package client

import (
	"context"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetbroker"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/protoid"
)

// QotSub 3001 - 订阅或者反订阅，该接口的S2C返回的是空
func (client *Client) QotSub(ctx context.Context, c2s *qotsub.C2S) error {
	req := &qotsub.Request{
		C2S: c2s,
	}

	ch := make(chan *qotsub.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotSub, req, infra.NewProtobufChan(ch)); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-client.closed:
		return ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return ErrChannelClosed
		}
		return infra.Error(resp)
	}
}

// QotGetSubInfo 3003 - 获取订阅状态
func (client *Client) QotGetSubInfo(ctx context.Context, c2s *qotgetsubinfo.C2S) (*qotgetsubinfo.S2C, error) {
	req := &qotgetsubinfo.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetsubinfo.Response, 1)
	defer close(ch)

	if err := client.Request(protoid.QotGetSubInfo, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetBasicQot 3004 - 获取已订阅股票的实时报价
func (client *Client) QotGetBasicQot(ctx context.Context, c2s *qotgetbasicqot.C2S) (*qotgetbasicqot.S2C, error) {
	req := &qotgetbasicqot.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetbasicqot.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetBasicQot, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetKL 3005 - 获取已订阅股票的实时K线数据
func (client *Client) QotGetKL(ctx context.Context, c2s *qotgetkl.C2S) (*qotgetkl.S2C, error) {
	req := &qotgetkl.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetkl.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetKL, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetRT 3008 - 获取已订阅股票的实时分时数据
func (client *Client) QotGetRT(ctx context.Context, c2s *qotgetrt.C2S) (*qotgetrt.S2C, error) {
	req := &qotgetrt.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetrt.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetRT, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetTicker 3010 - 获取已订阅股票的实时逐笔
func (client *Client) QotGetTicker(ctx context.Context, c2s *qotgetticker.C2S) (*qotgetticker.S2C, error) {
	req := &qotgetticker.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetticker.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetTicker, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOrderBook 3012 - 获取已订阅股票的实时摆盘
func (client *Client) QotGetOrderBook(ctx context.Context, c2s *qotgetorderbook.C2S) (*qotgetorderbook.S2C, error) {
	req := &qotgetorderbook.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetorderbook.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOrderBook, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetBroker 3014 - 获取已订阅股票的实时经纪队列
func (client *Client) QotGetBroker(ctx context.Context, c2s *qotgetbroker.C2S) (*qotgetbroker.S2C, error) {
	req := &qotgetbroker.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetbroker.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetBroker, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotRequestHistoryKL 3103 - 在线获取单只股票一段历史K线
func (client *Client) QotRequestHistoryKL(ctx context.Context, c2s *qotrequesthistorykl.C2S) (*qotrequesthistorykl.S2C, error) {
	req := &qotrequesthistorykl.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequesthistorykl.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotRequestHistoryKL, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetSecuritySnapshot 3203 - 获取股票快照
func (client *Client) QotGetSecuritySnapshot(ctx context.Context, c2s *qotgetsecuritysnapshot.C2S) (*qotgetsecuritysnapshot.S2C, error) {
	req := &qotgetsecuritysnapshot.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetsecuritysnapshot.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetSecuritySnapshot, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotStockFilter 3215 - 获取条件选股
func (client *Client) QotStockFilter(ctx context.Context, c2s *qotstockfilter.C2S) (*qotstockfilter.S2C, error) {
	req := &qotstockfilter.Request{
		C2S: c2s,
	}

	ch := make(chan *qotstockfilter.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotStockFilter, req, infra.NewProtobufChan(ch)); err != nil {
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
