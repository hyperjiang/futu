package futu

import (
	"context"
	"time"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotfiltercompetition"
	"github.com/hyperjiang/futu/pb/qotgetarkactivetransaction"
	"github.com/hyperjiang/futu/pb/qotgetarkfundholding"
	"github.com/hyperjiang/futu/pb/qotgetarkstockdynamic"
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
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
	"github.com/hyperjiang/futu/pb/qotgetratingchange"
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
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgettopmoverrank"
	"github.com/hyperjiang/futu/pb/qotgettoptenbuysellbrokers"
	"github.com/hyperjiang/futu/pb/qotgetusafterhoursrank"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetusovernightrank"
	"github.com/hyperjiang/futu/pb/qotgetuspremarketrank"
	"github.com/hyperjiang/futu/pb/qotgetvaluationdetail"
	"github.com/hyperjiang/futu/pb/qotgetvaluationplatestocklist"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotoptionscreen"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryeventcontractkl"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestindicatorcalc"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotstockscreen"
	"github.com/hyperjiang/futu/pb/qotwarrantscreen"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdflowsummary"
	"github.com/hyperjiang/futu/pb/trdgetcombomaxtrdqtys"
	"github.com/hyperjiang/futu/pb/trdgetmarginratio"
	"github.com/hyperjiang/futu/pb/trdmodifyorder"
	"github.com/hyperjiang/futu/pb/trdplacecomboorder"
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

// GetComboMaxTrdQtys 2112 - gets the max tradable quantity of a combo order.
//
// header: trading header
//
// qty: quantity
//
// orderType: order type, see adapt.OrderType_*
func (sdk *SDK) GetComboMaxTrdQtys(header *trdcommon.TrdHeader, qty float64, orderType int32, opts ...adapt.Option) (*trdgetcombomaxtrdqtys.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetComboMaxTrdQtysWithContext(ctx, header, qty, orderType, opts...)
}

// PlaceComboOrder 2227 - places the combo order.
//
// header: trading header
//
// qty: quantity
//
// orderType: order type, see adapt.OrderType_*
func (sdk *SDK) PlaceComboOrder(header *trdcommon.TrdHeader, qty float64, orderType int32, opts ...adapt.Option) (*trdplacecomboorder.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.PlaceComboOrderWithContext(ctx, header, qty, orderType, opts...)
}

// GetOptionQuote 3255 - gets the option quote.
func (sdk *SDK) GetOptionQuote(opts ...adapt.Option) (*qotgetoptionquote.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionQuoteWithContext(ctx, opts...)
}

// GetOptionStrategy 3256 - gets the option strategy.
//
// code: security code, e.g. HK.00700
//
// optionStrategy: option strategy type
func (sdk *SDK) GetOptionStrategy(code string, optionStrategy int32, opts ...adapt.Option) (*qotgetoptionstrategy.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionStrategyWithContext(ctx, code, optionStrategy, opts...)
}

// GetOptionStrategyAnalysis 3257 - gets the option strategy analysis.
func (sdk *SDK) GetOptionStrategyAnalysis(opts ...adapt.Option) (*qotgetoptionstrategyanalysis.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionStrategyAnalysisWithContext(ctx, opts...)
}

// GetOptionStrategySpread 3258 - gets the option strategy spread.
//
// code: security code, e.g. HK.00700
//
// optionStrategy: option strategy type
func (sdk *SDK) GetOptionStrategySpread(code string, optionStrategy int32, opts ...adapt.Option) (*qotgetoptionstrategyspreads.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionStrategySpreadWithContext(ctx, code, optionStrategy, opts...)
}

// GetIndicatorList 3259 - gets the indicator list.
func (sdk *SDK) GetIndicatorList(opts ...adapt.Option) (*qotgetindicatorlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIndicatorListWithContext(ctx, opts...)
}

// RequestIndicatorCalc 3260 - requests an asynchronous indicator calculation.
//
// shortName: short name of the indicator
//
// langType: script language type of the indicator
//
// data: indicator calculation data
func (sdk *SDK) RequestIndicatorCalc(shortName string, langType int32, data *qotrequestindicatorcalc.IndicatorCalcData, opts ...adapt.Option) (*qotrequestindicatorcalc.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestIndicatorCalcWithContext(ctx, shortName, langType, data, opts...)
}

// GetSearchQuote 3262 - gets the search quote.
//
// keyword: search keyword
func (sdk *SDK) GetSearchQuote(keyword string, opts ...adapt.Option) (*qotgetsearchquote.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetSearchQuoteWithContext(ctx, keyword, opts...)
}

// GetSearchNews 3263 - gets the search news.
//
// keyword: search keyword
func (sdk *SDK) GetSearchNews(keyword string, opts ...adapt.Option) (*qotgetsearchnews.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetSearchNewsWithContext(ctx, keyword, opts...)
}

// GetOptionMarketStatistic 3301 - gets the option market statistic.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// dataType: statistic data type
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) GetOptionMarketStatistic(optionMarket int32, dataType int32, beginTime string, endTime string, opts ...adapt.Option) (*qotgetoptionmarketstatistic.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionMarketStatisticWithContext(ctx, optionMarket, dataType, beginTime, endTime, opts...)
}

