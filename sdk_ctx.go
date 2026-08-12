package futu

import (
	"context"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
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
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdflowsummary"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetcombomaxtrdqtys"
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
	"github.com/hyperjiang/futu/pb/trdplacecomboorder"
	"github.com/hyperjiang/futu/pb/trdplaceorder"
	"github.com/hyperjiang/futu/pb/trdsubaccpush"
	"github.com/hyperjiang/futu/pb/trdunlocktrade"
	"google.golang.org/protobuf/proto"
)

// GetGlobalStateWithContext 1002 - gets the global state with context.
func (sdk *SDK) GetGlobalStateWithContext(ctx context.Context) (*getglobalstate.S2C, error) {
	return sdk.cli.GetGlobalState(ctx)
}

// GetAccListWithContext 2001 - gets the account list with context.
func (sdk *SDK) GetAccListWithContext(ctx context.Context, opts ...adapt.Option) ([]*trdcommon.TrdAcc, error) {
	o := adapt.NewOptions(opts...)

	var c2s trdgetacclist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetAccList(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetAccList(), nil
}

// UnlockTrade 2005 - unlocks or locks the trade.
//
// unlock: true for unlock, false for lock
//
// pwdMD5: MD5 of the password
//
// securityFirm: security firm
func (sdk *SDK) UnlockTradeWithContext(ctx context.Context, unlock bool, pwdMD5 string, securityFirm int32) error {
	c2s := &trdunlocktrade.C2S{
		Unlock:       proto.Bool(unlock),
		PwdMD5:       proto.String(pwdMD5),
		SecurityFirm: proto.Int32(securityFirm),
	}

	return sdk.cli.TrdUnlockTrade(ctx, c2s)
}

// SubscribeAccPushWithContext 2008 - subscribes the trading account push data.
//
// accIDList: account ID list
func (sdk *SDK) SubscribeAccPushWithContext(ctx context.Context, accIDList []uint64) error {
	c2s := &trdsubaccpush.C2S{
		AccIDList: accIDList,
	}

	return sdk.cli.TrdSubAccPush(ctx, c2s)
}

// GetFundsWithContext 2101 - gets the funds with context.
func (sdk *SDK) GetFundsWithContext(ctx context.Context, header *trdcommon.TrdHeader, opts ...adapt.Option) (*trdcommon.Funds, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header

	var c2s trdgetfunds.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetFunds(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetFunds(), nil
}

// GetPositionListWithContext 2102 - gets the position list with context.
func (sdk *SDK) GetPositionListWithContext(ctx context.Context, header *trdcommon.TrdHeader, opts ...adapt.Option) ([]*trdcommon.Position, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header

	var c2s trdgetpositionlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetPositionList(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetPositionList(), nil
}

// GetMaxTrdQtysWithContext 2111 - gets the maximum available trading quantities with context.
//
// header: trading header
//
// orderType: order type
//
// code: security code, e.g. AAPL
//
// price: price
func (sdk *SDK) GetMaxTrdQtysWithContext(ctx context.Context, header *trdcommon.TrdHeader, orderType int32, code string, price float64, opts ...adapt.Option) (*trdcommon.MaxTrdQtys, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["orderType"] = orderType
	o["price"] = price
	o.SetCodeForTrade(code)

	var c2s trdgetmaxtrdqtys.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetMaxTrdQtys(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetMaxTrdQtys(), nil
}

// GetOpenOrderListWithContext 2201 - gets the open order list with context.
func (sdk *SDK) GetOpenOrderListWithContext(ctx context.Context, header *trdcommon.TrdHeader, opts ...adapt.Option) ([]*trdcommon.Order, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header

	var c2s trdgetorderlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetOrderList(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOrderList(), nil
}

// PlaceOrderWithContext 2202 - places an order with context.
//
// header: trading header
//
// trdSide: trading side
//
// orderType: order type
//
// code: security code, e.g. US.AAPL
//
// qty: quantity
//
// price: price
func (sdk *SDK) PlaceOrderWithContext(ctx context.Context, header *trdcommon.TrdHeader, trdSide int32, orderType int32, code string, qty float64, price float64, opts ...adapt.Option) (*trdplaceorder.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["trdSide"] = trdSide
	o["orderType"] = orderType
	o["qty"] = qty
	o["price"] = price
	o.SetCodeForTrade(code)

	var c2s trdplaceorder.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.TrdPlaceOrder(ctx, &c2s)
}

// ModifyOrderWithContext 2205 - modifies an order with context.
//
// header: trading header
//
// orderID: order ID, use 0 if forAll=true
//
// modifyOrderOp: modify order operation
func (sdk *SDK) ModifyOrderWithContext(ctx context.Context, header *trdcommon.TrdHeader, orderID uint64, modifyOrderOp int32, opts ...adapt.Option) (*trdmodifyorder.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["orderID"] = orderID
	o["modifyOrderOp"] = modifyOrderOp

	var c2s trdmodifyorder.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.TrdModifyOrder(ctx, &c2s)
}

// GetOrderFillListWithContext 2211 - gets the filled order list with context.
func (sdk *SDK) GetOrderFillListWithContext(ctx context.Context, header *trdcommon.TrdHeader, opts ...adapt.Option) ([]*trdcommon.OrderFill, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header

	var c2s trdgetorderfilllist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetOrderFillList(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOrderFillList(), nil
}

// GetHistoryOrderListWithContext 2221 - gets the history order list with context.
func (sdk *SDK) GetHistoryOrderListWithContext(ctx context.Context, header *trdcommon.TrdHeader, fc *trdcommon.TrdFilterConditions, opts ...adapt.Option) ([]*trdcommon.Order, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["filterConditions"] = fc

	var c2s trdgethistoryorderlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetHistoryOrderList(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOrderList(), nil
}

// GetHistoryOrderFillListWithContext 2222 - gets the history filled order list with context.
func (sdk *SDK) GetHistoryOrderFillListWithContext(ctx context.Context, header *trdcommon.TrdHeader, fc *trdcommon.TrdFilterConditions, opts ...adapt.Option) ([]*trdcommon.OrderFill, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["filterConditions"] = fc

	var c2s trdgethistoryorderfilllist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.TrdGetHistoryOrderFillList(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOrderFillList(), nil
}

// GetMarginRatioWithContext 2223 - gets the margin ratio with context.
func (sdk *SDK) GetMarginRatioWithContext(ctx context.Context, header *trdcommon.TrdHeader, codes []string) ([]*trdgetmarginratio.MarginRatioInfo, error) {
	c2s := &trdgetmarginratio.C2S{
		Header:       header,
		SecurityList: adapt.NewSecurities(codes),
	}

	s2c, err := sdk.cli.TrdGetMarginRatio(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetMarginRatioInfoList(), nil
}

// GetOrderFeeWithContext 2225 - gets the order fee with context.
func (sdk *SDK) GetOrderFeeWithContext(ctx context.Context, header *trdcommon.TrdHeader, orderIdExList []string) ([]*trdcommon.OrderFee, error) {
	c2s := &trdgetorderfee.C2S{
		Header:        header,
		OrderIdExList: orderIdExList,
	}

	s2c, err := sdk.cli.TrdGetOrderFee(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOrderFeeList(), nil
}

// TrdFlowSummaryWithContext 2226 - gets the trading flow summary with context.
func (sdk *SDK) TrdFlowSummaryWithContext(ctx context.Context, header *trdcommon.TrdHeader, clearingDate string) ([]*trdflowsummary.FlowSummaryInfo, error) {
	c2s := &trdflowsummary.C2S{
		Header:       header,
		ClearingDate: proto.String(clearingDate),
	}

	s2c, err := sdk.cli.TrdFlowSummary(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetFlowSummaryInfoList(), nil
}

// SubscribeWithContext 3001 - subscribes or unsubscribes with context.
//
// codes: security codes
//
// subTypes: subscription types
//
// isSub: true for subscribe, false for unsubscribe
func (sdk *SDK) SubscribeWithContext(ctx context.Context, codes []string, subTypes []int32, isSub bool, opts ...adapt.Option) error {
	o := adapt.NewOptions(opts...)
	o["securityList"] = adapt.NewSecurities(codes)
	o["subTypeList"] = subTypes
	o["isSubOrUnSub"] = isSub

	var c2s qotsub.C2S
	if err := o.ToProto(&c2s); err != nil {
		return err
	}

	return sdk.cli.QotSub(ctx, &c2s)
}

// GetSubInfoWithContext 3003 - gets the subscription information with context.
func (sdk *SDK) GetSubInfoWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetsubinfo.S2C, error) {
	o := adapt.NewOptions(opts...)
	var c2s qotgetsubinfo.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetSubInfo(ctx, &c2s)
}

// GetBasicQotWithContext 3004 - gets the basic quotes of given securities with context.
func (sdk *SDK) GetBasicQotWithContext(ctx context.Context, codes []string) ([]*qotcommon.BasicQot, error) {
	c2s := &qotgetbasicqot.C2S{
		SecurityList: adapt.NewSecurities(codes),
	}

	s2c, err := sdk.cli.QotGetBasicQot(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetBasicQotList(), nil
}

// GetKLWithContext 3006 - gets K-line data with context.
func (sdk *SDK) GetKLWithContext(ctx context.Context, code string, klType int32, opts ...adapt.Option) (*qotgetkl.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["klType"] = klType

	if _, ok := o["rehabType"]; !ok {
		o["rehabType"] = adapt.RehabType_None
	}

	if _, ok := o["reqNum"]; !ok {
		o["reqNum"] = 1000
	}

	var c2s qotgetkl.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetKL(ctx, &c2s)
}

// GetRTWithContext 3008 - gets real-time data with context.
//
// code: security code
func (sdk *SDK) GetRTWithContext(ctx context.Context, code string) (*qotgetrt.S2C, error) {
	c2s := &qotgetrt.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetRT(ctx, c2s)
}

// GetTickerWithContext 3010 - gets the ticker data with context.
//
// code: security code
func (sdk *SDK) GetTickerWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetticker.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	if _, ok := o["maxRetNum"]; !ok {
		o["maxRetNum"] = 1000
	}

	var c2s qotgetticker.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetTicker(ctx, &c2s)
}

// GetOrderBookWithContext 3012 - gets the order book with context.
//
// code: security code
func (sdk *SDK) GetOrderBookWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetorderbook.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	if _, ok := o["num"]; !ok {
		o["num"] = 100
	}

	var c2s qotgetorderbook.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOrderBook(ctx, &c2s)
}

// GetBrokerWithContext 3014 - gets the broker with context.
//
// code: security code
func (sdk *SDK) GetBrokerWithContext(ctx context.Context, code string) (*qotgetbroker.S2C, error) {
	c2s := &qotgetbroker.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetBroker(ctx, c2s)
}

// RequestHistoryKLWithContext 3103 - requests the history K-line data with context.
//
// code: security code
//
// klType: K-line type
//
// beginTime: begin time, format: "yyyy-MM-dd"
//
// endTime: end time, format: "yyyy-MM-dd"
func (sdk *SDK) RequestHistoryKLWithContext(ctx context.Context, code string, klType int32, beginTime string, endTime string, opts ...adapt.Option) (*qotrequesthistorykl.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["klType"] = klType
	o["beginTime"] = beginTime
	o["endTime"] = endTime

	if _, ok := o["rehabType"]; !ok {
		o["rehabType"] = adapt.RehabType_None
	}

	var c2s qotrequesthistorykl.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotRequestHistoryKL(ctx, &c2s)
}

// RequestHistoryKLQuotaWithContext 3104 - requests the history K-line quota with context.
func (sdk *SDK) RequestHistoryKLQuotaWithContext(ctx context.Context, opts ...adapt.Option) (*qotrequesthistoryklquota.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotrequesthistoryklquota.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotRequestHistoryKLQuota(ctx, &c2s)
}

// RequestRehabWithContext 3105 - requests the rehab data with context.
//
// code: security code
func (sdk *SDK) RequestRehabWithContext(ctx context.Context, code string) (*qotrequestrehab.S2C, error) {
	c2s := &qotrequestrehab.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotRequestRehab(ctx, c2s)
}

// GetStaticInfoWithContext 3202 - gets the static information with context.
func (sdk *SDK) GetStaticInfoWithContext(ctx context.Context, opts ...adapt.Option) ([]*qotcommon.SecurityStaticInfo, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetstaticinfo.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	res, err := sdk.cli.QotGetStaticInfo(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return res.GetStaticInfoList(), nil
}

// GetSecuritySnapshotWithContext 3203 - gets the security snapshot with context.
//
// codes: security codes
func (sdk *SDK) GetSecuritySnapshotWithContext(ctx context.Context, codes []string) ([]*qotgetsecuritysnapshot.Snapshot, error) {
	c2s := &qotgetsecuritysnapshot.C2S{
		SecurityList: adapt.NewSecurities(codes),
	}

	s2c, err := sdk.cli.QotGetSecuritySnapshot(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetSnapshotList(), nil
}

// GetPlateSetWithContext 3204 - gets the plate set with context.
//
// market: market
//
// plateSetType: plate set type
func (sdk *SDK) GetPlateSetWithContext(ctx context.Context, market int32, plateSetType int32) ([]*qotcommon.PlateInfo, error) {
	c2s := &qotgetplateset.C2S{
		Market:       proto.Int32(market),
		PlateSetType: proto.Int32(plateSetType),
	}

	s2c, err := sdk.cli.QotGetPlateSet(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetPlateInfoList(), nil
}

// GetPlateSecurityWithContext 3205 - gets the plate securities with context.
//
// plateCode: plate code
func (sdk *SDK) GetPlateSecurityWithContext(ctx context.Context, plateCode string, opts ...adapt.Option) ([]*qotcommon.SecurityStaticInfo, error) {
	o := adapt.NewOptions(opts...)
	o["plate"] = adapt.NewSecurity(plateCode)

	var c2s qotgetplatesecurity.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.QotGetPlateSecurity(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetStaticInfoList(), nil
}

// GetReferenceWithContext 3206 - gets the reference with context.
//
// code: security code
//
// refType: reference type
func (sdk *SDK) GetReferenceWithContext(ctx context.Context, code string, refType int32) ([]*qotcommon.SecurityStaticInfo, error) {
	c2s := &qotgetreference.C2S{
		Security:      adapt.NewSecurity(code),
		ReferenceType: proto.Int32(refType),
	}

	s2c, err := sdk.cli.QotGetReference(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetStaticInfoList(), nil
}

// GetOwnerPlateWithContext 3207 - gets the owner plate with context.
//
// codes: security codes
func (sdk *SDK) GetOwnerPlateWithContext(ctx context.Context, codes []string) ([]*qotgetownerplate.SecurityOwnerPlate, error) {
	c2s := &qotgetownerplate.C2S{
		SecurityList: adapt.NewSecurities(codes),
	}

	s2c, err := sdk.cli.QotGetOwnerPlate(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOwnerPlateList(), nil
}

// GetOptionChainWithContext 3209 - gets the option chain with context.
//
// code: security code
//
// beginTime: begin time, format: "yyyy-MM-dd"
//
// endTime: end time, format: "yyyy-MM-dd"
func (sdk *SDK) GetOptionChainWithContext(ctx context.Context, code string, beginTime string, endTime string, opts ...adapt.Option) ([]*qotgetoptionchain.OptionChain, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)
	o["beginTime"] = beginTime
	o["endTime"] = endTime

	var c2s qotgetoptionchain.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.QotGetOptionChain(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetOptionChain(), nil
}

// GetWarrantWithContext 3210 - gets the warrant with context, only available in Hong Kong market.
// Sort by score in descending order by default.
//
// begin: begin index
//
// num: number of warrants
func (sdk *SDK) GetWarrantWithContext(ctx context.Context, begin int32, num int32, opts ...adapt.Option) (*qotgetwarrant.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["begin"] = begin
	o["num"] = num

	if _, ok := o["sortField"]; !ok {
		o["sortField"] = adapt.SortField_Score
	}

	if _, ok := o["ascend"]; !ok {
		o["ascend"] = false
	}

	var c2s qotgetwarrant.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetWarrant(ctx, &c2s)
}

// GetCapitalFlowWithContext 3211 - gets the capital flow with context.
//
// code: security code
func (sdk *SDK) GetCapitalFlowWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetcapitalflow.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	if _, ok := o["periodType"]; !ok {
		o["periodType"] = adapt.PeriodType_INTRADAY
	}

	var c2s qotgetcapitalflow.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetCapitalFlow(ctx, &c2s)
}

// GetCapitalDistributionWithContext 3212 - gets the capital distribution with context.
//
// code: security code
func (sdk *SDK) GetCapitalDistributionWithContext(ctx context.Context, code string) (*qotgetcapitaldistribution.S2C, error) {
	c2s := &qotgetcapitaldistribution.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetCapitalDistribution(ctx, c2s)
}

// GetUserSecurityWithContext 3213 - gets the user security with context.
//
// groupName: group name
func (sdk *SDK) GetUserSecurityWithContext(ctx context.Context, groupName string) ([]*qotcommon.SecurityStaticInfo, error) {
	c2s := &qotgetusersecurity.C2S{
		GroupName: proto.String(groupName),
	}

	s2c, err := sdk.cli.QotGetUserSecurity(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetStaticInfoList(), nil
}

// ModifyUserSecurityWithContext 3214 - modifies the user security with context.
//
// groupName: group name
//
// codes: security codes
//
// op: operation
func (sdk *SDK) ModifyUserSecurityWithContext(ctx context.Context, groupName string, codes []string, op int32) error {
	c2s := &qotmodifyusersecurity.C2S{
		GroupName:    proto.String(groupName),
		SecurityList: adapt.NewSecurities(codes),
		Op:           proto.Int32(op),
	}

	return sdk.cli.QotModifyUserSecurity(ctx, c2s)
}

// StockFilterWithContext 3215 - filters the stock with context.
//
// market: market
func (sdk *SDK) StockFilterWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotstockfilter.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	if _, ok := o["begin"]; !ok {
		o["begin"] = 0
	}

	if _, ok := o["num"]; !ok {
		o["num"] = 200
	}

	var c2s qotstockfilter.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotStockFilter(ctx, &c2s)
}

// GetIpoListWithContext 3217 - gets the IPO list with context.
//
// market: market
func (sdk *SDK) GetIpoListWithContext(ctx context.Context, market int32) ([]*qotgetipolist.IpoData, error) {
	c2s := &qotgetipolist.C2S{
		Market: proto.Int32(market),
	}

	s2c, err := sdk.cli.QotGetIpoList(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetIpoList(), nil
}

// GetFutureInfoWithContext 3218 - gets the future information with context.
//
// codes: security codes
func (sdk *SDK) GetFutureInfoWithContext(ctx context.Context, codes []string) ([]*qotgetfutureinfo.FutureInfo, error) {
	c2s := &qotgetfutureinfo.C2S{
		SecurityList: adapt.NewSecurities(codes),
	}

	s2c, err := sdk.cli.QotGetFutureInfo(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetFutureInfoList(), nil
}

// RequestTradeDateWithContext 3219 - requests the trade date with context.
//
// market: market
//
// code: security code
//
// beginTime: begin time, format: "yyyy-MM-dd"
//
// endTime: end time, format: "yyyy-MM-dd"
func (sdk *SDK) RequestTradeDateWithContext(ctx context.Context, market int32, code string, beginTime string, endTime string) ([]*qotrequesttradedate.TradeDate, error) {
	c2s := &qotrequesttradedate.C2S{
		Market:    proto.Int32(market),
		BeginTime: proto.String(beginTime),
		EndTime:   proto.String(endTime),
	}
	if code != "" {
		c2s.Security = adapt.NewSecurity(code)
	}

	s2c, err := sdk.cli.QotRequestTradeDate(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetTradeDateList(), nil
}

// SetPriceReminderWithContext 3220 - sets the price reminder with context.
//
// code: security code
//
// op: operation
func (sdk *SDK) SetPriceReminderWithContext(ctx context.Context, code string, op int32, opts ...adapt.Option) (int64, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["op"] = op

	var c2s qotsetpricereminder.C2S
	if err := o.ToProto(&c2s); err != nil {
		return 0, err
	}

	s2c, err := sdk.cli.QotSetPriceReminder(ctx, &c2s)
	if err != nil {
		return 0, err
	}

	return s2c.GetKey(), nil
}

// GetPriceReminderWithContext 3221 - gets the price reminder with context.
//
// code: security code
//
// market: market, if security is set, this param is ignored
func (sdk *SDK) GetPriceReminderWithContext(ctx context.Context, code string, market int32) ([]*qotgetpricereminder.PriceReminder, error) {
	c2s := &qotgetpricereminder.C2S{
		Security: adapt.NewSecurity(code),
		Market:   proto.Int32(market),
	}

	s2c, err := sdk.cli.QotGetPriceReminder(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetPriceReminderList(), nil
}

// GetUserSecurityGroupWithContext 3222 - gets the user security group with context.
//
// groupType: group type
func (sdk *SDK) GetUserSecurityGroupWithContext(ctx context.Context, groupType int32) ([]*qotgetusersecuritygroup.GroupData, error) {
	c2s := &qotgetusersecuritygroup.C2S{
		GroupType: proto.Int32(groupType),
	}

	s2c, err := sdk.cli.QotGetUserSecurityGroup(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetGroupList(), nil
}

// GetMarketStateWithContext 3223 - gets the market state with context.
//
// codes: security codes
func (sdk *SDK) GetMarketStateWithContext(ctx context.Context, codes []string) ([]*qotgetmarketstate.MarketInfo, error) {
	c2s := &qotgetmarketstate.C2S{
		SecurityList: adapt.NewSecurities(codes),
	}

	s2c, err := sdk.cli.QotGetMarketState(ctx, c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetMarketInfoList(), nil
}

// GetOptionExpirationDateWithContext 3224 - gets the option expiration date with context.
//
// code: security code
func (sdk *SDK) GetOptionExpirationDateWithContext(ctx context.Context, code string, opts ...adapt.Option) ([]*qotgetoptionexpirationdate.OptionExpirationDate, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)

	var c2s qotgetoptionexpirationdate.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	s2c, err := sdk.cli.QotGetOptionExpirationDate(ctx, &c2s)
	if err != nil {
		return nil, err
	}

	return s2c.GetDateList(), nil
}

// GetFinancialsEarningsPriceMoveWithContext 3225 - gets the financials earnings price move with context.
//
// code: security code
func (sdk *SDK) GetFinancialsEarningsPriceMoveWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetfinancialsearnpricemove.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetfinancialsearnpricemove.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetFinancialsEarningsPriceMove(ctx, &c2s)
}

// GetFinancialsEarningsPriceHistoryWithContext 3226 - gets the financials earnings price history with context.
//
// code: security code
func (sdk *SDK) GetFinancialsEarningsPriceHistoryWithContext(ctx context.Context, code string) (*qotgetfinancialsearnpricehist.S2C, error) {
	c2s := &qotgetfinancialsearnpricehist.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetFinancialsEarningsPriceHistory(ctx, c2s)
}

// GetFinancialsStatementsWithContext 3227 - gets the financial statements with context.
//
// code: security code
//
// statementType: financial statement type, see adapt.FinancialStatementsType_*
func (sdk *SDK) GetFinancialsStatementsWithContext(ctx context.Context, code string, statementType int32, opts ...adapt.Option) (*qotgetfinancialsstatements.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["statementType"] = statementType

	var c2s qotgetfinancialsstatements.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetFinancialsStatements(ctx, &c2s)
}

// GetFinancialsRevenueBreakdownWithContext 3228 - gets the financials revenue breakdown with context.
//
// code: security code
func (sdk *SDK) GetFinancialsRevenueBreakdownWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetfinancialrevenuebreakdown.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetfinancialrevenuebreakdown.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetFinancialsRevenueBreakdown(ctx, &c2s)
}

// GetResearchAnalystConsensusWithContext 3229 - gets the research analyst consensus with context.
//
// code: security code
func (sdk *SDK) GetResearchAnalystConsensusWithContext(ctx context.Context, code string) (*qotgetresearchanalystconsensus.S2C, error) {
	c2s := &qotgetresearchanalystconsensus.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetResearchAnalystConsensus(ctx, c2s)
}

// GetResearchRatingSummaryWithContext 3230 - gets the research rating summary with context.
//
// code: security code
func (sdk *SDK) GetResearchRatingSummaryWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetresearchratingsummary.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetresearchratingsummary.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetResearchRatingSummary(ctx, &c2s)
}

// GetResearchMorningstarReportWithContext 3231 - gets the research morningstar report with context.
//
// code: security code
func (sdk *SDK) GetResearchMorningstarReportWithContext(ctx context.Context, code string) (*qotgetresearchmorningstarrpt.S2C, error) {
	c2s := &qotgetresearchmorningstarrpt.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetResearchMorningstarReport(ctx, c2s)
}

// GetValuationDetailWithContext 3232 - gets the valuation detail with context.
//
// code: security code
//
// valuationType: valuation type, see adapt.ValuationType_*
func (sdk *SDK) GetValuationDetailWithContext(ctx context.Context, code string, valuationType int32, opts ...adapt.Option) (*qotgetvaluationdetail.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["valuationType"] = valuationType

	var c2s qotgetvaluationdetail.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetValuationDetail(ctx, &c2s)
}

// GetValuationPlateStockListWithContext 3233 - gets the valuation plate stock list with context.
//
// code: security code (plate code)
//
// valuationType: valuation type, see adapt.ValuationType_*
func (sdk *SDK) GetValuationPlateStockListWithContext(ctx context.Context, code string, valuationType int32, opts ...adapt.Option) (*qotgetvaluationplatestocklist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["valuationType"] = valuationType

	var c2s qotgetvaluationplatestocklist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetValuationPlateStockList(ctx, &c2s)
}

// GetCorporateActionsDividendsWithContext 3234 - gets the corporate actions dividends with context.
//
// code: security code
func (sdk *SDK) GetCorporateActionsDividendsWithContext(ctx context.Context, code string) (*qotgetcorporateactionsdividends.S2C, error) {
	c2s := &qotgetcorporateactionsdividends.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetCorporateActionsDividends(ctx, c2s)
}

// GetCorporateActionsBuybacksWithContext 3235 - gets the corporate actions buybacks with context.
//
// code: security code
func (sdk *SDK) GetCorporateActionsBuybacksWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetcorporateactionsbuybacks.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetcorporateactionsbuybacks.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetCorporateActionsBuybacks(ctx, &c2s)
}

// GetCorporateActionsStockSplitsWithContext 3236 - gets the corporate actions stock splits with context.
//
// code: security code
func (sdk *SDK) GetCorporateActionsStockSplitsWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetcorporateactionsstocksplits.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetcorporateactionsstocksplits.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetCorporateActionsStockSplits(ctx, &c2s)
}

// GetShareholdersOverviewWithContext 3237 - gets the shareholders overview with context.
//
// code: security code
func (sdk *SDK) GetShareholdersOverviewWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetshareholdersoverview.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetshareholdersoverview.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetShareholdersOverview(ctx, &c2s)
}

// GetShareholdersHoldingChangesWithContext 3238 - gets the shareholders holding changes with context.
//
// code: security code
func (sdk *SDK) GetShareholdersHoldingChangesWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetshareholdersholdingchanges.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetshareholdersholdingchanges.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetShareholdersHoldingChanges(ctx, &c2s)
}

// GetShareholdersHolderDetailWithContext 3239 - gets the shareholders holder detail with context.
//
// code: security code
func (sdk *SDK) GetShareholdersHolderDetailWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetshareholdersholderdetail.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetshareholdersholderdetail.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetShareholdersHolderDetail(ctx, &c2s)
}

// GetShareholdersInstitutionalWithContext 3240 - gets the shareholders institutional with context.
//
// code: security code
func (sdk *SDK) GetShareholdersInstitutionalWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetshareholdersinstitutional.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetshareholdersinstitutional.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetShareholdersInstitutional(ctx, &c2s)
}

// GetInsiderHolderListWithContext 3241 - gets the insider holder list with context.
//
// code: security code
func (sdk *SDK) GetInsiderHolderListWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetinsiderholderlist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetinsiderholderlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInsiderHolderList(ctx, &c2s)
}

// GetInsiderTradeListWithContext 3242 - gets the insider trade list with context.
//
// code: security code
func (sdk *SDK) GetInsiderTradeListWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetinsidertradelist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetinsidertradelist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInsiderTradeList(ctx, &c2s)
}

// GetCompanyProfileWithContext 3243 - gets the company profile with context.
//
// code: security code
func (sdk *SDK) GetCompanyProfileWithContext(ctx context.Context, code string) (*qotgetcompanyprofile.S2C, error) {
	c2s := &qotgetcompanyprofile.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetCompanyProfile(ctx, c2s)
}

// GetCompanyExecutivesWithContext 3244 - gets the company executives with context.
//
// code: security code
func (sdk *SDK) GetCompanyExecutivesWithContext(ctx context.Context, code string) (*qotgetcompanyexecutives.S2C, error) {
	c2s := &qotgetcompanyexecutives.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetCompanyExecutives(ctx, c2s)
}

// GetCompanyExecutiveBackgroundWithContext 3245 - gets the company executive background with context.
//
// code: security code
func (sdk *SDK) GetCompanyExecutiveBackgroundWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetcompanyexecutivebackground.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetcompanyexecutivebackground.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetCompanyExecutiveBackground(ctx, &c2s)
}

// GetCompanyOperationalEfficiencyWithContext 3246 - gets the company operational efficiency with context.
//
// code: security code
func (sdk *SDK) GetCompanyOperationalEfficiencyWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetcompanyoperationalefficiency.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetcompanyoperationalefficiency.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetCompanyOperationalEfficiency(ctx, &c2s)
}

// GetTopTenBuySellBrokersWithContext 3247 - gets the top ten buy/sell brokers with context.
//
// code: security code
func (sdk *SDK) GetTopTenBuySellBrokersWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgettoptenbuysellbrokers.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgettoptenbuysellbrokers.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetTopTenBuySellBrokers(ctx, &c2s)
}

// GetDailyShortVolumeWithContext 3248 - gets the daily short volume with context.
//
// code: security code
func (sdk *SDK) GetDailyShortVolumeWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetdailyshortvolume.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetdailyshortvolume.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetDailyShortVolume(ctx, &c2s)
}

// GetShortInterestWithContext 3249 - gets the short interest with context.
//
// code: security code
func (sdk *SDK) GetShortInterestWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetshortinterest.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetshortinterest.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetShortInterest(ctx, &c2s)
}

// GetOptionVolatilityWithContext 3250 - gets the option volatility with context.
//
// code: security code
func (sdk *SDK) GetOptionVolatilityWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgetoptionvolatility.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetoptionvolatility.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionVolatility(ctx, &c2s)
}

// GetOptionExerciseProbabilityWithContext 3251 - gets the option exercise probability with context.
//
// code: security code
func (sdk *SDK) GetOptionExerciseProbabilityWithContext(ctx context.Context, code string) (*qotgetoptionexerciseprobability.S2C, error) {
	c2s := &qotgetoptionexerciseprobability.C2S{
		Security: adapt.NewSecurity(code),
	}

	return sdk.cli.QotGetOptionExerciseProbability(ctx, c2s)
}

// StockScreenWithContext 3252 - stock screen with context.
func (sdk *SDK) StockScreenWithContext(ctx context.Context, opts ...adapt.Option) (*qotstockscreen.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotstockscreen.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotStockScreen(ctx, &c2s)
}

// OptionScreenWithContext 3253 - option screen with context.
func (sdk *SDK) OptionScreenWithContext(ctx context.Context, opts ...adapt.Option) (*qotoptionscreen.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotoptionscreen.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotOptionScreen(ctx, &c2s)
}

// WarrantScreenWithContext 3254 - warrant screen with context.
func (sdk *SDK) WarrantScreenWithContext(ctx context.Context, opts ...adapt.Option) (*qotwarrantscreen.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotwarrantscreen.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotWarrantScreen(ctx, &c2s)
}

// GetComboMaxTrdQtysWithContext 2112 - gets the max tradable quantity of a combo order with context.
//
// header: trading header
//
// qty: quantity
//
// orderType: order type, see adapt.OrderType_*
func (sdk *SDK) GetComboMaxTrdQtysWithContext(ctx context.Context, header *trdcommon.TrdHeader, qty float64, orderType int32, opts ...adapt.Option) (*trdgetcombomaxtrdqtys.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["qty"] = qty
	o["orderType"] = orderType

	var c2s trdgetcombomaxtrdqtys.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.TrdGetComboMaxTrdQtys(ctx, &c2s)
}

// PlaceComboOrderWithContext 2227 - places the combo order with context.
//
// header: trading header
//
// qty: quantity
//
// orderType: order type, see adapt.OrderType_*
func (sdk *SDK) PlaceComboOrderWithContext(ctx context.Context, header *trdcommon.TrdHeader, qty float64, orderType int32, opts ...adapt.Option) (*trdplacecomboorder.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["header"] = header
	o["qty"] = qty
	o["orderType"] = orderType

	var c2s trdplacecomboorder.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.TrdPlaceComboOrder(ctx, &c2s)
}

// GetOptionQuoteWithContext 3255 - gets the option quote with context.
func (sdk *SDK) GetOptionQuoteWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetoptionquote.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetoptionquote.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionQuote(ctx, &c2s)
}

// GetOptionStrategyWithContext 3256 - gets the option strategy with context.
//
// code: security code, e.g. HK.00700
//
// optionStrategy: option strategy type
func (sdk *SDK) GetOptionStrategyWithContext(ctx context.Context, code string, optionStrategy int32, opts ...adapt.Option) (*qotgetoptionstrategy.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)
	o["option_strategy"] = optionStrategy

	var c2s qotgetoptionstrategy.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionStrategy(ctx, &c2s)
}

