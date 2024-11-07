package adapt

import (
	"testing"

	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/stretchr/testify/require"
)

func TestEmptyOption(t *testing.T) {
	should := require.New(t)

	opts := NewOptions()
	var ta trdcommon.TrdHeader
	err := opts.ToProto(&ta)
	should.NoError(err)
	should.Equal(int32(0), ta.GetTrdMarket())
	should.Equal(uint64(0), ta.GetAccID())
	should.Equal(int32(0), ta.GetTrdEnv())
}

func TestSimpleOption(t *testing.T) {
	should := require.New(t)

	opts := NewOptions(
		With("trdEnv", 1),
		With("accID", 123),
		With("trdMarket", 1),
		WithSecurity("HK.00700"), // this should be ignored
	)

	var ta trdcommon.TrdHeader
	err := opts.ToProto(&ta)
	should.NoError(err)
	should.Equal(int32(1), ta.GetTrdMarket())
	should.Equal(uint64(123), ta.GetAccID())
	should.Equal(int32(1), ta.GetTrdEnv())
}

func TestInvalidOption(t *testing.T) {
	should := require.New(t)

	opts := NewOptions(
		With("trdEnv", 1),
		With("func", func() {}),
	)
	var ta trdcommon.TrdHeader
	err := opts.ToProto(&ta)
	should.Error(err)
}

func TestComplexOption(t *testing.T) {
	should := require.New(t)

	opts := NewOptions(
		WithSecurities([]string{"HK.00700", "US.AAPL"}),
	)

	var c2s qotgetbasicqot.C2S
	err := opts.ToProto(&c2s)
	should.NoError(err)
	should.Len(c2s.GetSecurityList(), 2)
	should.Equal(int32(1), c2s.GetSecurityList()[0].GetMarket())
	should.Equal("00700", c2s.GetSecurityList()[0].GetCode())
	should.Equal(int32(11), c2s.GetSecurityList()[1].GetMarket())
	should.Equal("AAPL", c2s.GetSecurityList()[1].GetCode())
}

func TestComplexOption2(t *testing.T) {
	should := require.New(t)

	opts := NewOptions(
		WithSecurities([]string{"HK.00700", "US.AAPL"}),
		With("subTypeList", []int32{1, 2, 3}),
		With("isSubOrUnSub", true),
	)

	var c2s qotsub.C2S
	err := opts.ToProto(&c2s)
	should.NoError(err)
	should.Len(c2s.GetSecurityList(), 2)
	should.Equal(int32(1), c2s.GetSecurityList()[0].GetMarket())
	should.Equal("00700", c2s.GetSecurityList()[0].GetCode())
	should.Equal(int32(11), c2s.GetSecurityList()[1].GetMarket())
	should.Equal("AAPL", c2s.GetSecurityList()[1].GetCode())
	should.Len(c2s.GetSubTypeList(), 3)
	should.True(c2s.GetIsSubOrUnSub())
	should.False(c2s.GetIsRegOrUnRegPush())
	should.False(c2s.GetIsFirstPush())
	should.Empty(c2s.GetRegPushRehabTypeList())
}