// GetOptionUnderlyingHisStatistic 3302 - gets the option underlying historical statistic.
//
// code: security code, e.g. HK.00700
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) GetOptionUnderlyingHisStatistic(code string, beginTime string, endTime string, opts ...adapt.Option) (*qotgetoptionunderlyinghisstatistic.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionUnderlyingHisStatisticWithContext(ctx, code, beginTime, endTime, opts...)
}

// GetOptionUnderlyingOverview 3303 - gets the option underlying overview.
func (sdk *SDK) GetOptionUnderlyingOverview(opts ...adapt.Option) (*qotgetoptionunderlyingoverview.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionUnderlyingOverviewWithContext(ctx, opts...)
}

// GetOptionUnderlyingHisVolatility 3304 - gets the option underlying historical volatility.
//
// code: security code, e.g. HK.00700
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) GetOptionUnderlyingHisVolatility(code string, beginTime string, endTime string, opts ...adapt.Option) (*qotgetoptionunderlyinghisvolatility.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionUnderlyingHisVolatilityWithContext(ctx, code, beginTime, endTime, opts...)
}

// GetOptionUnderlyingRank 3305 - gets the option underlying rank.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// sortType: sort field
func (sdk *SDK) GetOptionUnderlyingRank(optionMarket int32, sortType int32, opts ...adapt.Option) (*qotgetoptionunderlyingrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionUnderlyingRankWithContext(ctx, optionMarket, sortType, opts...)
}

// GetOptionRank 3306 - gets the option rank.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// sortType: sort field
func (sdk *SDK) GetOptionRank(optionMarket int32, sortType int32, opts ...adapt.Option) (*qotgetoptionrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionRankWithContext(ctx, optionMarket, sortType, opts...)
}

// GetOptionEvent 3307 - gets the option event.
//
// optionMarket: option market, see adapt.QotMarket_*
func (sdk *SDK) GetOptionEvent(optionMarket int32, opts ...adapt.Option) (*qotgetoptionevent.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionEventWithContext(ctx, optionMarket, opts...)
}

// GetOptionEventAlert 3308 - gets the option event alert.
func (sdk *SDK) GetOptionEventAlert(opts ...adapt.Option) (*qotgetoptioneventalert.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionEventAlertWithContext(ctx, opts...)
}

// SetOptionEventAlert 3309 - sets the option event alert.
//
// operType: operation type
func (sdk *SDK) SetOptionEventAlert(operType int32, opts ...adapt.Option) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.SetOptionEventAlertWithContext(ctx, operType, opts...)
}

// GetOptionZeroDteScreener 3311 - gets the zero DTE option screener.
//
// optionMarket: option market, see adapt.QotMarket_*
func (sdk *SDK) GetOptionZeroDteScreener(optionMarket int32, opts ...adapt.Option) (*qotgetoptionzerodtescreener.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionZeroDteScreenerWithContext(ctx, optionMarket, opts...)
}

// GetOptionZeroDteContract 3312 - gets the zero DTE option contract.
//
// code: security code, e.g. HK.00700
//
// strikeDateTimestamp: strike date timestamp in seconds
//
// chainInfo: option chain info
func (sdk *SDK) GetOptionZeroDteContract(code string, strikeDateTimestamp int64, chainInfo *qotgetoptionzerodtescreener.OptionChainInfo, opts ...adapt.Option) (*qotgetoptionzerodtecontract.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionZeroDteContractWithContext(ctx, code, strikeDateTimestamp, chainInfo, opts...)
}

// GetOptionEarningsScreener 3313 - gets the option earnings screener.
//
// optionMarket: option market, see adapt.QotMarket_*
func (sdk *SDK) GetOptionEarningsScreener(optionMarket int32, opts ...adapt.Option) (*qotgetoptionearningsscreener.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionEarningsScreenerWithContext(ctx, optionMarket, opts...)
}