// GetOptionStrategyAnalysisWithContext 3257 - gets the option strategy analysis with context.
func (sdk *SDK) GetOptionStrategyAnalysisWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetoptionstrategyanalysis.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetoptionstrategyanalysis.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionStrategyAnalysis(ctx, &c2s)
}

// GetOptionStrategySpreadWithContext 3258 - gets the option strategy spread with context.
//
// code: security code, e.g. HK.00700
//
// optionStrategy: option strategy type
func (sdk *SDK) GetOptionStrategySpreadWithContext(ctx context.Context, code string, optionStrategy int32, opts ...adapt.Option) (*qotgetoptionstrategyspreads.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)
	o["option_strategy"] = optionStrategy

	var c2s qotgetoptionstrategyspreads.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionStrategySpread(ctx, &c2s)
}

// GetIndicatorListWithContext 3259 - gets the indicator list with context.
func (sdk *SDK) GetIndicatorListWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetindicatorlist.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetindicatorlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetIndicatorList(ctx, &c2s)
}

// RequestIndicatorCalcWithContext 3260 - requests an asynchronous indicator calculation with context.
//
// shortName: short name of the indicator
//
// langType: script language type of the indicator
//
// data: indicator calculation data
func (sdk *SDK) RequestIndicatorCalcWithContext(ctx context.Context, shortName string, langType int32, data *qotrequestindicatorcalc.IndicatorCalcData, opts ...adapt.Option) (*qotrequestindicatorcalc.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["shortName"] = shortName
	o["langType"] = langType
	o["data"] = data

	var c2s qotrequestindicatorcalc.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotRequestIndicatorCalc(ctx, &c2s)
}

