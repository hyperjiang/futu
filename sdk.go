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
	"github.com/hyperjiang/futu/pb/qotgetfutureinfo"
	"github.com/hyperjiang/futu/pb/qotgetipolist"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetmarketstate"
	"github.com/hyperjiang/futu/pb/qotgetoptionchain"
	"github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetownerplate"
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
)

const defaultTimeout = time.Second * 5

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
func (sdk *SDK) GetFunds(header *trdcommon.TrdHeader, opts ...adapt.Option) (*trdgetfunds.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetFundsWithContext(ctx, header, opts...)
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
