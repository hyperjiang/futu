package futu

import (
	"context"
	"time"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
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
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
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
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgettoptenbuysellbrokers"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetvaluationdetail"
	"github.com/hyperjiang/futu/pb/qotgetvaluationplatestocklist"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotoptionscreen"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotstockscreen"
	"github.com/hyperjiang/futu/pb/qotwarrantscreen"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdflowsummary"
	"github.com/hyperjiang/futu/pb/trdgetmarginratio"
	"github.com/hyperjiang/futu/pb/trdmodifyorder"
	"github.com/hyperjiang/futu/pb/trdplaceorder"
)

const defaultTimeout = time.Second * 5

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

// SDK is Futu SDK.
type SDK struct {
	client.Options

	cli *client.Client
}

// NewSDK creates a new Futu SDK.
func NewSDK(opts ...client.Option) (*SDK, error) {
	cli, err := client.New(opts...)
	if err != nil {
		return nil, err
	}

	return &SDK{cli: cli}, nil
}

// Close closes the client.
func (sdk *SDK) Close() error {
	return sdk.cli.Close()
}

// GetClient returns the client.
func (sdk *SDK) GetClient() *client.Client {
	return sdk.cli
}

// RegisterHandler registers a handler for notifications of a specified protoID.
func (sdk *SDK) RegisterHandler(protoID uint32, h client.Handler) *SDK {
	sdk.cli.RegisterHandler(protoID, h)

	return sdk
}

// GetGlobalState 1002 - gets the global state.
func (sdk *SDK) GetGlobalState() (*getglobalstate.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetGlobalStateWithContext(ctx)
}

// GetAccList 2001 - gets the trading account list.
func (sdk *SDK) GetAccList(opts ...adapt.Option) ([]*trdcommon.TrdAcc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetAccListWithContext(ctx, opts...)
}

// UnlockTrade 2005 - unlocks or locks the trade.
//
// unlock: true for unlock, false for lock
//
// pwdMD5: MD5 of the password
//
// securityFirm: security firm
func (sdk *SDK) UnlockTrade(unlock bool, pwdMD5 string, securityFirm int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.UnlockTradeWithContext(ctx, unlock, pwdMD5, securityFirm)
}

// SubscribeAccPush 2008 - subscribes the trading account push data.
//
// accIDList: account ID list
func (sdk *SDK) SubscribeAccPush(accIDList []uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.SubscribeAccPushWithContext(ctx, accIDList)
}

// GetFunds 2101 - gets the funds.
func (sdk *SDK) GetFunds(header *trdcommon.TrdHeader, opts ...adapt.Option) (*trdcommon.Funds, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFundsWithContext(ctx, header, opts...)
}

// GetPositionList 2102 - gets the position list.
func (sdk *SDK) GetPositionList(header *trdcommon.TrdHeader, opts ...adapt.Option) ([]*trdcommon.Position, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetPositionListWithContext(ctx, header, opts...)
}

// GetMaxTrdQtys 2111 - gets the maximum available trading quantities.
//
// header: trading header
//
// orderType: order type
//
// code: security code, e.g. US.AAPL
//
// price: price
func (sdk *SDK) GetMaxTrdQtys(header *trdcommon.TrdHeader, orderType int32, code string, price float64, opts ...adapt.Option) (*trdcommon.MaxTrdQtys, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetMaxTrdQtysWithContext(ctx, header, orderType, code, price, opts...)
}

// GetOpenOrderList 2201 - gets the open order list.
func (sdk *SDK) GetOpenOrderList(header *trdcommon.TrdHeader, opts ...adapt.Option) ([]*trdcommon.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOpenOrderListWithContext(ctx, header, opts...)
}

// PlaceOrder 2202 - places an order.
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
func (sdk *SDK) PlaceOrder(header *trdcommon.TrdHeader, trdSide int32, orderType int32, code string, qty float64, price float64, opts ...adapt.Option) (*trdplaceorder.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.PlaceOrderWithContext(ctx, header, trdSide, orderType, code, qty, price, opts...)
}

