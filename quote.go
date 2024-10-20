package futu

import (
	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/protoid"
)

// QotGetBasicQot 获取股票基本报价
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

// QotRequestHistoryKL 在线获取单只股票一段历史K线
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
