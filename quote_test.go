package futu_test

import (
	"fmt"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (ts *FutuTestSuite) TestQotRequestHistoryKL() {
	should := require.New(ts.T())

	beginDate := "2024-10-01"
	endDate := "2024-10-10"
	security := &qotcommon.Security{
		Market: (*int32)(qotcommon.QotMarket_QotMarket_HK_Security.Enum()),
		Code:   proto.String("00700"),
	}

	c2s := &qotrequesthistorykl.C2S{
		RehabType:   proto.Int32(int32(qotcommon.RehabType_RehabType_Forward)),
		KlType:      proto.Int32(int32(qotcommon.KLType_KLType_Day)),
		Security:    security,
		BeginTime:   proto.String(beginDate),
		EndTime:     proto.String(endDate),
		MaxAckKLNum: proto.Int32(3),
	}

	res, err := ts.client.QotRequestHistoryKL(c2s)
	should.NoError(err)
	log.Info().Str("name", res.GetName()).Msg("QotRequestHistoryKL")

	for _, kl := range res.GetKlList() {
		log.Info().Str("date", kl.GetTime()).Float64("close", kl.GetClosePrice()).Msg("")
	}

	next := res.GetNextReqKey()
	for len(next) > 0 {
		c2s.NextReqKey = next
		res, err = ts.client.QotRequestHistoryKL(c2s)
		should.NoError(err)

		for _, kl := range res.GetKlList() {
			log.Info().Str("date", kl.GetTime()).Float64("close", kl.GetClosePrice()).Msg("")
		}

		next = res.GetNextReqKey()
	}
}

func (ts *FutuTestSuite) TestQotStockFilter() {
	should := require.New(ts.T())

	var begin int32 = 0
	var num int32 = 10
	f := &qotstockfilter.BaseFilter{
		FieldName:  (*int32)(qotstockfilter.StockField_StockField_MarketVal.Enum()),
		FilterMin:  proto.Float64(10000000000),
		SortDir:    (*int32)(qotstockfilter.SortDir_SortDir_Ascend.Enum()),
		IsNoFilter: proto.Bool(false),
	}

	c2s := &qotstockfilter.C2S{
		Begin:  &begin,
		Num:    &num,
		Market: (*int32)(qotcommon.QotMarket_QotMarket_HK_Security.Enum()),
		// Plate:          &qotcommon.Security{Code: proto.String("Motherboard"), Market: (*int32)(qotcommon.QotMarket_QotMarket_HK_Security.Enum())},
		BaseFilterList: []*qotstockfilter.BaseFilter{f},
	}

	res, err := ts.client.QotStockFilter(c2s)
	should.NoError(err)

	fmt.Println(res.GetAllCount())

	for _, stock := range res.GetDataList() {
		log.Info().Str("code", stock.GetSecurity().GetCode()).
			Str("name", stock.GetName()).
			Int32("market", stock.GetSecurity().GetMarket()).
			Msg("")

		snapshotC2S := &qotgetsecuritysnapshot.C2S{
			SecurityList: []*qotcommon.Security{stock.GetSecurity()},
		}

		snapshot, err := ts.client.QotGetSecuritySnapshot(snapshotC2S)
		should.NoError(err)
		fmt.Println(snapshot.GetSnapshotList()[0])
	}
}