// GetSearchQuoteWithContext 3262 - gets the search quote with context.
//
// keyword: search keyword
func (sdk *SDK) GetSearchQuoteWithContext(ctx context.Context, keyword string, opts ...adapt.Option) (*qotgetsearchquote.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["keyword"] = keyword

	var c2s qotgetsearchquote.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetSearchQuote(ctx, &c2s)
}

// GetSearchNewsWithContext 3263 - gets the search news with context.
//
// keyword: search keyword
func (sdk *SDK) GetSearchNewsWithContext(ctx context.Context, keyword string, opts ...adapt.Option) (*qotgetsearchnews.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["keyword"] = keyword

	var c2s qotgetsearchnews.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetSearchNews(ctx, &c2s)
}

// GetOptionMarketStatisticWithContext 3301 - gets the option market statistic with context.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// dataType: statistic data type
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) GetOptionMarketStatisticWithContext(ctx context.Context, optionMarket int32, dataType int32, beginTime string, endTime string, opts ...adapt.Option) (*qotgetoptionmarketstatistic.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket
	o["dataType"] = dataType
	o["beginTime"] = beginTime
	o["endTime"] = endTime

	var c2s qotgetoptionmarketstatistic.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionMarketStatistic(ctx, &c2s)
}

// GetOptionUnderlyingHisStatisticWithContext 3302 - gets the option underlying historical statistic with context.
//
// code: security code, e.g. HK.00700
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) GetOptionUnderlyingHisStatisticWithContext(ctx context.Context, code string, beginTime string, endTime string, opts ...adapt.Option) (*qotgetoptionunderlyinghisstatistic.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)
	o["beginTime"] = beginTime
	o["endTime"] = endTime

	var c2s qotgetoptionunderlyinghisstatistic.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionUnderlyingHisStatistic(ctx, &c2s)
}

