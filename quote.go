package futu

import (
	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/protoid"
)

// QotSub 3001 - 订阅或者反订阅，该接口的S2C返回的是空
func (client *Client) QotSub(c2s *qotsub.C2S) error {
	req := &qotsub.Request{
		C2S: c2s,
	}

	ch := make(chan *qotsub.Response)
	if err := client.Request(protoid.QotSub, req, infra.NewProtobufChan(ch)); err != nil {
		return err
	}

	select {
	case <-client.closed:
		return ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return ErrChannelClosed
		}
		close(ch)
		return infra.Error(resp)
	}
}

// QotGetBasicQot 3004 - 获取股票基本报价
func (client *Client) QotGetBasicQot(c2s *qotgetbasicqot.C2S) (*qotgetbasicqot.S2C, error) {
	req := &qotgetbasicqot.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetbasicqot.Response)
	if err := client.Request(protoid.QotGetBasicQot, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotRequestHistoryKL 3103 - 在线获取单只股票一段历史K线
func (client *Client) QotRequestHistoryKL(c2s *qotrequesthistorykl.C2S) (*qotrequesthistorykl.S2C, error) {
	req := &qotrequesthistorykl.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequesthistorykl.Response)
	if err := client.Request(protoid.QotRequestHistoryKL, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetSecuritySnapshot 3203 - 获取股票快照
func (client *Client) QotGetSecuritySnapshot(c2s *qotgetsecuritysnapshot.C2S) (*qotgetsecuritysnapshot.S2C, error) {
	req := &qotgetsecuritysnapshot.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetsecuritysnapshot.Response)
	if err := client.Request(protoid.QotGetSecuritySnapshot, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotStockFilter 3215 - 获取条件选股
func (client *Client) QotStockFilter(c2s *qotstockfilter.C2S) (*qotstockfilter.S2C, error) {
	req := &qotstockfilter.Request{
		C2S: c2s,
	}

	ch := make(chan *qotstockfilter.Response)
	if err := client.Request(protoid.QotStockFilter, req, infra.NewProtobufChan(ch)); err != nil {
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