// ModifyOrder 2205 - modifies an order with context.
//
// header: trading header
//
// orderID: order ID, use 0 if forAll=true
//
// modifyOrderOp: modify order operation
func (sdk *SDK) ModifyOrder(header *trdcommon.TrdHeader, orderID uint64, modifyOrderOp int32, opts ...adapt.Option) (*trdmodifyorder.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.ModifyOrderWithContext(ctx, header, orderID, modifyOrderOp, opts...)
}

// GetOrderFillList 2211 - gets the filled order list.
func (sdk *SDK) GetOrderFillList(header *trdcommon.TrdHeader, opts ...adapt.Option) ([]*trdcommon.OrderFill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOrderFillListWithContext(ctx, header, opts...)
}

// GetHistoryOrderList 2221 - gets the history order list.
func (sdk *SDK) GetHistoryOrderList(header *trdcommon.TrdHeader, fc *trdcommon.TrdFilterConditions, opts ...adapt.Option) ([]*trdcommon.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetHistoryOrderListWithContext(ctx, header, fc, opts...)
}

// GetHistoryOrderFillList 2222 - gets the history filled order list.
func (sdk *SDK) GetHistoryOrderFillList(header *trdcommon.TrdHeader, fc *trdcommon.TrdFilterConditions, opts ...adapt.Option) ([]*trdcommon.OrderFill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetHistoryOrderFillListWithContext(ctx, header, fc, opts...)
}

// GetMarginRatio 2223 - gets the margin ratio.
func (sdk *SDK) GetMarginRatio(header *trdcommon.TrdHeader, codes []string) ([]*trdgetmarginratio.MarginRatioInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetMarginRatioWithContext(ctx, header, codes)
}

// GetOrderFee 2225 - gets the order fee.
func (sdk *SDK) GetOrderFee(header *trdcommon.TrdHeader, orderIdExList []string) ([]*trdcommon.OrderFee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOrderFeeWithContext(ctx, header, orderIdExList)
}

// TrdFlowSummary 2226 - gets the trading flow summary.
func (sdk *SDK) TrdFlowSummary(header *trdcommon.TrdHeader, clearingDate string) ([]*trdflowsummary.FlowSummaryInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.TrdFlowSummaryWithContext(ctx, header, clearingDate)
}

// Subscribe 3001 - subscribes or unsubscribes.
//
// codes: security codes
//
// subTypes: subscription types
//
// isSub: true for subscribe, false for unsubscribe
func (sdk *SDK) Subscribe(codes []string, subTypes []int32, isSub bool, opts ...adapt.Option) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.SubscribeWithContext(ctx, codes, subTypes, isSub, opts...)
}

// GetSubInfo 3003 - gets the subscription information.
func (sdk *SDK) GetSubInfo(opts ...adapt.Option) (*qotgetsubinfo.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetSubInfoWithContext(ctx, opts...)
}

// GetBasicQot 3004 - gets the basic quotes of given securities.
func (sdk *SDK) GetBasicQot(codes []string) ([]*qotcommon.BasicQot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetBasicQotWithContext(ctx, codes)
}

// GetKL 3006 - gets K-line data.
//
// code: security code
//
// klType: K-line type
func (sdk *SDK) GetKL(code string, klType int32, opts ...adapt.Option) (*qotgetkl.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetKLWithContext(ctx, code, klType, opts...)
}

// GetRT 3008 - gets real-time data.
//
// code: security code
func (sdk *SDK) GetRT(code string) (*qotgetrt.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetRTWithContext(ctx, code)
}

// GetTicker 3010 - gets ticker data.
//
// code: security code
func (sdk *SDK) GetTicker(code string, opts ...adapt.Option) (*qotgetticker.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetTickerWithContext(ctx, code, opts...)
}

// GetOrderBook 3012 - gets order book data.
//
// code: security code
func (sdk *SDK) GetOrderBook(code string, opts ...adapt.Option) (*qotgetorderbook.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOrderBookWithContext(ctx, code, opts...)
}