// GetOptionSellerScreener 3314 - gets the option seller screener.
//
// optionMarket: option market, see adapt.QotMarket_*
//
// sellerType: seller strategy type
func (sdk *SDK) GetOptionSellerScreener(optionMarket int32, sellerType int32, opts ...adapt.Option) (*qotgetoptionsellerscreener.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetOptionSellerScreenerWithContext(ctx, optionMarket, sellerType, opts...)
}

// GetEarningsCalendar 3401 - gets the earnings calendar.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetEarningsCalendar(market int32, opts ...adapt.Option) (*qotgetearningscalendar.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEarningsCalendarWithContext(ctx, market, opts...)
}

// GetMacroIndicatorList 3402 - gets the macro indicator list.
//
// region: region of the macro indicator
func (sdk *SDK) GetMacroIndicatorList(region int32, opts ...adapt.Option) (*qotgetmacroindicatorlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetMacroIndicatorListWithContext(ctx, region, opts...)
}

// GetMacroIndicatorHistory 3403 - gets the macro indicator history.
//
// indicatorId: macro indicator ID
func (sdk *SDK) GetMacroIndicatorHistory(indicatorId uint64, opts ...adapt.Option) (*qotgetmacroindicatorhistory.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetMacroIndicatorHistoryWithContext(ctx, indicatorId, opts...)
}

// GetFedWatchTargetRate 3404 - gets the fed watch target rate.
func (sdk *SDK) GetFedWatchTargetRate(opts ...adapt.Option) (*qotgetfedwatchtargetrate.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFedWatchTargetRateWithContext(ctx, opts...)
}

// GetFedWatchDotPlot 3405 - gets the fed watch dot plot.
func (sdk *SDK) GetFedWatchDotPlot(opts ...adapt.Option) (*qotgetfedwatchdotplot.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFedWatchDotPlotWithContext(ctx, opts...)
}

// GetEarningsBeatRank 3406 - gets the earnings beat rank.
//
// market: market, see adapt.QotMarket_*
//
// beatType: earnings beat type
func (sdk *SDK) GetEarningsBeatRank(market int32, beatType int32, opts ...adapt.Option) (*qotgetearningsbeatrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEarningsBeatRankWithContext(ctx, market, beatType, opts...)
}

// GetDividendRank 3407 - gets the dividend rank.
//
// market: market, see adapt.QotMarket_*
//
// rankType: dividend rank type
func (sdk *SDK) GetDividendRank(market int32, rankType int32, opts ...adapt.Option) (*qotgetdividendrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetDividendRankWithContext(ctx, market, rankType, opts...)
}

// GetDividendCalendar 3408 - gets the dividend calendar.
//
// market: market, see adapt.QotMarket_*
//
// date: date, format: yyyy-MM-dd
func (sdk *SDK) GetDividendCalendar(market int32, date string, opts ...adapt.Option) (*qotgetdividendcalendar.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetDividendCalendarWithContext(ctx, market, date, opts...)
}

// GetEconomicCalendar 3409 - gets the economic calendar.
//
// beginDate: begin date, format: yyyy-MM-dd
func (sdk *SDK) GetEconomicCalendar(beginDate string, opts ...adapt.Option) (*qotgeteconomiccalendar.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEconomicCalendarWithContext(ctx, beginDate, opts...)
}

// GetUSPreMarketRank 3410 - gets the US pre market rank.
func (sdk *SDK) GetUSPreMarketRank(opts ...adapt.Option) (*qotgetuspremarketrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetUSPreMarketRankWithContext(ctx, opts...)
}

// GetUSAfterHoursRank 3411 - gets the US after hours rank.
func (sdk *SDK) GetUSAfterHoursRank(opts ...adapt.Option) (*qotgetusafterhoursrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetUSAfterHoursRankWithContext(ctx, opts...)
}

// GetUSOvernightRank 3412 - gets the US overnight rank.
func (sdk *SDK) GetUSOvernightRank(opts ...adapt.Option) (*qotgetusovernightrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetUSOvernightRankWithContext(ctx, opts...)
}

// GetTopMoversRank 3413 - gets the top movers rank.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetTopMoversRank(market int32, opts ...adapt.Option) (*qotgettopmoverrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetTopMoversRankWithContext(ctx, market, opts...)
}

// GetHotList 3414 - gets the hot list.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetHotList(market int32, opts ...adapt.Option) (*qotgethotlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetHotListWithContext(ctx, market, opts...)
}

