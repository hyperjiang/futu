package futu

import (
	"context"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetbroker"
	"github.com/hyperjiang/futu/pb/qotgetcapitaldistribution"
	"github.com/hyperjiang/futu/pb/qotgetcapitalflow"
	"github.com/hyperjiang/futu/pb/qotgetfutureinfo"
	"github.com/hyperjiang/futu/pb/qotgetipolist"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetmarketstate"
	"github.com/hyperjiang/futu/pb/qotgetoptionchain"
	"github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetownerplate"
	"github.com/hyperjiang/futu/pb/qotgetplatesecurity"
	"github.com/hyperjiang/futu/pb/qotgetplateset"
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
	"github.com/hyperjiang/futu/pb/qotgetreference"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetstaticinfo"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgetusersecurity"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotmodifyusersecurity"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotsetpricereminder"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdflowsummary"
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