// GetBroker 3014 - gets broker data.
//
// code: security code
func (sdk *SDK) GetBroker(code string) (*qotgetbroker.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetBrokerWithContext(ctx, code)
}

// RequestHistoryKL 3103 - requests the history K-line data.
//
// code: security code
//
// klType: K-line type
//
// beginTime: begin time, format: "yyyy-MM-dd"
//
// endTime: end time, format: "yyyy-MM-dd"
func (sdk *SDK) RequestHistoryKL(code string, klType int32, beginTime string, endTime string, opts ...adapt.Option) (*qotrequesthistorykl.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestHistoryKLWithContext(ctx, code, klType, beginTime, endTime, opts...)
}

// RequestHistoryKLQuota 3104 - requests the history K-line quota.
func (sdk *SDK) RequestHistoryKLQuota(opts ...adapt.Option) (*qotrequesthistoryklquota.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestHistoryKLQuotaWithContext(ctx, opts...)
}

// RequestRehab 3105 - requests the rehab data.
func (sdk *SDK) RequestRehab(code string) (*qotrequestrehab.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestRehabWithContext(ctx, code)
}

// GetStaticInfo 3202 - gets the static information.
func (sdk *SDK) GetStaticInfo(opts ...adapt.Option) ([]*qotcommon.SecurityStaticInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetStaticInfoWithContext(ctx, opts...)
}

// GetSecuritySnapshot 3203 - gets the security snapshot.
//
// codes: security codes
func (sdk *SDK) GetSecuritySnapshot(codes []string) ([]*qotgetsecuritysnapshot.Snapshot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetSecuritySnapshotWithContext(ctx, codes)
}

// GetPlateSet 3204 - gets the plate set.
//
// market: market
//
// plateSetType: plate set type
func (sdk *SDK) GetPlateSet(market int32, plateSetType int32) ([]*qotcommon.PlateInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetPlateSetWithContext(ctx, market, plateSetType)
}

// GetPlateSecurity 3205 - gets the plate securities.
//
// plateCode: plate code
func (sdk *SDK) GetPlateSecurity(plateCode string, opts ...adapt.Option) ([]*qotcommon.SecurityStaticInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetPlateSecurityWithContext(ctx, plateCode, opts...)
}

// GetReference 3206 - gets the reference data.
//
// code: security code
//
// refType: reference type
func (sdk *SDK) GetReference(code string, refType int32) ([]*qotcommon.SecurityStaticInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetReferenceWithContext(ctx, code, refType)
}

// GetOwnerPlate 3207 - gets the owner plate.
//
// codes: security codes
func (sdk *SDK) GetOwnerPlate(codes []string) ([]*qotgetownerplate.SecurityOwnerPlate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOwnerPlateWithContext(ctx, codes)
}

// GetOptionChain 3209 - gets the option chain with context.
//
// code: security code
//
// beginTime: begin time, format: "yyyy-MM-dd"
//
// endTime: end time, format: "yyyy-MM-dd"
func (sdk *SDK) GetOptionChain(code string, beginTime string, endTime string, opts ...adapt.Option) ([]*qotgetoptionchain.OptionChain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionChainWithContext(ctx, code, beginTime, endTime, opts...)
}

// GetWarrant 3210 - gets the warrant, only available in Hong Kong market.
// Sort by score in descending order by default.
//
// begin: begin index
//
// num: number of warrants
func (sdk *SDK) GetWarrant(begin int32, num int32, opts ...adapt.Option) (*qotgetwarrant.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetWarrantWithContext(ctx, begin, num, opts...)
}

// GetCapitalFlow 3211 - gets the capital flow.
//
// code: security code
func (sdk *SDK) GetCapitalFlow(code string, opts ...adapt.Option) (*qotgetcapitalflow.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCapitalFlowWithContext(ctx, code, opts...)
}

// GetCapitalDistribution 3212 - gets the capital distribution.
//
// code: security code
func (sdk *SDK) GetCapitalDistribution(code string) (*qotgetcapitaldistribution.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCapitalDistributionWithContext(ctx, code)
}

