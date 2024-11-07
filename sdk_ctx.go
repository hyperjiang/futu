package futu

import (
	"context"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetbroker"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetplateset"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetstaticinfo"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotsub"
	"google.golang.org/protobuf/proto"
)

// GetGlobalStateWithContext 1002 - gets the global state with context.
func (sdk *SDK) GetGlobalStateWithContext(ctx context.Context) (*getglobalstate.S2C, error) {
	return sdk.cli.GetGlobalState(ctx)
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
// begin: begin time, format: "yyyy-MM-dd"
//
// end: end time, format: "yyyy-MM-dd"
func (sdk *SDK) RequestHistoryKLWithContext(ctx context.Context, code string, klType int32, begin string, end string, opts ...adapt.Option) (*qotrequesthistorykl.S2C, error) {
	o := adapt.NewOptions(opts...)
	o["security"] = adapt.NewSecurity(code)
	o["klType"] = klType
	o["beginTime"] = begin
	o["endTime"] = end

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