// GetShortSellingRank 3415 - gets the short selling rank.
func (sdk *SDK) GetShortSellingRank(opts ...adapt.Option) (*qotgetshortsellingrank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetShortSellingRankWithContext(ctx, opts...)
}

// GetPeriodChangeRank 3416 - gets the period change rank.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetPeriodChangeRank(market int32, opts ...adapt.Option) (*qotgetperiodchangerank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetPeriodChangeRankWithContext(ctx, market, opts...)
}

// GetHighDividendSOERank 3417 - gets the high dividend SOE rank.
func (sdk *SDK) GetHighDividendSOERank(opts ...adapt.Option) (*qotgethighdividendsoerank.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetHighDividendSOERankWithContext(ctx, opts...)
}

// GetInstitutionList 3418 - gets the institution list.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetInstitutionList(market int32, opts ...adapt.Option) (*qotgetinstitutionlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInstitutionListWithContext(ctx, market, opts...)
}

// GetInstitutionProfile 3419 - gets the institution profile.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionProfile(market int32, institutionId int32) (*qotgetinstitutionprofile.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInstitutionProfileWithContext(ctx, market, institutionId)
}

// GetInstitutionDistribution 3420 - gets the institution distribution.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionDistribution(market int32, institutionId int32) (*qotgetinstitutiondistribution.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInstitutionDistributionWithContext(ctx, market, institutionId)
}

// GetInstitutionHoldingChange 3421 - gets the institution holding change.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionHoldingChange(market int32, institutionId int32, opts ...adapt.Option) (*qotgetinstitutionholdingchange.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInstitutionHoldingChangeWithContext(ctx, market, institutionId, opts...)
}

// GetInstitutionHoldingList 3422 - gets the institution holding list.
//
// market: market, see adapt.QotMarket_*
//
// institutionId: institution ID
func (sdk *SDK) GetInstitutionHoldingList(market int32, institutionId int32, opts ...adapt.Option) (*qotgetinstitutionholdinglist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetInstitutionHoldingListWithContext(ctx, market, institutionId, opts...)
}

// GetArkFundHolding 3423 - gets the ARK fund holding.
func (sdk *SDK) GetArkFundHolding(opts ...adapt.Option) (*qotgetarkfundholding.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetArkFundHoldingWithContext(ctx, opts...)
}

// GetArkStockDynamic 3424 - gets the ARK stock dynamic.
//
// code: security code, e.g. HK.00700
func (sdk *SDK) GetArkStockDynamic(code string) (*qotgetarkstockdynamic.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetArkStockDynamicWithContext(ctx, code)
}

// GetArkActiveTransaction 3425 - gets the ARK active transaction.
func (sdk *SDK) GetArkActiveTransaction(opts ...adapt.Option) (*qotgetarkactivetransaction.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetArkActiveTransactionWithContext(ctx, opts...)
}

// GetRatingChange 3426 - gets the rating change.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetRatingChange(market int32, opts ...adapt.Option) (*qotgetratingchange.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetRatingChangeWithContext(ctx, market, opts...)
}

// GetIndustrialChainList 3427 - gets the industrial chain list.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetIndustrialChainList(market int32, opts ...adapt.Option) (*qotgetindustrialchainlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIndustrialChainListWithContext(ctx, market, opts...)
}

// GetIndustrialChainDetail 3428 - gets the industrial chain detail.
//
// chainId: industrial chain ID
func (sdk *SDK) GetIndustrialChainDetail(chainId int64) (*qotgetindustrialchaindetail.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIndustrialChainDetailWithContext(ctx, chainId)
}

// GetIndustrialChainByPlate 3429 - gets the industrial chain by plate.
//
// plateId: industrial plate ID
func (sdk *SDK) GetIndustrialChainByPlate(plateId int64) (*qotgetindustrialchainbyplate.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIndustrialChainByPlateWithContext(ctx, plateId)
}

// GetIndustrialPlateInfo 3430 - gets the industrial plate info.
//
// plateId: industrial plate ID
func (sdk *SDK) GetIndustrialPlateInfo(plateId int64) (*qotgetindustrialplateinfo.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIndustrialPlateInfoWithContext(ctx, plateId)
}

// GetIndustrialPlateStock 3431 - gets the industrial plate stock.
func (sdk *SDK) GetIndustrialPlateStock(opts ...adapt.Option) (*qotgetindustrialplatestock.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetIndustrialPlateStockWithContext(ctx, opts...)
}

// GetHeatMapData 3432 - gets the heat map data.
//
// market: market, see adapt.QotMarket_*
func (sdk *SDK) GetHeatMapData(market int32, opts ...adapt.Option) (*qotgetheatmapdata.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetHeatMapDataWithContext(ctx, market, opts...)
}