// GetOptionUnderlyingOverviewWithContext 3303 - gets the option underlying overview with context.
func (sdk *SDK) GetOptionUnderlyingOverviewWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetoptionunderlyingoverview.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetoptionunderlyingoverview.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionUnderlyingOverview(ctx, &c2s)
}

// GetOptionUnderlyingHisVolatilityWithContext 3304 - gets the option underlying historical volatility with context.
//
// code: security code, e.g. HK.00700
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) GetOptionUnderlyingHisVolatilityWithContext(ctx context.Context, code string, beginTime string, endTime string, opts ...adapt.Option) (*qotgetoptionunderlyinghisvolatility.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)
	o["beginTime"] = beginTime
	o["endTime"] = endTime

	var c2s qotgetoptionunderlyinghisvolatility.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionUnderlyingHisVolatility(ctx, &c2s)
}

// GetOptionUnderlyingRankWithContext 3305 - gets the option underlying rank with context.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// sortType: sort field
func (sdk *SDK) GetOptionUnderlyingRankWithContext(ctx context.Context, optionMarket int32, sortType int32, opts ...adapt.Option) (*qotgetoptionunderlyingrank.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket
	o["sortType"] = sortType

	var c2s qotgetoptionunderlyingrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionUnderlyingRank(ctx, &c2s)
}

