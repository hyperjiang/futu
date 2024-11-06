package futu

import (
	"context"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotsub"
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
