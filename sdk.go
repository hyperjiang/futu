package futu

import (
	"context"
	"time"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbroker"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
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
// begin: begin time, format: "yyyy-MM-dd"
//
// end: end time, format: "yyyy-MM-dd"
func (sdk *SDK) RequestHistoryKL(code string, klType int32, begin string, end string, opts ...adapt.Option) (*qotrequesthistorykl.S2C, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.RequestHistoryKLWithContext(ctx, code, klType, begin, end, opts...)
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