// GetUserSecurity 3213 - gets the user security.
//
// groupName: group name
func (sdk *SDK) GetUserSecurity(groupName string) ([]*qotcommon.SecurityStaticInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetUserSecurityWithContext(ctx, groupName)
}

// ModifyUserSecurity 3214 - modifies the user security.
//
// groupName: group name
//
// codes: security codes
//
// op: operation, 1 for add, 2 for delete
func (sdk *SDK) ModifyUserSecurity(groupName string, codes []string, op int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.ModifyUserSecurityWithContext(ctx, groupName, codes, op)
}

// StockFilter 3215 - filters the stocks.
//
// market: market
func (sdk *SDK) StockFilter(market int32, opts ...adapt.Option) (*qotstockfilter.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.StockFilterWithContext(ctx, market, opts...)
}

// GetIpoList 3217 - gets the IPO list.
//
// market: market
func (sdk *SDK) GetIpoList(market int32) ([]*qotgetipolist.IpoData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIpoListWithContext(ctx, market)
}

// GetFutureInfo 3218 - gets the future information.
//
// codes: security codes
func (sdk *SDK) GetFutureInfo(codes []string) ([]*qotgetfutureinfo.FutureInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFutureInfoWithContext(ctx, codes)
}

// RequestTradeDate 3219 - requests the trade date.
//
// market: market
//
// code: security code
//
// beginTime: begin time, format: "yyyy-MM-dd"
//
// endTime: end time, format: "yyyy-MM-dd"
func (sdk *SDK) RequestTradeDate(market int32, code string, beginTime string, endTime string) ([]*qotrequesttradedate.TradeDate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestTradeDateWithContext(ctx, market, code, beginTime, endTime)
}

// SetPriceReminder 3220 - sets the price reminder.
//
// code: security code
//
// op: operation, 1 for add, 2 for delete
func (sdk *SDK) SetPriceReminder(code string, op int32, opts ...adapt.Option) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.SetPriceReminderWithContext(ctx, code, op, opts...)
}

// GetPriceReminder 3221 - gets the price reminder.
//
// code: security code
//
// market: market, if security is set, this param is ignored
func (sdk *SDK) GetPriceReminder(code string, market int32) ([]*qotgetpricereminder.PriceReminder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetPriceReminderWithContext(ctx, code, market)
}

// GetUserSecurityGroup 3222 - gets the user security group.
//
// groupType: group type
func (sdk *SDK) GetUserSecurityGroup(groupType int32) ([]*qotgetusersecuritygroup.GroupData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetUserSecurityGroupWithContext(ctx, groupType)
}

// GetMarketState 3223 - gets the market state.
//
// codes: security codes
func (sdk *SDK) GetMarketState(codes []string) ([]*qotgetmarketstate.MarketInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetMarketStateWithContext(ctx, codes)
}

// GetOptionExpirationDate 3224 - gets the option expiration date.
//
// code: security code
func (sdk *SDK) GetOptionExpirationDate(code string, opts ...adapt.Option) ([]*qotgetoptionexpirationdate.OptionExpirationDate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionExpirationDateWithContext(ctx, code, opts...)
}

// GetFinancialsEarningsPriceMove 3225 - gets the financials earnings price move.
//
// code: security code
func (sdk *SDK) GetFinancialsEarningsPriceMove(code string, opts ...adapt.Option) (*qotgetfinancialsearnpricemove.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFinancialsEarningsPriceMoveWithContext(ctx, code, opts...)
}

// GetFinancialsEarningsPriceHistory 3226 - gets the financials earnings price history.
//
// code: security code
func (sdk *SDK) GetFinancialsEarningsPriceHistory(code string) (*qotgetfinancialsearnpricehist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFinancialsEarningsPriceHistoryWithContext(ctx, code)
}