// GetOptionRankWithContext 3306 - gets the option rank with context.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// sortType: sort field
func (sdk *SDK) GetOptionRankWithContext(ctx context.Context, optionMarket int32, sortType int32, opts ...adapt.Option) (*qotgetoptionrank.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket
	o["sortType"] = sortType

	var c2s qotgetoptionrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionRank(ctx, &c2s)
}

// GetOptionEventWithContext 3307 - gets the option event with context.
//
// optionMarket: option market, see adapt.QotMarket_*
func (sdk *SDK) GetOptionEventWithContext(ctx context.Context, optionMarket int32, opts ...adapt.Option) (*qotgetoptionevent.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket

	var c2s qotgetoptionevent.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionEvent(ctx, &c2s)
}

// GetOptionEventAlertWithContext 3308 - gets the option event alert with context.
func (sdk *SDK) GetOptionEventAlertWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetoptioneventalert.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetoptioneventalert.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionEventAlert(ctx, &c2s)
}

// SetOptionEventAlertWithContext 3309 - sets the option event alert with context.
//
// operType: operation type
func (sdk *SDK) SetOptionEventAlertWithContext(ctx context.Context, operType int32, opts ...adapt.Option) error {
	o := adapt.NewOptions(opts...)
	o["operType"] = operType

	var c2s qotsetoptioneventalert.C2S
	if err := o.ToProto(&c2s); err != nil {
		return err
	}

	return sdk.cli.QotSetOptionEventAlert(ctx, &c2s)
}

