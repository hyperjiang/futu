package client_test

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (ts *ClientTestSuite) TestQotGetBasicQot() {
	should := require.New(ts.T())

	c2s := &qotgetbasicqot.C2S{
		SecurityList: []*qotcommon.Security{alibaba},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	s2c, err := ts.client.QotGetBasicQot(ctx, c2s)
	should.NoError(err)
	fmt.Println(s2c.GetBasicQotList())
}

func (ts *ClientTestSuite) TestQotGetSubInfo() {
	should := require.New(ts.T())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	info, err := ts.client.QotGetSubInfo(ctx, &qotgetsubinfo.C2S{})
	should.NoError(err)
	fmt.Println(info)
}

func (ts *ClientTestSuite) TestQotGetKL() {
	should := require.New(ts.T())

	c2s := &qotgetkl.C2S{
		RehabType: proto.Int32(int32(qotcommon.RehabType_RehabType_Forward)),
		KlType:    proto.Int32(int32(qotcommon.KLType_KLType_Day)),
		Security:  tencent,
		ReqNum:    proto.Int32(3),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s2c, err := ts.client.QotGetKL(ctx, c2s)
	should.NoError(err)
	fmt.Println(s2c.GetKlList())
}

func (ts *ClientTestSuite) TestQotRequestHistoryKL() {
	should := require.New(ts.T())

	beginDate := "2024-10-01"
	endDate := "2024-10-10"

	c2s := &qotrequesthistorykl.C2S{
		RehabType:   proto.Int32(int32(qotcommon.RehabType_RehabType_Forward)),
		KlType:      proto.Int32(int32(qotcommon.KLType_KLType_Day)),
		Security:    tencent,
		BeginTime:   proto.String(beginDate),
		EndTime:     proto.String(endDate),
		MaxAckKLNum: proto.Int32(3),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotRequestHistoryKL(ctx, c2s)
	should.NoError(err)
	log.Info().Str("name", res.GetName()).Msg("QotRequestHistoryKL")

	for _, kl := range res.GetKlList() {
		log.Info().Str("date", kl.GetTime()).Float64("close", kl.GetClosePrice()).Msg("")
	}

	next := res.GetNextReqKey()
	for len(next) > 0 {
		c2s.NextReqKey = next
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err = ts.client.QotRequestHistoryKL(ctx, c2s)
		should.NoError(err)

		for _, kl := range res.GetKlList() {
			log.Info().Str("date", kl.GetTime()).Float64("close", kl.GetClosePrice()).Msg("")
		}

		next = res.GetNextReqKey()
	}
}

func (ts *ClientTestSuite) TestQotStockFilter() {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotStockFilter(ctx, c2s)
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

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		snapshot, err := ts.client.QotGetSecuritySnapshot(ctx, snapshotC2S)
		should.NoError(err)
		fmt.Println(snapshot.GetSnapshotList()[0])
	}
}