// GetFinancialsStatements 3227 - gets the financial statements.
//
// code: security code
//
// statementType: financial statement type, see adapt.FinancialStatementsType_*
func (sdk *SDK) GetFinancialsStatements(code string, statementType int32, opts ...adapt.Option) (*qotgetfinancialsstatements.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFinancialsStatementsWithContext(ctx, code, statementType, opts...)
}

// GetFinancialsRevenueBreakdown 3228 - gets the financials revenue breakdown.
//
// code: security code
func (sdk *SDK) GetFinancialsRevenueBreakdown(code string, opts ...adapt.Option) (*qotgetfinancialrevenuebreakdown.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFinancialsRevenueBreakdownWithContext(ctx, code, opts...)
}

// GetResearchAnalystConsensus 3229 - gets the research analyst consensus.
//
// code: security code
func (sdk *SDK) GetResearchAnalystConsensus(code string) (*qotgetresearchanalystconsensus.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetResearchAnalystConsensusWithContext(ctx, code)
}

// GetResearchRatingSummary 3230 - gets the research rating summary.
//
// code: security code
func (sdk *SDK) GetResearchRatingSummary(code string, opts ...adapt.Option) (*qotgetresearchratingsummary.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetResearchRatingSummaryWithContext(ctx, code, opts...)
}

// GetResearchMorningstarReport 3231 - gets the research morningstar report.
//
// code: security code
func (sdk *SDK) GetResearchMorningstarReport(code string) (*qotgetresearchmorningstarrpt.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetResearchMorningstarReportWithContext(ctx, code)
}

// GetValuationDetail 3232 - gets the valuation detail.
//
// code: security code
//
// valuationType: valuation type, see adapt.ValuationType_*
func (sdk *SDK) GetValuationDetail(code string, valuationType int32, opts ...adapt.Option) (*qotgetvaluationdetail.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetValuationDetailWithContext(ctx, code, valuationType, opts...)
}

// GetValuationPlateStockList 3233 - gets the valuation plate stock list.
//
// code: security code (plate code)
//
// valuationType: valuation type, see adapt.ValuationType_*
func (sdk *SDK) GetValuationPlateStockList(code string, valuationType int32, opts ...adapt.Option) (*qotgetvaluationplatestocklist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetValuationPlateStockListWithContext(ctx, code, valuationType, opts...)
}

// GetCorporateActionsDividends 3234 - gets the corporate actions dividends.
//
// code: security code
func (sdk *SDK) GetCorporateActionsDividends(code string) (*qotgetcorporateactionsdividends.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCorporateActionsDividendsWithContext(ctx, code)
}

// GetCorporateActionsBuybacks 3235 - gets the corporate actions buybacks.
//
// code: security code
func (sdk *SDK) GetCorporateActionsBuybacks(code string, opts ...adapt.Option) (*qotgetcorporateactionsbuybacks.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCorporateActionsBuybacksWithContext(ctx, code, opts...)
}

// GetCorporateActionsStockSplits 3236 - gets the corporate actions stock splits.
//
// code: security code
func (sdk *SDK) GetCorporateActionsStockSplits(code string, opts ...adapt.Option) (*qotgetcorporateactionsstocksplits.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCorporateActionsStockSplitsWithContext(ctx, code, opts...)
}

// GetShareholdersOverview 3237 - gets the shareholders overview.
//
// code: security code
func (sdk *SDK) GetShareholdersOverview(code string, opts ...adapt.Option) (*qotgetshareholdersoverview.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetShareholdersOverviewWithContext(ctx, code, opts...)
}

// GetShareholdersHoldingChanges 3238 - gets the shareholders holding changes.
//
// code: security code
func (sdk *SDK) GetShareholdersHoldingChanges(code string, opts ...adapt.Option) (*qotgetshareholdersholdingchanges.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetShareholdersHoldingChangesWithContext(ctx, code, opts...)
}

// GetShareholdersHolderDetail 3239 - gets the shareholders holder detail.
//
// code: security code
func (sdk *SDK) GetShareholdersHolderDetail(code string, opts ...adapt.Option) (*qotgetshareholdersholderdetail.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetShareholdersHolderDetailWithContext(ctx, code, opts...)
}

