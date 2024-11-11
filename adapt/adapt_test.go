package adapt

import (
	"testing"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
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

func TestSecurityToCode(t *testing.T) {
	should := require.New(t)

	s := &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		Code:   proto.String("00700"),
	}
	should.Equal("HK.00700", SecurityToCode(s))
}

func TestNewTradeAccount(t *testing.T) {
	should := require.New(t)

	ta := NewTradeAccount(123, TrdMarket_HK)
	should.NotNil(ta)
	should.Equal(int32(1), ta.GetTrdMarket())
	should.Equal(uint64(123), ta.GetAccID())
	should.Equal(int32(1), ta.GetTrdEnv())
}

func TestNewTestingTradeAccount(t *testing.T) {
	should := require.New(t)

	ta := NewTestingTradeAccount(123, TrdMarket_HK)
	should.NotNil(ta)
	should.Equal(int32(1), ta.GetTrdMarket())
	should.Equal(uint64(123), ta.GetAccID())
	should.Equal(int32(0), ta.GetTrdEnv())
}

func TestNewBaseFilter(t *testing.T) {
	should := require.New(t)

	f := NewBaseFilter(
		qotstockfilter.StockField_StockField_MarketVal,
		10000000000,
		20000000000,
		qotstockfilter.SortDir_SortDir_Ascend,
	)
	should.NotNil(f)
	should.Equal(int32(qotstockfilter.StockField_StockField_MarketVal), f.GetFieldName())
	should.Equal(float64(10000000000), f.GetFilterMin())
	should.Equal(float64(20000000000), f.GetFilterMax())
	should.Equal(int32(qotstockfilter.SortDir_SortDir_Ascend), f.GetSortDir())
	should.False(f.GetIsNoFilter())
}

func TestNewAccumulateFilter(t *testing.T) {
	should := require.New(t)

	f := NewAccumulateFilter(
		qotstockfilter.AccumulateField_AccumulateField_TurnoverRate,
		10,
		50,
		5,
		qotstockfilter.SortDir_SortDir_Ascend,
	)
	should.NotNil(f)
	should.Equal(int32(qotstockfilter.AccumulateField_AccumulateField_TurnoverRate), f.GetFieldName())
	should.Equal(float64(10), f.GetFilterMin())
	should.Equal(float64(50), f.GetFilterMax())
	should.Equal(int32(5), f.GetDays())
	should.Equal(int32(qotstockfilter.SortDir_SortDir_Ascend), f.GetSortDir())
	should.False(f.GetIsNoFilter())
}

func TestNewFinancialFilter(t *testing.T) {
	should := require.New(t)

	f := NewFinancialFilter(
		qotstockfilter.FinancialField_FinancialField_NetProfit,
		10000000000,
		0,
		3,
		qotstockfilter.SortDir_SortDir_Ascend,
	)
	should.NotNil(f)
	should.Equal(int32(qotstockfilter.FinancialField_FinancialField_NetProfit), f.GetFieldName())
	should.Equal(float64(10000000000), f.GetFilterMin())
	should.Equal(float64(0), f.GetFilterMax())
	should.Equal(int32(3), f.GetQuarter())
	should.Equal(int32(qotstockfilter.SortDir_SortDir_Ascend), f.GetSortDir())
	should.False(f.GetIsNoFilter())
}

func TestNewPatternFilter(t *testing.T) {
	should := require.New(t)

	f := NewPatternFilter(
		qotstockfilter.PatternField_PatternField_MAAlignmentLong,
		qotcommon.KLType_KLType_Day,
		3,
	)
	should.NotNil(f)
	should.Equal(int32(qotstockfilter.PatternField_PatternField_MAAlignmentLong), f.GetFieldName())
	should.Equal(int32(qotcommon.KLType_KLType_Day), f.GetKlType())
	should.Equal(int32(3), f.GetConsecutivePeriod())
	should.False(f.GetIsNoFilter())
}

func TestNewCustomIndicatorFilter(t *testing.T) {
	should := require.New(t)

	f := NewCustomIndicatorFilter(
		With("firstFieldName", qotstockfilter.CustomIndicatorField_CustomIndicatorField_Price),
		With("secondFieldName", qotstockfilter.CustomIndicatorField_CustomIndicatorField_MA20),
		With("relativePosition", qotstockfilter.RelativePosition_RelativePosition_More),
		With("fieldValue", 10),
		With("klType", qotcommon.KLType_KLType_Day),
		With("firstFieldParaList", []int32{1, 4}),
		With("secondFieldParaList", []int32{1, 4}),
		With("consecutivePeriod", 3),
	)
	should.NotNil(f)
	should.Equal(int32(qotstockfilter.CustomIndicatorField_CustomIndicatorField_Price), f.GetFirstFieldName())
	should.Equal(int32(qotstockfilter.CustomIndicatorField_CustomIndicatorField_MA20), f.GetSecondFieldName())
	should.Equal(int32(qotstockfilter.RelativePosition_RelativePosition_More), f.GetRelativePosition())
	should.Equal(float64(10), f.GetFieldValue())
	should.Equal(int32(qotcommon.KLType_KLType_Day), f.GetKlType())
	should.Equal([]int32{1, 4}, f.GetFirstFieldParaList())
	should.Equal([]int32{1, 4}, f.GetSecondFieldParaList())
	should.Equal(int32(3), f.GetConsecutivePeriod())
	should.False(f.GetIsNoFilter())
}
