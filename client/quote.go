package client

import (
	"context"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/qotfiltercompetition"
	"github.com/hyperjiang/futu/pb/qotgetarkactivetransaction"
	"github.com/hyperjiang/futu/pb/qotgetarkfundholding"
	"github.com/hyperjiang/futu/pb/qotgetarkstockdynamic"
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
	"github.com/hyperjiang/futu/pb/qotgetdividendcalendar"
	"github.com/hyperjiang/futu/pb/qotgetdividendrank"
	"github.com/hyperjiang/futu/pb/qotgetearningsbeatrank"
	"github.com/hyperjiang/futu/pb/qotgetearningscalendar"
	"github.com/hyperjiang/futu/pb/qotgeteconomiccalendar"
	"github.com/hyperjiang/futu/pb/qotgeteventcontract"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractcategory"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractcombolist"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractcomborfq"
	"github.com/hyperjiang/futu/pb/qotgeteventcontracteventlist"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractkline"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractmilestonelist"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractorderbook"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractserieslist"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractsnapshot"
	"github.com/hyperjiang/futu/pb/qotgeteventcontractticker"
	"github.com/hyperjiang/futu/pb/qotgetfedwatchdotplot"
	"github.com/hyperjiang/futu/pb/qotgetfedwatchtargetrate"
	"github.com/hyperjiang/futu/pb/qotgetfinancialrevenuebreakdown"
	"github.com/hyperjiang/futu/pb/qotgetfinancialsearnpricehist"
	"github.com/hyperjiang/futu/pb/qotgetfinancialsearnpricemove"
	"github.com/hyperjiang/futu/pb/qotgetfinancialsstatements"
	"github.com/hyperjiang/futu/pb/qotgetfutureinfo"
	"github.com/hyperjiang/futu/pb/qotgetheatmapdata"
	"github.com/hyperjiang/futu/pb/qotgethighdividendsoerank"
	"github.com/hyperjiang/futu/pb/qotgethotlist"
	"github.com/hyperjiang/futu/pb/qotgetindicatorlist"
	"github.com/hyperjiang/futu/pb/qotgetindustrialchainbyplate"
	"github.com/hyperjiang/futu/pb/qotgetindustrialchaindetail"
	"github.com/hyperjiang/futu/pb/qotgetindustrialchainlist"
	"github.com/hyperjiang/futu/pb/qotgetindustrialplateinfo"
	"github.com/hyperjiang/futu/pb/qotgetindustrialplatestock"
	"github.com/hyperjiang/futu/pb/qotgetinsiderholderlist"
	"github.com/hyperjiang/futu/pb/qotgetinsidertradelist"
	"github.com/hyperjiang/futu/pb/qotgetinstitutiondistribution"
	"github.com/hyperjiang/futu/pb/qotgetinstitutionholdingchange"
	"github.com/hyperjiang/futu/pb/qotgetinstitutionholdinglist"
	"github.com/hyperjiang/futu/pb/qotgetinstitutionlist"
	"github.com/hyperjiang/futu/pb/qotgetinstitutionprofile"
	"github.com/hyperjiang/futu/pb/qotgetipolist"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetmacroindicatorhistory"
	"github.com/hyperjiang/futu/pb/qotgetmacroindicatorlist"
	"github.com/hyperjiang/futu/pb/qotgetmarketstate"
	"github.com/hyperjiang/futu/pb/qotgetoptionchain"
	"github.com/hyperjiang/futu/pb/qotgetoptionearningsscreener"
	"github.com/hyperjiang/futu/pb/qotgetoptionevent"
	"github.com/hyperjiang/futu/pb/qotgetoptioneventalert"
	"github.com/hyperjiang/futu/pb/qotgetoptionexerciseprobability"
	"github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate"
	"github.com/hyperjiang/futu/pb/qotgetoptionmarketstatistic"
	"github.com/hyperjiang/futu/pb/qotgetoptionquote"
	"github.com/hyperjiang/futu/pb/qotgetoptionrank"
	"github.com/hyperjiang/futu/pb/qotgetoptionsellerscreener"
	"github.com/hyperjiang/futu/pb/qotgetoptionstrategy"
	"github.com/hyperjiang/futu/pb/qotgetoptionstrategyanalysis"
	"github.com/hyperjiang/futu/pb/qotgetoptionstrategyspreads"
	"github.com/hyperjiang/futu/pb/qotgetoptionunderlyinghisstatistic"
	"github.com/hyperjiang/futu/pb/qotgetoptionunderlyinghisvolatility"
	"github.com/hyperjiang/futu/pb/qotgetoptionunderlyingoverview"
	"github.com/hyperjiang/futu/pb/qotgetoptionunderlyingrank"
	"github.com/hyperjiang/futu/pb/qotgetoptionvolatility"
	"github.com/hyperjiang/futu/pb/qotgetoptionzerodtecontract"
	"github.com/hyperjiang/futu/pb/qotgetoptionzerodtescreener"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetownerplate"
	"github.com/hyperjiang/futu/pb/qotgetperiodchangerank"
	"github.com/hyperjiang/futu/pb/qotgetplatesecurity"
	"github.com/hyperjiang/futu/pb/qotgetplateset"
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
	"github.com/hyperjiang/futu/pb/qotgetratingchange"
	"github.com/hyperjiang/futu/pb/qotgetreference"
	"github.com/hyperjiang/futu/pb/qotgetresearchanalystconsensus"
	"github.com/hyperjiang/futu/pb/qotgetresearchmorningstarrpt"
	"github.com/hyperjiang/futu/pb/qotgetresearchratingsummary"
	"github.com/hyperjiang/futu/pb/qotgetrisefalldistr"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsearchnews"
	"github.com/hyperjiang/futu/pb/qotgetsearchquote"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersholderdetail"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersholdingchanges"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersinstitutional"
	"github.com/hyperjiang/futu/pb/qotgetshareholdersoverview"
	"github.com/hyperjiang/futu/pb/qotgetshortinterest"
	"github.com/hyperjiang/futu/pb/qotgetshortsellingrank"
	"github.com/hyperjiang/futu/pb/qotgetstaticinfo"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgettopmoverrank"
	"github.com/hyperjiang/futu/pb/qotgettoptenbuysellbrokers"
	"github.com/hyperjiang/futu/pb/qotgetusafterhoursrank"
	"github.com/hyperjiang/futu/pb/qotgetusersecurity"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetusovernightrank"
	"github.com/hyperjiang/futu/pb/qotgetuspremarketrank"
	"github.com/hyperjiang/futu/pb/qotgetvaluationdetail"
	"github.com/hyperjiang/futu/pb/qotgetvaluationplatestocklist"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotmodifyusersecurity"
	"github.com/hyperjiang/futu/pb/qotoptionscreen"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryeventcontractkl"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestindicatorcalc"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotsetoptioneventalert"
	"github.com/hyperjiang/futu/pb/qotsetpricereminder"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotstockscreen"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/pb/qotsubeventcontract"
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

// QotGetOptionQuote 3255 - 获取期权行情
func (client *Client) QotGetOptionQuote(ctx context.Context, c2s *qotgetoptionquote.C2S) (*qotgetoptionquote.S2C, error) {
	req := &qotgetoptionquote.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionquote.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionQuote, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionStrategy 3256 - 获取期权策略
func (client *Client) QotGetOptionStrategy(ctx context.Context, c2s *qotgetoptionstrategy.C2S) (*qotgetoptionstrategy.S2C, error) {
	req := &qotgetoptionstrategy.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionstrategy.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionStrategy, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionStrategyAnalysis 3257 - 获取期权策略分析
func (client *Client) QotGetOptionStrategyAnalysis(ctx context.Context, c2s *qotgetoptionstrategyanalysis.C2S) (*qotgetoptionstrategyanalysis.S2C, error) {
	req := &qotgetoptionstrategyanalysis.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionstrategyanalysis.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionStrategyAnalysis, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionStrategySpread 3258 - 获取期权策略价差
func (client *Client) QotGetOptionStrategySpread(ctx context.Context, c2s *qotgetoptionstrategyspreads.C2S) (*qotgetoptionstrategyspreads.S2C, error) {
	req := &qotgetoptionstrategyspreads.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionstrategyspreads.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionStrategySpread, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIndicatorList 3259 - 获取指标列表
func (client *Client) QotGetIndicatorList(ctx context.Context, c2s *qotgetindicatorlist.C2S) (*qotgetindicatorlist.S2C, error) {
	req := &qotgetindicatorlist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetindicatorlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIndicatorList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotRequestIndicatorCalc 3260 - 异步发起指标计算
func (client *Client) QotRequestIndicatorCalc(ctx context.Context, c2s *qotrequestindicatorcalc.C2S) (*qotrequestindicatorcalc.S2C, error) {
	req := &qotrequestindicatorcalc.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequestindicatorcalc.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotRequestIndicatorCalc, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetSearchQuote 3262 - 搜索行情
func (client *Client) QotGetSearchQuote(ctx context.Context, c2s *qotgetsearchquote.C2S) (*qotgetsearchquote.S2C, error) {
	req := &qotgetsearchquote.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetsearchquote.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetSearchQuote, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetSearchNews 3263 - 搜索资讯
func (client *Client) QotGetSearchNews(ctx context.Context, c2s *qotgetsearchnews.C2S) (*qotgetsearchnews.S2C, error) {
	req := &qotgetsearchnews.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetsearchnews.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetSearchNews, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionMarketStatistic 3301 - 获取期权市场统计
func (client *Client) QotGetOptionMarketStatistic(ctx context.Context, c2s *qotgetoptionmarketstatistic.C2S) (*qotgetoptionmarketstatistic.S2C, error) {
	req := &qotgetoptionmarketstatistic.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionmarketstatistic.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionMarketStatistic, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionUnderlyingHisStatistic 3302 - 获取期权标的历史统计
func (client *Client) QotGetOptionUnderlyingHisStatistic(ctx context.Context, c2s *qotgetoptionunderlyinghisstatistic.C2S) (*qotgetoptionunderlyinghisstatistic.S2C, error) {
	req := &qotgetoptionunderlyinghisstatistic.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionunderlyinghisstatistic.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionUnderlyingHisStatistic, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionUnderlyingOverview 3303 - 获取批量标的最新数据
func (client *Client) QotGetOptionUnderlyingOverview(ctx context.Context, c2s *qotgetoptionunderlyingoverview.C2S) (*qotgetoptionunderlyingoverview.S2C, error) {
	req := &qotgetoptionunderlyingoverview.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionunderlyingoverview.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionUnderlyingOverview, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionUnderlyingHisVolatility 3304 - 获取历史波动率
func (client *Client) QotGetOptionUnderlyingHisVolatility(ctx context.Context, c2s *qotgetoptionunderlyinghisvolatility.C2S) (*qotgetoptionunderlyinghisvolatility.S2C, error) {
	req := &qotgetoptionunderlyinghisvolatility.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionunderlyinghisvolatility.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionUnderlyingHisVolatility, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionUnderlyingRank 3305 - 获取标的排行
func (client *Client) QotGetOptionUnderlyingRank(ctx context.Context, c2s *qotgetoptionunderlyingrank.C2S) (*qotgetoptionunderlyingrank.S2C, error) {
	req := &qotgetoptionunderlyingrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionunderlyingrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionUnderlyingRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionRank 3306 - 获取期权合约排行
func (client *Client) QotGetOptionRank(ctx context.Context, c2s *qotgetoptionrank.C2S) (*qotgetoptionrank.S2C, error) {
	req := &qotgetoptionrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionEvent 3307 - 获取期权异动列表
func (client *Client) QotGetOptionEvent(ctx context.Context, c2s *qotgetoptionevent.C2S) (*qotgetoptionevent.S2C, error) {
	req := &qotgetoptionevent.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionevent.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionEvent, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionEventAlert 3308 - 获取期权异动告警设置
func (client *Client) QotGetOptionEventAlert(ctx context.Context, c2s *qotgetoptioneventalert.C2S) (*qotgetoptioneventalert.S2C, error) {
	req := &qotgetoptioneventalert.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptioneventalert.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionEventAlert, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotSetOptionEventAlert 3309 - 修改期权异动告警条件，该接口的S2C返回的是空
func (client *Client) QotSetOptionEventAlert(ctx context.Context, c2s *qotsetoptioneventalert.C2S) error {
	req := &qotsetoptioneventalert.Request{
		C2S: c2s,
	}

	ch := make(chan *qotsetoptioneventalert.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotSetOptionEventAlert, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionZeroDteScreener 3311 - 获取末日期权标的列表
func (client *Client) QotGetOptionZeroDteScreener(ctx context.Context, c2s *qotgetoptionzerodtescreener.C2S) (*qotgetoptionzerodtescreener.S2C, error) {
	req := &qotgetoptionzerodtescreener.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionzerodtescreener.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionZeroDteScreener, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionZeroDteContract 3312 - 获取末日期权合约列表
func (client *Client) QotGetOptionZeroDteContract(ctx context.Context, c2s *qotgetoptionzerodtecontract.C2S) (*qotgetoptionzerodtecontract.S2C, error) {
	req := &qotgetoptionzerodtecontract.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionzerodtecontract.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionZeroDteContract, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionEarningsScreener 3313 - 获取财报期权标的列表
func (client *Client) QotGetOptionEarningsScreener(ctx context.Context, c2s *qotgetoptionearningsscreener.C2S) (*qotgetoptionearningsscreener.S2C, error) {
	req := &qotgetoptionearningsscreener.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionearningsscreener.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionEarningsScreener, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetOptionSellerScreener 3314 - 获取期权卖方策略列表
func (client *Client) QotGetOptionSellerScreener(ctx context.Context, c2s *qotgetoptionsellerscreener.C2S) (*qotgetoptionsellerscreener.S2C, error) {
	req := &qotgetoptionsellerscreener.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetoptionsellerscreener.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetOptionSellerScreener, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEarningsCalendar 3401 - 获取财报日历
func (client *Client) QotGetEarningsCalendar(ctx context.Context, c2s *qotgetearningscalendar.C2S) (*qotgetearningscalendar.S2C, error) {
	req := &qotgetearningscalendar.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetearningscalendar.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEarningsCalendar, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetMacroIndicatorList 3402 - 获取宏观指标列表
func (client *Client) QotGetMacroIndicatorList(ctx context.Context, c2s *qotgetmacroindicatorlist.C2S) (*qotgetmacroindicatorlist.S2C, error) {
	req := &qotgetmacroindicatorlist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetmacroindicatorlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetMacroIndicatorList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetMacroIndicatorHistory 3403 - 获取宏观指标历史数据
func (client *Client) QotGetMacroIndicatorHistory(ctx context.Context, c2s *qotgetmacroindicatorhistory.C2S) (*qotgetmacroindicatorhistory.S2C, error) {
	req := &qotgetmacroindicatorhistory.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetmacroindicatorhistory.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetMacroIndicatorHistory, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFedWatchTargetRate 3404 - 获取美联储利率预测(FedWatch概率)
func (client *Client) QotGetFedWatchTargetRate(ctx context.Context, c2s *qotgetfedwatchtargetrate.C2S) (*qotgetfedwatchtargetrate.S2C, error) {
	req := &qotgetfedwatchtargetrate.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetfedwatchtargetrate.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFedWatchTargetRate, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetFedWatchDotPlot 3405 - 获取CME利率点阵图
func (client *Client) QotGetFedWatchDotPlot(ctx context.Context, c2s *qotgetfedwatchdotplot.C2S) (*qotgetfedwatchdotplot.S2C, error) {
	req := &qotgetfedwatchdotplot.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetfedwatchdotplot.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetFedWatchDotPlot, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEarningsBeatRank 3406 - 获取盈利超预期排行
func (client *Client) QotGetEarningsBeatRank(ctx context.Context, c2s *qotgetearningsbeatrank.C2S) (*qotgetearningsbeatrank.S2C, error) {
	req := &qotgetearningsbeatrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetearningsbeatrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEarningsBeatRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetDividendRank 3407 - 获取股息排行
func (client *Client) QotGetDividendRank(ctx context.Context, c2s *qotgetdividendrank.C2S) (*qotgetdividendrank.S2C, error) {
	req := &qotgetdividendrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetdividendrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetDividendRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetDividendCalendar 3408 - 获取派息日历
func (client *Client) QotGetDividendCalendar(ctx context.Context, c2s *qotgetdividendcalendar.C2S) (*qotgetdividendcalendar.S2C, error) {
	req := &qotgetdividendcalendar.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetdividendcalendar.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetDividendCalendar, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEconomicCalendar 3409 - 获取财经日历(经济数据事件)
func (client *Client) QotGetEconomicCalendar(ctx context.Context, c2s *qotgeteconomiccalendar.C2S) (*qotgeteconomiccalendar.S2C, error) {
	req := &qotgeteconomiccalendar.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteconomiccalendar.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEconomicCalendar, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetUSPreMarketRank 3410 - 获取盘前榜(美股)
func (client *Client) QotGetUSPreMarketRank(ctx context.Context, c2s *qotgetuspremarketrank.C2S) (*qotgetuspremarketrank.S2C, error) {
	req := &qotgetuspremarketrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetuspremarketrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetUSPreMarketRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetUSAfterHoursRank 3411 - 获取盘后榜(美股)
func (client *Client) QotGetUSAfterHoursRank(ctx context.Context, c2s *qotgetusafterhoursrank.C2S) (*qotgetusafterhoursrank.S2C, error) {
	req := &qotgetusafterhoursrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetusafterhoursrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetUSAfterHoursRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetUSOvernightRank 3412 - 获取夜盘榜(美股)
func (client *Client) QotGetUSOvernightRank(ctx context.Context, c2s *qotgetusovernightrank.C2S) (*qotgetusovernightrank.S2C, error) {
	req := &qotgetusovernightrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetusovernightrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetUSOvernightRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetTopMoversRank 3413 - 获取领涨/领跌榜(盘中)
func (client *Client) QotGetTopMoversRank(ctx context.Context, c2s *qotgettopmoverrank.C2S) (*qotgettopmoverrank.S2C, error) {
	req := &qotgettopmoverrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgettopmoverrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetTopMoversRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetHotList 3414 - 获取热议榜
func (client *Client) QotGetHotList(ctx context.Context, c2s *qotgethotlist.C2S) (*qotgethotlist.S2C, error) {
	req := &qotgethotlist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgethotlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetHotList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetShortSellingRank 3415 - 获取卖空异动榜(美股)
func (client *Client) QotGetShortSellingRank(ctx context.Context, c2s *qotgetshortsellingrank.C2S) (*qotgetshortsellingrank.S2C, error) {
	req := &qotgetshortsellingrank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetshortsellingrank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetShortSellingRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetPeriodChangeRank 3416 - 获取区间涨跌幅
func (client *Client) QotGetPeriodChangeRank(ctx context.Context, c2s *qotgetperiodchangerank.C2S) (*qotgetperiodchangerank.S2C, error) {
	req := &qotgetperiodchangerank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetperiodchangerank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetPeriodChangeRank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetHighDividendSOERank 3417 - 获取破净高股息国央企(港股)
func (client *Client) QotGetHighDividendSOERank(ctx context.Context, c2s *qotgethighdividendsoerank.C2S) (*qotgethighdividendsoerank.S2C, error) {
	req := &qotgethighdividendsoerank.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgethighdividendsoerank.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetHighDividendSOERank, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInstitutionList 3418 - 获取机构持仓列表
func (client *Client) QotGetInstitutionList(ctx context.Context, c2s *qotgetinstitutionlist.C2S) (*qotgetinstitutionlist.S2C, error) {
	req := &qotgetinstitutionlist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetinstitutionlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInstitutionList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInstitutionProfile 3419 - 获取机构概况
func (client *Client) QotGetInstitutionProfile(ctx context.Context, c2s *qotgetinstitutionprofile.C2S) (*qotgetinstitutionprofile.S2C, error) {
	req := &qotgetinstitutionprofile.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetinstitutionprofile.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInstitutionProfile, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInstitutionDistribution 3420 - 获取机构持仓行业分布
func (client *Client) QotGetInstitutionDistribution(ctx context.Context, c2s *qotgetinstitutiondistribution.C2S) (*qotgetinstitutiondistribution.S2C, error) {
	req := &qotgetinstitutiondistribution.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetinstitutiondistribution.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInstitutionDistribution, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInstitutionHoldingChange 3421 - 获取机构持仓变动
func (client *Client) QotGetInstitutionHoldingChange(ctx context.Context, c2s *qotgetinstitutionholdingchange.C2S) (*qotgetinstitutionholdingchange.S2C, error) {
	req := &qotgetinstitutionholdingchange.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetinstitutionholdingchange.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInstitutionHoldingChange, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetInstitutionHoldingList 3422 - 获取机构持股列表
func (client *Client) QotGetInstitutionHoldingList(ctx context.Context, c2s *qotgetinstitutionholdinglist.C2S) (*qotgetinstitutionholdinglist.S2C, error) {
	req := &qotgetinstitutionholdinglist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetinstitutionholdinglist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetInstitutionHoldingList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetArkFundHolding 3423 - 获取ARK基金持仓
func (client *Client) QotGetArkFundHolding(ctx context.Context, c2s *qotgetarkfundholding.C2S) (*qotgetarkfundholding.S2C, error) {
	req := &qotgetarkfundholding.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetarkfundholding.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetArkFundHolding, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetArkStockDynamic 3424 - 获取ARK个股交易动态
func (client *Client) QotGetArkStockDynamic(ctx context.Context, c2s *qotgetarkstockdynamic.C2S) (*qotgetarkstockdynamic.S2C, error) {
	req := &qotgetarkstockdynamic.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetarkstockdynamic.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetArkStockDynamic, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetArkActiveTransaction 3425 - 获取ARK主动交易聚合
func (client *Client) QotGetArkActiveTransaction(ctx context.Context, c2s *qotgetarkactivetransaction.C2S) (*qotgetarkactivetransaction.S2C, error) {
	req := &qotgetarkactivetransaction.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetarkactivetransaction.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetArkActiveTransaction, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetRatingChange 3426 - 获取评级变动
func (client *Client) QotGetRatingChange(ctx context.Context, c2s *qotgetratingchange.C2S) (*qotgetratingchange.S2C, error) {
	req := &qotgetratingchange.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetratingchange.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetRatingChange, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIndustrialChainList 3427 - 获取产业链列表
func (client *Client) QotGetIndustrialChainList(ctx context.Context, c2s *qotgetindustrialchainlist.C2S) (*qotgetindustrialchainlist.S2C, error) {
	req := &qotgetindustrialchainlist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetindustrialchainlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIndustrialChainList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIndustrialChainDetail 3428 - 获取产业链详情
func (client *Client) QotGetIndustrialChainDetail(ctx context.Context, c2s *qotgetindustrialchaindetail.C2S) (*qotgetindustrialchaindetail.S2C, error) {
	req := &qotgetindustrialchaindetail.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetindustrialchaindetail.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIndustrialChainDetail, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIndustrialChainByPlate 3429 - 获取板块关联产业链
func (client *Client) QotGetIndustrialChainByPlate(ctx context.Context, c2s *qotgetindustrialchainbyplate.C2S) (*qotgetindustrialchainbyplate.S2C, error) {
	req := &qotgetindustrialchainbyplate.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetindustrialchainbyplate.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIndustrialChainByPlate, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIndustrialPlateInfo 3430 - 获取产业板块信息
func (client *Client) QotGetIndustrialPlateInfo(ctx context.Context, c2s *qotgetindustrialplateinfo.C2S) (*qotgetindustrialplateinfo.S2C, error) {
	req := &qotgetindustrialplateinfo.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetindustrialplateinfo.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIndustrialPlateInfo, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetIndustrialPlateStock 3431 - 获取产业板块成分股
func (client *Client) QotGetIndustrialPlateStock(ctx context.Context, c2s *qotgetindustrialplatestock.C2S) (*qotgetindustrialplatestock.S2C, error) {
	req := &qotgetindustrialplatestock.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetindustrialplatestock.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetIndustrialPlateStock, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetHeatMapData 3432 - 获取热力图数据
func (client *Client) QotGetHeatMapData(ctx context.Context, c2s *qotgetheatmapdata.C2S) (*qotgetheatmapdata.S2C, error) {
	req := &qotgetheatmapdata.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetheatmapdata.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetHeatMapData, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetRiseFallDistribution 3433 - 获取涨跌分布
func (client *Client) QotGetRiseFallDistribution(ctx context.Context, c2s *qotgetrisefalldistr.C2S) (*qotgetrisefalldistr.S2C, error) {
	req := &qotgetrisefalldistr.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgetrisefalldistr.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetRiseFallDistribution, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractCategory 3434 - 获取事件合约分类
func (client *Client) QotGetEventContractCategory(ctx context.Context, c2s *qotgeteventcontractcategory.C2S) (*qotgeteventcontractcategory.S2C, error) {
	req := &qotgeteventcontractcategory.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractcategory.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractCategory, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotFilterCompetition 3435 - 获取赛事筛选选项
func (client *Client) QotFilterCompetition(ctx context.Context, c2s *qotfiltercompetition.C2S) (*qotfiltercompetition.S2C, error) {
	req := &qotfiltercompetition.Request{
		C2S: c2s,
	}

	ch := make(chan *qotfiltercompetition.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotFilterCompetition, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractSeriesList 3436 - 获取事件合约Series列表
func (client *Client) QotGetEventContractSeriesList(ctx context.Context, c2s *qotgeteventcontractserieslist.C2S) (*qotgeteventcontractserieslist.S2C, error) {
	req := &qotgeteventcontractserieslist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractserieslist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractSeriesList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractEventList 3437 - 获取事件合约Event列表
func (client *Client) QotGetEventContractEventList(ctx context.Context, c2s *qotgeteventcontracteventlist.C2S) (*qotgeteventcontracteventlist.S2C, error) {
	req := &qotgeteventcontracteventlist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontracteventlist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractEventList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContract 3438 - 获取事件合约(合约列表)
func (client *Client) QotGetEventContract(ctx context.Context, c2s *qotgeteventcontract.C2S) (*qotgeteventcontract.S2C, error) {
	req := &qotgeteventcontract.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontract.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContract, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractMilestoneList 3439 - 获取事件合约里程碑列表
func (client *Client) QotGetEventContractMilestoneList(ctx context.Context, c2s *qotgeteventcontractmilestonelist.C2S) (*qotgeteventcontractmilestonelist.S2C, error) {
	req := &qotgeteventcontractmilestonelist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractmilestonelist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractMilestoneList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractSnapshot 3445 - 获取事件合约快照
func (client *Client) QotGetEventContractSnapshot(ctx context.Context, c2s *qotgeteventcontractsnapshot.C2S) (*qotgeteventcontractsnapshot.S2C, error) {
	req := &qotgeteventcontractsnapshot.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractsnapshot.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractSnapshot, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractOrderBook 3446 - 获取事件合约摆盘
func (client *Client) QotGetEventContractOrderBook(ctx context.Context, c2s *qotgeteventcontractorderbook.C2S) (*qotgeteventcontractorderbook.S2C, error) {
	req := &qotgeteventcontractorderbook.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractorderbook.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractOrderBook, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractKline 3447 - 获取事件合约K线
func (client *Client) QotGetEventContractKline(ctx context.Context, c2s *qotgeteventcontractkline.C2S) (*qotgeteventcontractkline.S2C, error) {
	req := &qotgeteventcontractkline.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractkline.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractKline, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractTicker 3448 - 获取事件合约逐笔
func (client *Client) QotGetEventContractTicker(ctx context.Context, c2s *qotgeteventcontractticker.C2S) (*qotgeteventcontractticker.S2C, error) {
	req := &qotgeteventcontractticker.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractticker.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractTicker, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractComboList 3453 - 获取事件合约可Combo列表
func (client *Client) QotGetEventContractComboList(ctx context.Context, c2s *qotgeteventcontractcombolist.C2S) (*qotgeteventcontractcombolist.S2C, error) {
	req := &qotgeteventcontractcombolist.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractcombolist.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractComboList, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotGetEventContractComboRfq 3454 - 事件合约Combo询价
func (client *Client) QotGetEventContractComboRfq(ctx context.Context, c2s *qotgeteventcontractcomborfq.C2S) (*qotgeteventcontractcomborfq.S2C, error) {
	req := &qotgeteventcontractcomborfq.Request{
		C2S: c2s,
	}

	ch := make(chan *qotgeteventcontractcomborfq.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotGetEventContractComboRfq, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotSubEventContract 3455 - 事件合约订阅/反订阅，该接口的S2C返回的是空
func (client *Client) QotSubEventContract(ctx context.Context, c2s *qotsubeventcontract.C2S) error {
	req := &qotsubeventcontract.Request{
		C2S: c2s,
	}

	ch := make(chan *qotsubeventcontract.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotSubEventContract, req, infra.NewProtobufChan(ch)); err != nil {
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

// QotRequestHistoryEventContractKL 3456 - 拉取事件合约历史K线
func (client *Client) QotRequestHistoryEventContractKL(ctx context.Context, c2s *qotrequesthistoryeventcontractkl.C2S) (*qotrequesthistoryeventcontractkl.S2C, error) {
	req := &qotrequesthistoryeventcontractkl.Request{
		C2S: c2s,
	}

	ch := make(chan *qotrequesthistoryeventcontractkl.Response, 1)
	defer close(ch)
	if err := client.Request(protoid.QotRequestHistoryEventContractKL, req, infra.NewProtobufChan(ch)); err != nil {
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