// GetOptionZeroDteScreenerWithContext 3311 - gets the zero DTE option screener with context.
//
// optionMarket: option market, see adapt.QotMarket_*
func (sdk *SDK) GetOptionZeroDteScreenerWithContext(ctx context.Context, optionMarket int32, opts ...adapt.Option) (*qotgetoptionzerodtescreener.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket

	var c2s qotgetoptionzerodtescreener.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionZeroDteScreener(ctx, &c2s)
}

// GetOptionZeroDteContractWithContext 3312 - gets the zero DTE option contract with context.
//
// code: security code, e.g. HK.00700
//
// strikeDateTimestamp: strike date timestamp in seconds
//
// chainInfo: option chain info
func (sdk *SDK) GetOptionZeroDteContractWithContext(ctx context.Context, code string, strikeDateTimestamp int64, chainInfo *qotgetoptionzerodtescreener.OptionChainInfo, opts ...adapt.Option) (*qotgetoptionzerodtecontract.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["owner"] = adapt.NewSecurity(code)
	o["strikeDateTimestamp"] = strikeDateTimestamp
	o["chainInfo"] = chainInfo

	var c2s qotgetoptionzerodtecontract.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionZeroDteContract(ctx, &c2s)
}

// GetOptionEarningsScreenerWithContext 3313 - gets the option earnings screener with context.
//
// optionMarket: option market, see adapt.QotMarket_*
func (sdk *SDK) GetOptionEarningsScreenerWithContext(ctx context.Context, optionMarket int32, opts ...adapt.Option) (*qotgetoptionearningsscreener.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket

	var c2s qotgetoptionearningsscreener.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionEarningsScreener(ctx, &c2s)
}

// GetOptionSellerScreenerWithContext 3314 - gets the option seller screener with context.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// sellerType: seller strategy type
func (sdk *SDK) GetOptionSellerScreenerWithContext(ctx context.Context, optionMarket int32, sellerType int32, opts ...adapt.Option) (*qotgetoptionsellerscreener.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["optionMarket"] = optionMarket
	o["sellerType"] = sellerType

	var c2s qotgetoptionsellerscreener.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetOptionSellerScreener(ctx, &c2s)
}

// GetEarningsCalendarWithContext 3401 - gets the earnings calendar with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetEarningsCalendarWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgetearningscalendar.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgetearningscalendar.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEarningsCalendar(ctx, &c2s)
}

// GetMacroIndicatorListWithContext 3402 - gets the macro indicator list with context.
//
// region: region of the macro indicator
func (sdk *SDK) GetMacroIndicatorListWithContext(ctx context.Context, region int32, opts ...adapt.Option) (*qotgetmacroindicatorlist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["region"] = region

	var c2s qotgetmacroindicatorlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetMacroIndicatorList(ctx, &c2s)
}

// GetMacroIndicatorHistoryWithContext 3403 - gets the macro indicator history with context.
//
// indicatorId: macro indicator ID
func (sdk *SDK) GetMacroIndicatorHistoryWithContext(ctx context.Context, indicatorId uint64, opts ...adapt.Option) (*qotgetmacroindicatorhistory.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["indicatorId"] = indicatorId

	var c2s qotgetmacroindicatorhistory.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetMacroIndicatorHistory(ctx, &c2s)
}

// GetFedWatchTargetRateWithContext 3404 - gets the fed watch target rate with context.
func (sdk *SDK) GetFedWatchTargetRateWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetfedwatchtargetrate.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetfedwatchtargetrate.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetFedWatchTargetRate(ctx, &c2s)
}

