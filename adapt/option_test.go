package adapt

import (
	"testing"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdgetpositionlist"
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

func TestFilters(t *testing.T) {
	should := require.New(t)

	opts := NewOptions(
		WithBaseFilters(
			NewBaseFilter(qotstockfilter.StockField_StockField_MarketVal, 10000000000, 0, qotstockfilter.SortDir_SortDir_Ascend),
		),
		WithAccumulateFilters(
			NewAccumulateFilter(qotstockfilter.AccumulateField_AccumulateField_TurnoverRate, 10, 50, 5, qotstockfilter.SortDir_SortDir_Ascend),
		),
		WithFinancialFilters(
			NewFinancialFilter(qotstockfilter.FinancialField_FinancialField_NetProfit, 10000000000, 0, 3, qotstockfilter.SortDir_SortDir_Ascend),
		),
		WithPatternFilters(
			NewPatternFilter(qotstockfilter.PatternField_PatternField_MAAlignmentLong, qotcommon.KLType_KLType_Day, 3),
		),
		WithCustomIndicatorFilters(
			NewCustomIndicatorFilter(
				With("firstFieldName", qotstockfilter.CustomIndicatorField_CustomIndicatorField_Price),
				With("secondFieldName", qotstockfilter.CustomIndicatorField_CustomIndicatorField_MA20),
				With("relativePosition", qotstockfilter.RelativePosition_RelativePosition_More),
				With("fieldValue", 10),
				With("klType", qotcommon.KLType_KLType_Day),
				With("firstFieldParaList", []int32{1, 4}),
				With("secondFieldParaList", []int32{1, 4}),
				With("consecutivePeriod", 3),
			),
		),
	)

	var c2s qotstockfilter.C2S
	err := opts.ToProto(&c2s)
	should.NoError(err)
	should.Len(c2s.GetBaseFilterList(), 1)
	should.Len(c2s.GetAccumulateFilterList(), 1)
	should.Len(c2s.GetFinancialFilterList(), 1)
	should.Len(c2s.GetPatternFilterList(), 1)
	should.Len(c2s.GetCustomIndicatorFilterList(), 1)
}

func TestFilterConditions(t *testing.T) {
	should := require.New(t)

	fc := NewFilterConditions(
		With("codeList", []string{"00700", "AAPL"}),
		With("idList", []uint64{123, 456}),
		With("beginTime", "2021-01-01"),
		With("endTime", "2021-12-31"),
		With("orderIDExList", []string{"123", "456"}),
	)

	opts := NewOptions(
		WithFilterConditions(fc),
	)

	var c2s trdgetpositionlist.C2S
	err := opts.ToProto(&c2s)
	should.NoError(err)
	should.Len(c2s.GetFilterConditions().GetCodeList(), 2)
	should.Len(c2s.GetFilterConditions().GetIdList(), 2)
	should.Equal("2021-01-01", c2s.GetFilterConditions().GetBeginTime())
	should.Equal("2021-12-31", c2s.GetFilterConditions().GetEndTime())
	should.Len(c2s.GetFilterConditions().GetOrderIDExList(), 2)
}