// GetShareholdersInstitutional 3240 - gets the shareholders institutional.
//
// code: security code
func (sdk *SDK) GetShareholdersInstitutional(code string, opts ...adapt.Option) (*qotgetshareholdersinstitutional.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetShareholdersInstitutionalWithContext(ctx, code, opts...)
}

// GetInsiderHolderList 3241 - gets the insider holder list.
//
// code: security code
func (sdk *SDK) GetInsiderHolderList(code string, opts ...adapt.Option) (*qotgetinsiderholderlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInsiderHolderListWithContext(ctx, code, opts...)
}

// GetInsiderTradeList 3242 - gets the insider trade list.
//
// code: security code
func (sdk *SDK) GetInsiderTradeList(code string, opts ...adapt.Option) (*qotgetinsidertradelist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInsiderTradeListWithContext(ctx, code, opts...)
}

// GetCompanyProfile 3243 - gets the company profile.
//
// code: security code
func (sdk *SDK) GetCompanyProfile(code string) (*qotgetcompanyprofile.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCompanyProfileWithContext(ctx, code)
}

// GetCompanyExecutives 3244 - gets the company executives.
//
// code: security code
func (sdk *SDK) GetCompanyExecutives(code string) (*qotgetcompanyexecutives.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCompanyExecutivesWithContext(ctx, code)
}

// GetCompanyExecutiveBackground 3245 - gets the company executive background.
//
// code: security code
func (sdk *SDK) GetCompanyExecutiveBackground(code string, opts ...adapt.Option) (*qotgetcompanyexecutivebackground.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCompanyExecutiveBackgroundWithContext(ctx, code, opts...)
}

// GetCompanyOperationalEfficiency 3246 - gets the company operational efficiency.
//
// code: security code
func (sdk *SDK) GetCompanyOperationalEfficiency(code string, opts ...adapt.Option) (*qotgetcompanyoperationalefficiency.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetCompanyOperationalEfficiencyWithContext(ctx, code, opts...)
}

// GetTopTenBuySellBrokers 3247 - gets the top ten buy/sell brokers.
//
// code: security code
func (sdk *SDK) GetTopTenBuySellBrokers(code string, opts ...adapt.Option) (*qotgettoptenbuysellbrokers.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetTopTenBuySellBrokersWithContext(ctx, code, opts...)
}

// GetDailyShortVolume 3248 - gets the daily short volume.
//
// code: security code
func (sdk *SDK) GetDailyShortVolume(code string, opts ...adapt.Option) (*qotgetdailyshortvolume.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetDailyShortVolumeWithContext(ctx, code, opts...)
}

// GetShortInterest 3249 - gets the short interest.
//
// code: security code
func (sdk *SDK) GetShortInterest(code string, opts ...adapt.Option) (*qotgetshortinterest.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetShortInterestWithContext(ctx, code, opts...)
}

// GetOptionVolatility 3250 - gets the option volatility.
//
// code: security code
func (sdk *SDK) GetOptionVolatility(code string, opts ...adapt.Option) (*qotgetoptionvolatility.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionVolatilityWithContext(ctx, code, opts...)
}

// GetOptionExerciseProbability 3251 - gets the option exercise probability.
//
// code: security code
func (sdk *SDK) GetOptionExerciseProbability(code string) (*qotgetoptionexerciseprobability.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionExerciseProbabilityWithContext(ctx, code)
}

// StockScreen 3252 - stock screen.
func (sdk *SDK) StockScreen(opts ...adapt.Option) (*qotstockscreen.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.StockScreenWithContext(ctx, opts...)
}

// OptionScreen 3253 - option screen.
func (sdk *SDK) OptionScreen(opts ...adapt.Option) (*qotoptionscreen.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.OptionScreenWithContext(ctx, opts...)
}

// WarrantScreen 3254 - warrant screen.
func (sdk *SDK) WarrantScreen(opts ...adapt.Option) (*qotwarrantscreen.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.WarrantScreenWithContext(ctx, opts...)
}
