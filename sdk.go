package futu

import (
	"context"
	"time"

	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/getglobalstate"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
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
