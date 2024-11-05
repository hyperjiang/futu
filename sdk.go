package futu

import (
	"context"
	"time"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotsub"
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

// GetGlobalStateWithContext 1002 - gets the global state with context.
func (sdk *SDK) GetGlobalStateWithContext(ctx context.Context) (*getglobalstate.S2C, error) {
	return sdk.cli.GetGlobalState(ctx)
}

// Subscribe 3001 - subscribes or unsubscribes.
func (sdk *SDK) Subscribe(opts ...adapt.Option) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.SubscribeWithContext(ctx, opts...)
}

// SubscribeWithContext 3001 - subscribes or unsubscribes with context.
func (sdk *SDK) SubscribeWithContext(ctx context.Context, opts ...adapt.Option) error {
	o := adapt.NewOptions(opts...)

	var c2s qotsub.C2S
	if err := o.ToProto(&c2s); err != nil {
		return err
	}

	return sdk.cli.QotSub(ctx, &c2s)
}

// GetBasicQot 3004 - gets the basic quotes of given securities.
func (sdk *SDK) GetBasicQot(codes []string) ([]*qotcommon.BasicQot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	return sdk.GetBasicQotWithContext(ctx, codes)
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