// GetFedWatchDotPlotWithContext 3405 - gets the fed watch dot plot with context.
func (sdk *SDK) GetFedWatchDotPlotWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetfedwatchdotplot.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetfedwatchdotplot.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetFedWatchDotPlot(ctx, &c2s)
}

// GetEarningsBeatRankWithContext 3406 - gets the earnings beat rank with context.
//
// market: market, see adapt.QotMarket_*
//
// beatType: earnings beat type
func (sdk *SDK) GetEarningsBeatRankWithContext(ctx context.Context, market int32, beatType int32, opts ...adapt.Option) (*qotgetearningsbeatrank.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market
	o["beatType"] = beatType

	var c2s qotgetearningsbeatrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEarningsBeatRank(ctx, &c2s)
}

// GetDividendRankWithContext 3407 - gets the dividend rank with context.
//
// market: market, see adapt.QotMarket_*
//
// rankType: dividend rank type
func (sdk *SDK) GetDividendRankWithContext(ctx context.Context, market int32, rankType int32, opts ...adapt.Option) (*qotgetdividendrank.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market
	o["rankType"] = rankType

	var c2s qotgetdividendrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetDividendRank(ctx, &c2s)
}

// GetDividendCalendarWithContext 3408 - gets the dividend calendar with context.
//
// market: market, see adapt.QotMarket_*
//
// date: date, format: yyyy-MM-dd
func (sdk *SDK) GetDividendCalendarWithContext(ctx context.Context, market int32, date string, opts ...adapt.Option) (*qotgetdividendcalendar.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market
	o["date"] = date

	var c2s qotgetdividendcalendar.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetDividendCalendar(ctx, &c2s)
}

// GetEconomicCalendarWithContext 3409 - gets the economic calendar with context.
//
// beginDate: begin date, format: yyyy-MM-dd
func (sdk *SDK) GetEconomicCalendarWithContext(ctx context.Context, beginDate string, opts ...adapt.Option) (*qotgeteconomiccalendar.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["beginDate"] = beginDate

	var c2s qotgeteconomiccalendar.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEconomicCalendar(ctx, &c2s)
}

// GetUSPreMarketRankWithContext 3410 - gets the US pre market rank with context.
func (sdk *SDK) GetUSPreMarketRankWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetuspremarketrank.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetuspremarketrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetUSPreMarketRank(ctx, &c2s)
}

// GetUSAfterHoursRankWithContext 3411 - gets the US after hours rank with context.
func (sdk *SDK) GetUSAfterHoursRankWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetusafterhoursrank.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetusafterhoursrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetUSAfterHoursRank(ctx, &c2s)
}

// GetUSOvernightRankWithContext 3412 - gets the US overnight rank with context.
func (sdk *SDK) GetUSOvernightRankWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetusovernightrank.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetusovernightrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetUSOvernightRank(ctx, &c2s)
}

// GetTopMoversRankWithContext 3413 - gets the top movers rank with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetTopMoversRankWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgettopmoverrank.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgettopmoverrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetTopMoversRank(ctx, &c2s)
}

// GetHotListWithContext 3414 - gets the hot list with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetHotListWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgethotlist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgethotlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetHotList(ctx, &c2s)
}

// GetShortSellingRankWithContext 3415 - gets the short selling rank with context.
func (sdk *SDK) GetShortSellingRankWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetshortsellingrank.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetshortsellingrank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetShortSellingRank(ctx, &c2s)
}

// GetPeriodChangeRankWithContext 3416 - gets the period change rank with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetPeriodChangeRankWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgetperiodchangerank.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgetperiodchangerank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetPeriodChangeRank(ctx, &c2s)
}

// GetHighDividendSOERankWithContext 3417 - gets the high dividend SOE rank with context.
func (sdk *SDK) GetHighDividendSOERankWithContext(ctx context.Context, opts ...adapt.Option) (*qotgethighdividendsoerank.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgethighdividendsoerank.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetHighDividendSOERank(ctx, &c2s)
}

// GetInstitutionListWithContext 3418 - gets the institution list with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetInstitutionListWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgetinstitutionlist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgetinstitutionlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInstitutionList(ctx, &c2s)
}

// GetInstitutionProfileWithContext 3419 - gets the institution profile with context.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionProfileWithContext(ctx context.Context, market int32, institutionId int32) (*qotgetinstitutionprofile.S2C, error) {
	o := make(adapt.Options)
	o["market"] = market
	o["institutionId"] = institutionId

	var c2s qotgetinstitutionprofile.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInstitutionProfile(ctx, &c2s)
}

// GetInstitutionDistributionWithContext 3420 - gets the institution distribution with context.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionDistributionWithContext(ctx context.Context, market int32, institutionId int32) (*qotgetinstitutiondistribution.S2C, error) {
	o := make(adapt.Options)
	o["market"] = market
	o["institutionId"] = institutionId

	var c2s qotgetinstitutiondistribution.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInstitutionDistribution(ctx, &c2s)
}

// GetInstitutionHoldingChangeWithContext 3421 - gets the institution holding change with context.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionHoldingChangeWithContext(ctx context.Context, market int32, institutionId int32, opts ...adapt.Option) (*qotgetinstitutionholdingchange.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market
	o["institutionId"] = institutionId

	var c2s qotgetinstitutionholdingchange.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInstitutionHoldingChange(ctx, &c2s)
}

// GetInstitutionHoldingListWithContext 3422 - gets the institution holding list with context.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionHoldingListWithContext(ctx context.Context, market int32, institutionId int32, opts ...adapt.Option) (*qotgetinstitutionholdinglist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market
	o["institutionId"] = institutionId

	var c2s qotgetinstitutionholdinglist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetInstitutionHoldingList(ctx, &c2s)
}

// GetArkFundHoldingWithContext 3423 - gets the ARK fund holding with context.
func (sdk *SDK) GetArkFundHoldingWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetarkfundholding.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetarkfundholding.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetArkFundHolding(ctx, &c2s)
}

// GetArkStockDynamicWithContext 3424 - gets the ARK stock dynamic with context.
//
// code: security code, e.g. HK.00700
func (sdk *SDK) GetArkStockDynamicWithContext(ctx context.Context, code string) (*qotgetarkstockdynamic.S2C, error) {
	o := make(adapt.Options)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgetarkstockdynamic.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetArkStockDynamic(ctx, &c2s)
}

// GetArkActiveTransactionWithContext 3425 - gets the ARK active transaction with context.
func (sdk *SDK) GetArkActiveTransactionWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetarkactivetransaction.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetarkactivetransaction.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetArkActiveTransaction(ctx, &c2s)
}

// GetRatingChangeWithContext 3426 - gets the rating change with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetRatingChangeWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgetratingchange.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgetratingchange.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetRatingChange(ctx, &c2s)
}

// GetIndustrialChainListWithContext 3427 - gets the industrial chain list with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetIndustrialChainListWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgetindustrialchainlist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgetindustrialchainlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetIndustrialChainList(ctx, &c2s)
}

// GetIndustrialChainDetailWithContext 3428 - gets the industrial chain detail with context.
//
// chainId: industrial chain ID
func (sdk *SDK) GetIndustrialChainDetailWithContext(ctx context.Context, chainId int64) (*qotgetindustrialchaindetail.S2C, error) {
	o := make(adapt.Options)
	o["chainId"] = chainId

	var c2s qotgetindustrialchaindetail.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetIndustrialChainDetail(ctx, &c2s)
}