// GetRiseFallDistribution 3433 - gets the rise fall distribution.
func (sdk *SDK) GetRiseFallDistribution(opts ...adapt.Option) (*qotgetrisefalldistr.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetRiseFallDistributionWithContext(ctx, opts...)
}

// GetEventContractCategory 3434 - gets the event contract category.
func (sdk *SDK) GetEventContractCategory(opts ...adapt.Option) (*qotgeteventcontractcategory.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractCategoryWithContext(ctx, opts...)
}

// FilterCompetition 3435 - filters the competition.
func (sdk *SDK) FilterCompetition(opts ...adapt.Option) (*qotfiltercompetition.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.FilterCompetitionWithContext(ctx, opts...)
}

// GetEventContractSeriesList 3436 - gets the event contract series list.
func (sdk *SDK) GetEventContractSeriesList(opts ...adapt.Option) (*qotgeteventcontractserieslist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractSeriesListWithContext(ctx, opts...)
}

// GetEventContractEventList 3437 - gets the event contract event list.
//
// code: series code of the event contract
func (sdk *SDK) GetEventContractEventList(code string, opts ...adapt.Option) (*qotgeteventcontracteventlist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractEventListWithContext(ctx, code, opts...)
}

// GetEventContract 3438 - gets the event contract.
//
// code: event code of the event contract
func (sdk *SDK) GetEventContract(code string, opts ...adapt.Option) (*qotgeteventcontract.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractWithContext(ctx, code, opts...)
}

// GetEventContractMilestoneList 3439 - gets the event contract milestone list.
func (sdk *SDK) GetEventContractMilestoneList(opts ...adapt.Option) (*qotgeteventcontractmilestonelist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractMilestoneListWithContext(ctx, opts...)
}

// GetEventContractSnapshot 3445 - gets the event contract snapshot.
func (sdk *SDK) GetEventContractSnapshot(opts ...adapt.Option) (*qotgeteventcontractsnapshot.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractSnapshotWithContext(ctx, opts...)
}

// GetEventContractOrderBook 3446 - gets the event contract order book.
//
// code: security code, e.g. HK.00700
//
// num: number of the order book levels
func (sdk *SDK) GetEventContractOrderBook(code string, num int32) (*qotgeteventcontractorderbook.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractOrderBookWithContext(ctx, code, num)
}

// GetEventContractKline 3447 - gets the event contract K line.
//
// code: security code, e.g. HK.00700
func (sdk *SDK) GetEventContractKline(code string, opts ...adapt.Option) (*qotgeteventcontractkline.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractKlineWithContext(ctx, code, opts...)
}

// GetEventContractTicker 3448 - gets the event contract ticker.
//
// code: security code, e.g. HK.00700
func (sdk *SDK) GetEventContractTicker(code string, opts ...adapt.Option) (*qotgeteventcontractticker.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractTickerWithContext(ctx, code, opts...)
}

// GetEventContractComboList 3453 - gets the event contract combo list.
func (sdk *SDK) GetEventContractComboList(opts ...adapt.Option) (*qotgeteventcontractcombolist.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractComboListWithContext(ctx, opts...)
}

// GetEventContractComboRfq 3454 - gets the event contract combo RFQ.
//
// mvc: combo identifier returned by the valid combo list
func (sdk *SDK) GetEventContractComboRfq(mvc string, opts ...adapt.Option) (*qotgeteventcontractcomborfq.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetEventContractComboRfqWithContext(ctx, mvc, opts...)
}

// SubEventContract 3455 - subscribes or unsubscribes event contract.
//
// isSubOrUnSub: true to subscribe, false to unsubscribe
func (sdk *SDK) SubEventContract(isSubOrUnSub bool, opts ...adapt.Option) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.SubEventContractWithContext(ctx, isSubOrUnSub, opts...)
}

// RequestHistoryEventContractKL 3456 - requests the history K line of an event contract.
//
// code: security code, e.g. HK.00700
//
// klType: K line type, see adapt.KLType_*
//
// beginTime: begin time, format: yyyy-MM-dd
//
// endTime: end time, format: yyyy-MM-dd
func (sdk *SDK) RequestHistoryEventContractKL(code string, klType int32, beginTime string, endTime string, opts ...adapt.Option) (*qotrequesthistoryeventcontractkl.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestHistoryEventContractKLWithContext(ctx, code, klType, beginTime, endTime, opts...)
}
