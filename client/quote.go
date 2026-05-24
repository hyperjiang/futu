package client

import (
	"context"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetbroker"
	"github.com/hyperjiang/futu/pb/qotgetcapitaldistribution"
	"github.com/hyperjiang/futu/pb/qotgetcapitalflow"
	"github.com/hyperjiang/futu/pb/qotgetcompanyexecutivebackground"
	"github.com/hyperjiang/futu/pb/qotgetcompanyexecutives"
	"github.com/hyperjiang/futu/pb/qotgetcompanyoperationalefficiency"
	"github.com/hyperjiang/futu/pb/qotgetcompanyprofile"
	"github.com/hyperjiang/futu/pb/qotgetcorporateactionsbuybacks"
	"github.com/hyperjiang/futu/pb/qotgetcorporateactionsdividends"
	"github.com/hyperjiang/futu/pb/qotgetcorporateactionsstocksplits"
	"github.com/hyperjiang/futu/pb/qotgetdailyshortvolume"
	"github.com/hyperjiang/futu/pb/qotgetfinancialrevenuebreakdown"
	"github.com/hyperjiang/futu/pb/qotgetfinancialsearnpricehist"
	"github.com/hyperjiang/futu/pb/qotgetfinancialsearnpricemove"
	"github.com/hyperjiang/futu/pb/qotgetfinancialsstatements"
	"github.com/hyperjiang/futu/pb/qotgetfutureinfo"
	"github.com/hyperjiang/futu/pb/qotgetinsiderholderlist"
	"github.com/hyperjiang/futu/pb/qotgetinsidertradelist"
	"github.com/hyperjiang/futu/pb/qotgetipolist"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetmarketstate"
	"github.com/hyperjiang/futu/pb/qotgetoptionchain"
	"github.com/hyperjiang/futu/pb/qotgetoptionexerciseprobability"
	"github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate"
	"github.com/hyperjiang/futu/pb/qotgetoptionvolatility"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetownerplate"
	"github.com/hyperjiang/futu/pb/qotgetplatesecurity"
	"github.com/hyperjiang/futu/pb/qotgetplateset"
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
	"github.com/hyperjiang/futu/pb/qotgetreference"
	"github.com/hyperjiang/futu/pb/qotgetresearchanalystconsensus"
	"github.com/hyperjiang/futu/pb/qotgetresearchmorningstarrpt"
	"github.com/hyperjiang/futu/pb/qotgetresearchratingsummary"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersholderdetail"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersholdingchanges"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersinstitutional"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersoverview"
	"github.com/hyperjiang/futu/pb/qotgetshortinterest"
	"github.com/hyperjiang/futu/pb/qotgetstaticinfo"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgettoptenbuysellbrokers"
	"github.com/hyperjiang/futu/pb/qotgetusersecurity"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetvaluationdetail"
	"github.com/hyperjiang/futu/pb/qotgetvaluationplatestocklist"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotmodifyusersecurity"
	"github.com/hyperjiang/futu/pb/qotoptionscreen"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotsetpricereminder"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotstockscreen"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/pb/qotwarrantscreen"
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

// QotGetKL 3006 - 获取已订阅股票的实时K线数据
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

// QotRequestHistoryKLQuota 3104 - 获取历史K线额度使用明细
func (client *Client) QotRequestHistoryKLQuota(ctx context.Context, c2s *qotrequesthistoryklquota.C2S) (*qotrequesthistoryklquota.S2C, error) {
	req := &qotrequesthistoryklquota.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequesthistoryklquota.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotRequestHistoryKLQuota, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotRequestRehab 3105 - 获取股票的复权因子
func (client *Client) QotRequestRehab(ctx context.Context, c2s *qotrequestrehab.C2S) (*qotrequestrehab.S2C, error) {
	req := &qotrequestrehab.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequestrehab.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotRequestRehab, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetStaticInfo 3202 - 获取股票静态信息
func (client *Client) QotGetStaticInfo(ctx context.Context, c2s *qotgetstaticinfo.C2S) (*qotgetstaticinfo.S2C, error) {
	req := &qotgetstaticinfo.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetstaticinfo.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetStaticInfo, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetPlateSet 3204 - 获取板块列表
func (client *Client) QotGetPlateSet(ctx context.Context, c2s *qotgetplateset.C2S) (*qotgetplateset.S2C, error) {
	req := &qotgetplateset.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetplateset.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetPlateSet, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetPlateSecurity 3205 - 获取指定板块内的股票列表，获取股指的成分股
func (client *Client) QotGetPlateSecurity(ctx context.Context, c2s *qotgetplatesecurity.C2S) (*qotgetplatesecurity.S2C, error) {
	req := &qotgetplatesecurity.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetplatesecurity.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetPlateSecurity, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetReference 3206 - 获取证券的关联数据，如：获取正股相关窝轮、获取期货相关合约
func (client *Client) QotGetReference(ctx context.Context, c2s *qotgetreference.C2S) (*qotgetreference.S2C, error) {
	req := &qotgetreference.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetreference.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetReference, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOwnerPlate 3207 - 获取股票所属板块
func (client *Client) QotGetOwnerPlate(ctx context.Context, c2s *qotgetownerplate.C2S) (*qotgetownerplate.S2C, error) {
	req := &qotgetownerplate.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetownerplate.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOwnerPlate, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionChain 3209 - 获取期权链
func (client *Client) QotGetOptionChain(ctx context.Context, c2s *qotgetoptionchain.C2S) (*qotgetoptionchain.S2C, error) {
	req := &qotgetoptionchain.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionchain.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionChain, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetWarrant 3210 - 筛选窝轮（仅用于香港市场）
func (client *Client) QotGetWarrant(ctx context.Context, c2s *qotgetwarrant.C2S) (*qotgetwarrant.S2C, error) {
	req := &qotgetwarrant.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetwarrant.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetWarrant, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCapitalFlow 3211 - 获取资金流向
func (client *Client) QotGetCapitalFlow(ctx context.Context, c2s *qotgetcapitalflow.C2S) (*qotgetcapitalflow.S2C, error) {
	req := &qotgetcapitalflow.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetcapitalflow.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCapitalFlow, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCapitalDistribution 3212 - 获取资金分布
func (client *Client) QotGetCapitalDistribution(ctx context.Context, c2s *qotgetcapitaldistribution.C2S) (*qotgetcapitaldistribution.S2C, error) {
	req := &qotgetcapitaldistribution.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetcapitaldistribution.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCapitalDistribution, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetUserSecurity 3213 - 获取指定分组的自选股列表
func (client *Client) QotGetUserSecurity(ctx context.Context, c2s *qotgetusersecurity.C2S) (*qotgetusersecurity.S2C, error) {
	req := &qotgetusersecurity.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetusersecurity.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetUserSecurity, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotModifyUserSecurity 3214 - 修改自选股分组下的股票，该接口的S2C返回的是空
func (client *Client) QotModifyUserSecurity(ctx context.Context, c2s *qotmodifyusersecurity.C2S) error {
	req := &qotmodifyusersecurity.Request{
		C2S: c2s,
	}

	ch := make(chan *qotmodifyusersecurity.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotModifyUserSecurity, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIpoList 3217 - 获取指定市场的IPO信息
func (client *Client) QotGetIpoList(ctx context.Context, c2s *qotgetipolist.C2S) (*qotgetipolist.S2C, error) {
	req := &qotgetipolist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetipolist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIpoList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFutureInfo 3218 - 获取期货合约资料
func (client *Client) QotGetFutureInfo(ctx context.Context, c2s *qotgetfutureinfo.C2S) (*qotgetfutureinfo.S2C, error) {
	req := &qotgetfutureinfo.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetfutureinfo.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFutureInfo, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotRequestTradeDate 3219 - 请求指定市场/指定标的的交易日历
func (client *Client) QotRequestTradeDate(ctx context.Context, c2s *qotrequesttradedate.C2S) (*qotrequesttradedate.S2C, error) {
	req := &qotrequesttradedate.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequesttradedate.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotRequestTradeDate, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotSetPriceReminder 3220 - 新增、删除、修改、启用、禁用指定股票的到价提醒
func (client *Client) QotSetPriceReminder(ctx context.Context, c2s *qotsetpricereminder.C2S) (*qotsetpricereminder.S2C, error) {
	req := &qotsetpricereminder.Request{
		C2S: c2s,
	}

	ch := make(chan *qotsetpricereminder.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotSetPriceReminder, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetPriceReminder 3221 - 获取对指定股票/指定市场设置的到价提醒列表
func (client *Client) QotGetPriceReminder(ctx context.Context, c2s *qotgetpricereminder.C2S) (*qotgetpricereminder.S2C, error) {
	req := &qotgetpricereminder.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetpricereminder.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetPriceReminder, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetUserSecurityGroup 3222 - 获取自选股分组列表
func (client *Client) QotGetUserSecurityGroup(ctx context.Context, c2s *qotgetusersecuritygroup.C2S) (*qotgetusersecuritygroup.S2C, error) {
	req := &qotgetusersecuritygroup.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetusersecuritygroup.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetUserSecurityGroup, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetMarketState 3223 - 获取指定标的的市场状态
func (client *Client) QotGetMarketState(ctx context.Context, c2s *qotgetmarketstate.C2S) (*qotgetmarketstate.S2C, error) {
	req := &qotgetmarketstate.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetmarketstate.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetMarketState, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionExpirationDate 3224 - 获取期权到期日
func (client *Client) QotGetOptionExpirationDate(ctx context.Context, c2s *qotgetoptionexpirationdate.C2S) (*qotgetoptionexpirationdate.S2C, error) {
	req := &qotgetoptionexpirationdate.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionexpirationdate.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionExpirationDate, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFinancialsEarningsPriceMove 3225 - 获取财报盈利预测变动
func (client *Client) QotGetFinancialsEarningsPriceMove(ctx context.Context, c2s *qotgetfinancialsearnpricemove.C2S) (*qotgetfinancialsearnpricemove.S2C, error) {
	req := &qotgetfinancialsearnpricemove.Request{C2S: c2s}
	ch := make(chan *qotgetfinancialsearnpricemove.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFinancialsEarningsPriceMove, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFinancialsEarningsPriceHistory 3226 - 获取财报盈利预测历史
func (client *Client) QotGetFinancialsEarningsPriceHistory(ctx context.Context, c2s *qotgetfinancialsearnpricehist.C2S) (*qotgetfinancialsearnpricehist.S2C, error) {
	req := &qotgetfinancialsearnpricehist.Request{C2S: c2s}
	ch := make(chan *qotgetfinancialsearnpricehist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFinancialsEarningsPriceHistory, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFinancialsStatements 3227 - 获取财务报表
func (client *Client) QotGetFinancialsStatements(ctx context.Context, c2s *qotgetfinancialsstatements.C2S) (*qotgetfinancialsstatements.S2C, error) {
	req := &qotgetfinancialsstatements.Request{C2S: c2s}
	ch := make(chan *qotgetfinancialsstatements.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFinancialsStatements, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFinancialsRevenueBreakdown 3228 - 获取主营构成
func (client *Client) QotGetFinancialsRevenueBreakdown(ctx context.Context, c2s *qotgetfinancialrevenuebreakdown.C2S) (*qotgetfinancialrevenuebreakdown.S2C, error) {
	req := &qotgetfinancialrevenuebreakdown.Request{C2S: c2s}
	ch := make(chan *qotgetfinancialrevenuebreakdown.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFinancialsRevenueBreakdown, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetResearchAnalystConsensus 3229 - 获取分析师评级汇总
func (client *Client) QotGetResearchAnalystConsensus(ctx context.Context, c2s *qotgetresearchanalystconsensus.C2S) (*qotgetresearchanalystconsensus.S2C, error) {
	req := &qotgetresearchanalystconsensus.Request{C2S: c2s}
	ch := make(chan *qotgetresearchanalystconsensus.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetResearchAnalystConsensus, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetResearchRatingSummary 3230 - 获取分析师评级详情
func (client *Client) QotGetResearchRatingSummary(ctx context.Context, c2s *qotgetresearchratingsummary.C2S) (*qotgetresearchratingsummary.S2C, error) {
	req := &qotgetresearchratingsummary.Request{C2S: c2s}
	ch := make(chan *qotgetresearchratingsummary.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetResearchRatingSummary, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetResearchMorningstarReport 3231 - 获取晨星研报
func (client *Client) QotGetResearchMorningstarReport(ctx context.Context, c2s *qotgetresearchmorningstarrpt.C2S) (*qotgetresearchmorningstarrpt.S2C, error) {
	req := &qotgetresearchmorningstarrpt.Request{C2S: c2s}
	ch := make(chan *qotgetresearchmorningstarrpt.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetResearchMorningstarReport, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetValuationDetail 3232 - 获取估值详情
func (client *Client) QotGetValuationDetail(ctx context.Context, c2s *qotgetvaluationdetail.C2S) (*qotgetvaluationdetail.S2C, error) {
	req := &qotgetvaluationdetail.Request{C2S: c2s}
	ch := make(chan *qotgetvaluationdetail.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetValuationDetail, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetValuationPlateStockList 3233 - 获取板块估值股票列表
func (client *Client) QotGetValuationPlateStockList(ctx context.Context, c2s *qotgetvaluationplatestocklist.C2S) (*qotgetvaluationplatestocklist.S2C, error) {
	req := &qotgetvaluationplatestocklist.Request{C2S: c2s}
	ch := make(chan *qotgetvaluationplatestocklist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetValuationPlateStockList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCorporateActionsDividends 3234 - 获取除权除息
func (client *Client) QotGetCorporateActionsDividends(ctx context.Context, c2s *qotgetcorporateactionsdividends.C2S) (*qotgetcorporateactionsdividends.S2C, error) {
	req := &qotgetcorporateactionsdividends.Request{C2S: c2s}
	ch := make(chan *qotgetcorporateactionsdividends.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCorporateActionsDividends, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCorporateActionsBuybacks 3235 - 获取回购
func (client *Client) QotGetCorporateActionsBuybacks(ctx context.Context, c2s *qotgetcorporateactionsbuybacks.C2S) (*qotgetcorporateactionsbuybacks.S2C, error) {
	req := &qotgetcorporateactionsbuybacks.Request{C2S: c2s}
	ch := make(chan *qotgetcorporateactionsbuybacks.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCorporateActionsBuybacks, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCorporateActionsStockSplits 3236 - 获取拆合股
func (client *Client) QotGetCorporateActionsStockSplits(ctx context.Context, c2s *qotgetcorporateactionsstocksplits.C2S) (*qotgetcorporateactionsstocksplits.S2C, error) {
	req := &qotgetcorporateactionsstocksplits.Request{C2S: c2s}
	ch := make(chan *qotgetcorporateactionsstocksplits.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCorporateActionsStockSplits, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetShareholdersOverview 3237 - 获取股东概要
func (client *Client) QotGetShareholdersOverview(ctx context.Context, c2s *qotgetshareholdersoverview.C2S) (*qotgetshareholdersoverview.S2C, error) {
	req := &qotgetshareholdersoverview.Request{C2S: c2s}
	ch := make(chan *qotgetshareholdersoverview.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetShareholdersOverview, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetShareholdersHoldingChanges 3238 - 获取持股变动
func (client *Client) QotGetShareholdersHoldingChanges(ctx context.Context, c2s *qotgetshareholdersholdingchanges.C2S) (*qotgetshareholdersholdingchanges.S2C, error) {
	req := &qotgetshareholdersholdingchanges.Request{C2S: c2s}
	ch := make(chan *qotgetshareholdersholdingchanges.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetShareholdersHoldingChanges, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetShareholdersHolderDetail 3239 - 获取股东持仓明细
func (client *Client) QotGetShareholdersHolderDetail(ctx context.Context, c2s *qotgetshareholdersholderdetail.C2S) (*qotgetshareholdersholderdetail.S2C, error) {
	req := &qotgetshareholdersholderdetail.Request{C2S: c2s}
	ch := make(chan *qotgetshareholdersholderdetail.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetShareholdersHolderDetail, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetShareholdersInstitutional 3240 - 获取机构持仓
func (client *Client) QotGetShareholdersInstitutional(ctx context.Context, c2s *qotgetshareholdersinstitutional.C2S) (*qotgetshareholdersinstitutional.S2C, error) {
	req := &qotgetshareholdersinstitutional.Request{C2S: c2s}
	ch := make(chan *qotgetshareholdersinstitutional.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetShareholdersInstitutional, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInsiderHolderList 3241 - 获取内部持有人列表
func (client *Client) QotGetInsiderHolderList(ctx context.Context, c2s *qotgetinsiderholderlist.C2S) (*qotgetinsiderholderlist.S2C, error) {
	req := &qotgetinsiderholderlist.Request{C2S: c2s}
	ch := make(chan *qotgetinsiderholderlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInsiderHolderList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInsiderTradeList 3242 - 获取内部人交易列表
func (client *Client) QotGetInsiderTradeList(ctx context.Context, c2s *qotgetinsidertradelist.C2S) (*qotgetinsidertradelist.S2C, error) {
	req := &qotgetinsidertradelist.Request{C2S: c2s}
	ch := make(chan *qotgetinsidertradelist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInsiderTradeList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCompanyProfile 3243 - 获取公司资料
func (client *Client) QotGetCompanyProfile(ctx context.Context, c2s *qotgetcompanyprofile.C2S) (*qotgetcompanyprofile.S2C, error) {
	req := &qotgetcompanyprofile.Request{C2S: c2s}
	ch := make(chan *qotgetcompanyprofile.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCompanyProfile, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCompanyExecutives 3244 - 获取公司高管
func (client *Client) QotGetCompanyExecutives(ctx context.Context, c2s *qotgetcompanyexecutives.C2S) (*qotgetcompanyexecutives.S2C, error) {
	req := &qotgetcompanyexecutives.Request{C2S: c2s}
	ch := make(chan *qotgetcompanyexecutives.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCompanyExecutives, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCompanyExecutiveBackground 3245 - 获取高管背景
func (client *Client) QotGetCompanyExecutiveBackground(ctx context.Context, c2s *qotgetcompanyexecutivebackground.C2S) (*qotgetcompanyexecutivebackground.S2C, error) {
	req := &qotgetcompanyexecutivebackground.Request{C2S: c2s}
	ch := make(chan *qotgetcompanyexecutivebackground.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCompanyExecutiveBackground, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetCompanyOperationalEfficiency 3246 - 获取营运效率
func (client *Client) QotGetCompanyOperationalEfficiency(ctx context.Context, c2s *qotgetcompanyoperationalefficiency.C2S) (*qotgetcompanyoperationalefficiency.S2C, error) {
	req := &qotgetcompanyoperationalefficiency.Request{C2S: c2s}
	ch := make(chan *qotgetcompanyoperationalefficiency.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetCompanyOperationalEfficiency, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetTopTenBuySellBrokers 3247 - 获取十大经纪商
func (client *Client) QotGetTopTenBuySellBrokers(ctx context.Context, c2s *qotgettoptenbuysellbrokers.C2S) (*qotgettoptenbuysellbrokers.S2C, error) {
	req := &qotgettoptenbuysellbrokers.Request{C2S: c2s}
	ch := make(chan *qotgettoptenbuysellbrokers.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetTopTenBuySellBrokers, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetDailyShortVolume 3248 - 获取每日做空量
func (client *Client) QotGetDailyShortVolume(ctx context.Context, c2s *qotgetdailyshortvolume.C2S) (*qotgetdailyshortvolume.S2C, error) {
	req := &qotgetdailyshortvolume.Request{C2S: c2s}
	ch := make(chan *qotgetdailyshortvolume.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetDailyShortVolume, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetShortInterest 3249 - 获取做空比例
func (client *Client) QotGetShortInterest(ctx context.Context, c2s *qotgetshortinterest.C2S) (*qotgetshortinterest.S2C, error) {
	req := &qotgetshortinterest.Request{C2S: c2s}
	ch := make(chan *qotgetshortinterest.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetShortInterest, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionVolatility 3250 - 获取期权波动率
func (client *Client) QotGetOptionVolatility(ctx context.Context, c2s *qotgetoptionvolatility.C2S) (*qotgetoptionvolatility.S2C, error) {
	req := &qotgetoptionvolatility.Request{C2S: c2s}
	ch := make(chan *qotgetoptionvolatility.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionVolatility, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionExerciseProbability 3251 - 获取期权行权概率
func (client *Client) QotGetOptionExerciseProbability(ctx context.Context, c2s *qotgetoptionexerciseprobability.C2S) (*qotgetoptionexerciseprobability.S2C, error) {
	req := &qotgetoptionexerciseprobability.Request{C2S: c2s}
	ch := make(chan *qotgetoptionexerciseprobability.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionExerciseProbability, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotStockScreen 3252 - 条件选股
func (client *Client) QotStockScreen(ctx context.Context, c2s *qotstockscreen.C2S) (*qotstockscreen.S2C, error) {
	req := &qotstockscreen.Request{C2S: c2s}
	ch := make(chan *qotstockscreen.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotStockScreen, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotOptionScreen 3253 - 期权筛选
func (client *Client) QotOptionScreen(ctx context.Context, c2s *qotoptionscreen.C2S) (*qotoptionscreen.S2C, error) {
	req := &qotoptionscreen.Request{C2S: c2s}
	ch := make(chan *qotoptionscreen.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotOptionScreen, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotWarrantScreen 3254 - 窝轮筛选
func (client *Client) QotWarrantScreen(ctx context.Context, c2s *qotwarrantscreen.C2S) (*qotwarrantscreen.S2C, error) {
	req := &qotwarrantscreen.Request{C2S: c2s}
	ch := make(chan *qotwarrantscreen.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotWarrantScreen, req, infra.NewProtobufChan(ch)); err != nil {
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
