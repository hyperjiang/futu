package client_test

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
	"github.com/hyperjiang/futu/pb/trdgethistoryorderfilllist"
	"github.com/hyperjiang/futu/pb/trdgethistoryorderlist"
	"github.com/hyperjiang/futu/pb/trdgetmarginratio"
	"github.com/hyperjiang/futu/pb/trdgetmaxtrdqtys"
	"github.com/hyperjiang/futu/pb/trdgetorderfee"
	"github.com/hyperjiang/futu/pb/trdgetorderfilllist"
	"github.com/hyperjiang/futu/pb/trdgetorderlist"
	"github.com/hyperjiang/futu/pb/trdgetpositionlist"
	"github.com/hyperjiang/futu/pb/trdmodifyorder"
	"github.com/hyperjiang/futu/pb/trdplaceorder"
	"github.com/hyperjiang/futu/pb/trdsubaccpush"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (ts *ClientTestSuite) TestTrdGetAccList_TrdGetFunds() {
	should := require.New(ts.T())

	c2s := &trdgetacclist.C2S{
		TrdCategory: proto.Int32(int32(trdcommon.TrdCategory_TrdCategory_Security)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.TrdGetAccList(ctx, c2s)
	should.NoError(err)

	for _, acc := range res.GetAccList() {
		fmt.Printf("acc: %+v\n", acc)
		if acc.GetTrdEnv() != int32(trdcommon.TrdEnv_TrdEnv_Simulate) {
			continue
		}

		header := &trdcommon.TrdHeader{
			TrdEnv:    proto.Int32(int32(acc.GetTrdEnv())), // trdcommon.TrdEnv_TrdEnv_Simulate
			AccID:     proto.Uint64(acc.GetAccID()),
			TrdMarket: proto.Int32(int32(acc.GetTrdMarketAuthList()[0])), // trdcommon.TrdMarket_TrdMarket_HK
		}

		c2s := &trdgetfunds.C2S{
			Header: header,
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := ts.client.TrdGetFunds(ctx, c2s)
		should.NoError(err)
		log.Info().Interface("data", res.GetFunds()).Msg("TrdGetFunds")
	}
}

func (ts *ClientTestSuite) TestTrdSubAccPush() {
	should := require.New(ts.T())

	c2s := &trdsubaccpush.C2S{
		AccIDList: []uint64{usAccount.GetAccID()},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := ts.client.TrdSubAccPush(ctx, c2s)
	should.NoError(err)
}

func (ts *ClientTestSuite) TestTrdGetPositionList() {
	should := require.New(ts.T())

	c2s := &trdgetpositionlist.C2S{
		Header: usAccount,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetPositionList(ctx, c2s)
	should.NoError(err)
	for _, position := range res.GetPositionList() {
		log.Info().Interface("position", position).Msg("TrdGetPositionList")
	}
}

func (ts *ClientTestSuite) TestTrdGetMaxTrdQtys() {
	should := require.New(ts.T())

	c2s := &trdgetmaxtrdqtys.C2S{
		Header:    usAccount,
		OrderType: proto.Int32(int32(trdcommon.OrderType_OrderType_Normal)),
		Code:      proto.String("AAPL"),
		Price:     proto.Float64(230),
		SecMarket: proto.Int32(int32(trdcommon.TrdSecMarket_TrdSecMarket_US)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetMaxTrdQtys(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetMaxTrdQtys()).Msg("TrdGetMaxTrdQtys")
}

func (ts *ClientTestSuite) TestTrdPlaceOrder() {
	should := require.New(ts.T())

	c2s := &trdplaceorder.C2S{
		Header:    usAccount,
		TrdSide:   proto.Int32(int32(trdcommon.TrdSide_TrdSide_Buy)),
		OrderType: proto.Int32(int32(trdcommon.OrderType_OrderType_Market)),
		// OrderType: proto.Int32(int32(trdcommon.OrderType_OrderType_Normal)),
		Code: proto.String("AAPL"),
		// Price:     proto.Float64(230),
		Qty:       proto.Float64(1),
		SecMarket: proto.Int32(int32(trdcommon.TrdSecMarket_TrdSecMarket_US)),
		Remark:    proto.String("test via futu go sdk"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdPlaceOrder(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("order", res).Msg("place order")

	if err == nil {
		// cancel the order
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		c2s := &trdmodifyorder.C2S{
			Header:        usAccount,
			OrderID:       proto.Uint64(res.GetOrderID()),
			ModifyOrderOp: proto.Int32(int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Cancel)),
		}

		res, err := ts.client.TrdModifyOrder(ctx, c2s)
		should.NoError(err)
		log.Info().Interface("order", res).Msg("cancel order")
	}
}

func (ts *ClientTestSuite) TestTrdGetOrderList() {
	should := require.New(ts.T())

	c2s := &trdgetorderlist.C2S{
		Header: usAccount,
		// FilterConditions: &trdcommon.TrdFilterConditions{},
		FilterStatusList: []int32{int32(trdcommon.OrderStatus_OrderStatus_Submitted)},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetOrderList(ctx, c2s)
	should.NoError(err)
	for _, order := range res.GetOrderList() {
		log.Info().Interface("pending order", order).Msg("TrdGetOrderList")
	}
}

func (ts *ClientTestSuite) TestTrdGetOrderFillList() {
	should := require.New(ts.T())

	c2s := &trdgetorderfilllist.C2S{
		Header: usAccount,
		// FilterConditions: &trdcommon.TrdFilterConditions{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetOrderFillList(ctx, c2s)
	should.Error(err) // 模拟交易不支持成交数据
	for _, order := range res.GetOrderFillList() {
		log.Info().Interface("completed order", order).Msg("TrdGetOrderFillList")
	}
}

func (ts *ClientTestSuite) TestTrdGetHistoryOrderList() {
	should := require.New(ts.T())

	c2s := &trdgethistoryorderlist.C2S{
		Header: usAccount,
		FilterConditions: &trdcommon.TrdFilterConditions{
			BeginTime: proto.String(time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")),
			EndTime:   proto.String(time.Now().Format("2006-01-02 15:04:05")),
		},
		FilterStatusList: []int32{int32(trdcommon.OrderStatus_OrderStatus_Filled_All)},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetHistoryOrderList(ctx, c2s)
	should.NoError(err)

	var orderIDs []string // 已成交的订单ID
	for _, order := range res.GetOrderList() {
		log.Info().Interface("history order", order).Msg("GetOrderList")
		orderIDs = append(orderIDs, order.GetOrderIDEx())
	}

	if len(orderIDs) > 0 {
		c2s := &trdgetorderfee.C2S{
			Header:        usAccount,
			OrderIdExList: orderIDs,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := ts.client.TrdGetOrderFee(ctx, c2s)
		should.Error(err) // 暂时不支持模拟交易
	}
}

func (ts *ClientTestSuite) TestTrdGetHistoryOrderFillList() {
	should := require.New(ts.T())

	c2s := &trdgethistoryorderfilllist.C2S{
		Header: usAccount,
		FilterConditions: &trdcommon.TrdFilterConditions{
			BeginTime: proto.String(time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")),
			EndTime:   proto.String(time.Now().Format("2006-01-02 15:04:05")),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := ts.client.TrdGetHistoryOrderFillList(ctx, c2s)
	should.Error(err) // 模拟交易不支持成交数据
}

func (ts *ClientTestSuite) TestTrdGetMarginRatio() {
	should := require.New(ts.T())

	c2s := &trdgetmarginratio.C2S{
		Header:       usAccount,
		SecurityList: []*qotcommon.Security{apple},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetMarginRatio(ctx, c2s)
	should.NoError(err)
	for _, item := range res.GetMarginRatioInfoList() {
		log.Info().Interface("margin ratio", item).Msg("TrdGetMarginRatio")
	}
}