// GetIndustrialChainByPlateWithContext 3429 - gets the industrial chain by plate with context.
//
// plateId: industrial plate ID
func (sdk *SDK) GetIndustrialChainByPlateWithContext(ctx context.Context, plateId int64) (*qotgetindustrialchainbyplate.S2C, error) {
	o := make(adapt.Options)
	o["plateId"] = plateId

	var c2s qotgetindustrialchainbyplate.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetIndustrialChainByPlate(ctx, &c2s)
}

// GetIndustrialPlateInfoWithContext 3430 - gets the industrial plate info with context.
//
// plateId: industrial plate ID
func (sdk *SDK) GetIndustrialPlateInfoWithContext(ctx context.Context, plateId int64) (*qotgetindustrialplateinfo.S2C, error) {
	o := make(adapt.Options)
	o["plateId"] = plateId

	var c2s qotgetindustrialplateinfo.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetIndustrialPlateInfo(ctx, &c2s)
}

// GetIndustrialPlateStockWithContext 3431 - gets the industrial plate stock with context.
func (sdk *SDK) GetIndustrialPlateStockWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetindustrialplatestock.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetindustrialplatestock.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetIndustrialPlateStock(ctx, &c2s)
}

// GetHeatMapDataWithContext 3432 - gets the heat map data with context.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetHeatMapDataWithContext(ctx context.Context, market int32, opts ...adapt.Option) (*qotgetheatmapdata.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["market"] = market

	var c2s qotgetheatmapdata.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetHeatMapData(ctx, &c2s)
}

// GetRiseFallDistributionWithContext 3433 - gets the rise fall distribution with context.
func (sdk *SDK) GetRiseFallDistributionWithContext(ctx context.Context, opts ...adapt.Option) (*qotgetrisefalldistr.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgetrisefalldistr.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetRiseFallDistribution(ctx, &c2s)
}

// GetEventContractCategoryWithContext 3434 - gets the event contract category with context.
func (sdk *SDK) GetEventContractCategoryWithContext(ctx context.Context, opts ...adapt.Option) (*qotgeteventcontractcategory.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgeteventcontractcategory.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractCategory(ctx, &c2s)
}

// FilterCompetitionWithContext 3435 - filters the competition with context.
func (sdk *SDK) FilterCompetitionWithContext(ctx context.Context, opts ...adapt.Option) (*qotfiltercompetition.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotfiltercompetition.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotFilterCompetition(ctx, &c2s)
}

// GetEventContractSeriesListWithContext 3436 - gets the event contract series list with context.
func (sdk *SDK) GetEventContractSeriesListWithContext(ctx context.Context, opts ...adapt.Option) (*qotgeteventcontractserieslist.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgeteventcontractserieslist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractSeriesList(ctx, &c2s)
}

// GetEventContractEventListWithContext 3437 - gets the event contract event list with context.
//
// code: series code of the event contract
func (sdk *SDK) GetEventContractEventListWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgeteventcontracteventlist.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["series"] = adapt.NewSecurity(code)

	var c2s qotgeteventcontracteventlist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractEventList(ctx, &c2s)
}

// GetEventContractWithContext 3438 - gets the event contract with context.
//
// code: event code of the event contract
func (sdk *SDK) GetEventContractWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgeteventcontract.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["event"] = adapt.NewSecurity(code)

	var c2s qotgeteventcontract.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContract(ctx, &c2s)
}

// GetEventContractMilestoneListWithContext 3439 - gets the event contract milestone list with context.
func (sdk *SDK) GetEventContractMilestoneListWithContext(ctx context.Context, opts ...adapt.Option) (*qotgeteventcontractmilestonelist.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgeteventcontractmilestonelist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractMilestoneList(ctx, &c2s)
}

// GetEventContractSnapshotWithContext 3445 - gets the event contract snapshot with context.
func (sdk *SDK) GetEventContractSnapshotWithContext(ctx context.Context, opts ...adapt.Option) (*qotgeteventcontractsnapshot.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgeteventcontractsnapshot.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractSnapshot(ctx, &c2s)
}

// GetEventContractOrderBookWithContext 3446 - gets the event contract order book with context.
//
// code: security code, e.g. HK.00700
//
// num: number of the order book levels
func (sdk *SDK) GetEventContractOrderBookWithContext(ctx context.Context, code string, num int32) (*qotgeteventcontractorderbook.S2C, error) {
	o := make(adapt.Options)
	o["security"] = adapt.NewSecurity(code)
	o["num"] = num

	var c2s qotgeteventcontractorderbook.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractOrderBook(ctx, &c2s)
}

// GetEventContractKlineWithContext 3447 - gets the event contract K line with context.
//
// code: security code, e.g. HK.00700
func (sdk *SDK) GetEventContractKlineWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgeteventcontractkline.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgeteventcontractkline.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractKline(ctx, &c2s)
}

// GetEventContractTickerWithContext 3448 - gets the event contract ticker with context.
//
// code: security code, e.g. HK.00700
func (sdk *SDK) GetEventContractTickerWithContext(ctx context.Context, code string, opts ...adapt.Option) (*qotgeteventcontractticker.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)

	var c2s qotgeteventcontractticker.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractTicker(ctx, &c2s)
}

// GetEventContractComboListWithContext 3453 - gets the event contract combo list with context.
func (sdk *SDK) GetEventContractComboListWithContext(ctx context.Context, opts ...adapt.Option) (*qotgeteventcontractcombolist.S2C, error) {
	o := adapt.NewOptions(opts...)

	var c2s qotgeteventcontractcombolist.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractComboList(ctx, &c2s)
}

// GetEventContractComboRfqWithContext 3454 - gets the event contract combo RFQ with context.
//
// mvc: combo identifier returned by the valid combo list
func (sdk *SDK) GetEventContractComboRfqWithContext(ctx context.Context, mvc string, opts ...adapt.Option) (*qotgeteventcontractcomborfq.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["mvc"] = mvc

	var c2s qotgeteventcontractcomborfq.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotGetEventContractComboRfq(ctx, &c2s)
}

// SubEventContractWithContext 3455 - subscribes or unsubscribes event contract with context.
//
// isSubOrUnSub: true to subscribe, false to unsubscribe
func (sdk *SDK) SubEventContractWithContext(ctx context.Context, isSubOrUnSub bool, opts ...adapt.Option) error {
	o := adapt.NewOptions(opts...)
	o["isSubOrUnSub"] = isSubOrUnSub

	var c2s qotsubeventcontract.C2S
	if err := o.ToProto(&c2s); err != nil {
		return err
	}

	return sdk.cli.QotSubEventContract(ctx, &c2s)
}

// RequestHistoryEventContractKLWithContext 3456 - requests the history K line of an event contract with context.
//
// code: security code, e.g. HK.00700
//
// klType: K line type, see adapt.KLType_*
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) RequestHistoryEventContractKLWithContext(ctx context.Context, code string, klType int32, beginTime string, endTime string, opts ...adapt.Option) (*qotrequesthistoryeventcontractkl.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["klType"] = klType
	o["beginTime"] = beginTime
	o["endTime"] = endTime

	var c2s qotrequesthistoryeventcontractkl.C2S
	if err := o.ToProto(&c2s); err != nil {
		return nil, err
	}

	return sdk.cli.QotRequestHistoryEventContractKL(ctx, &c2s)
}
