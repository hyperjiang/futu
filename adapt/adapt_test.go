package adapt

import (
	"testing"

	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/stretchr/testify/require"
)

func TestNewSecurity(t *testing.T) {
	should := require.New(t)

	s := NewSecurity("HK.00700")
	should.NotNil(s)
	should.Equal(int32(1), s.GetMarket())
	should.Equal("00700", s.GetCode())

	should.Nil(NewSecurity("00700"))
}

func TestNewSecurities(t *testing.T) {
	should := require.New(t)

	sa := NewSecurities([]string{"HK.00700", "US.AAPL"})
	should.Len(sa, 2)
	should.Equal(int32(1), sa[0].GetMarket())
	should.Equal("00700", sa[0].GetCode())
	should.Equal(int32(11), sa[1].GetMarket())
	should.Equal("AAPL", sa[1].GetCode())

	sa = NewSecurities([]string{"HK.00700", "00700"})
	should.Len(sa, 1)
	should.Equal(int32(1), sa[0].GetMarket())
	should.Equal("00700", sa[0].GetCode())
}

func TestNewTradeAccount(t *testing.T) {
	should := require.New(t)

	ta := NewTradeAccount(123, 1)
	should.NotNil(ta)
	should.Equal(int32(1), ta.GetTrdMarket())
	should.Equal(uint64(123), ta.GetAccID())
	should.Equal(int32(1), ta.GetTrdEnv())
}

func TestNewTestingTradeAccount(t *testing.T) {
	should := require.New(t)

	ta := NewTestingTradeAccount(123, trdcommon.TrdMarket_TrdMarket_HK)
	should.NotNil(ta)
	should.Equal(int32(1), ta.GetTrdMarket())
	should.Equal(uint64(123), ta.GetAccID())
	should.Equal(int32(0), ta.GetTrdEnv())
}
