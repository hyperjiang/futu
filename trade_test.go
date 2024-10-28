package futu_test

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
	"github.com/hyperjiang/futu/pb/trdgetorderlist"
	"github.com/hyperjiang/futu/pb/trdmodifyorder"
	"github.com/hyperjiang/futu/pb/trdplaceorder"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (ts *FutuTestSuite) TestTrdGetAccList_TrdGetFunds() {
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
		fmt.Println(res.GetFunds())
	}
}

func (ts *FutuTestSuite) TestTrdPlaceOrder() {
	should := require.New(ts.T())

	c2s := &trdplaceorder.C2S{
		Header:  usAccount,
		TrdSide: proto.Int32(int32(trdcommon.TrdSide_TrdSide_Buy)),
		// OrderType: proto.Int32(int32(trdcommon.OrderType_OrderType_Market)),
		OrderType: proto.Int32(int32(trdcommon.OrderType_OrderType_Normal)),
		Code:      proto.String("AAPL"),
		Price:     proto.Float64(234),
		Qty:       proto.Float64(1),
		SecMarket: proto.Int32(int32(trdcommon.TrdSecMarket_TrdSecMarket_US)),
		Remark:    proto.String("test via futu go sdk"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdPlaceOrder(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("order", res).Msg("place order")

	// cancel the order
	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c2s2 := &trdmodifyorder.C2S{
		Header:        usAccount,
		OrderID:       proto.Uint64(res.GetOrderID()),
		ModifyOrderOp: proto.Int32(int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Cancel)),
	}

	res2, err := ts.client.TrdModifyOrder(ctx2, c2s2)
	should.NoError(err)
	log.Info().Interface("order", res2).Msg("cancel order")
}

func (ts *FutuTestSuite) TestTrdGetOrderList() {
	should := require.New(ts.T())

	c2s := &trdgetorderlist.C2S{
		Header: usAccount,
		// FilterConditions: &trdcommon.TrdFilterConditions{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := ts.client.TrdGetOrderList(ctx, c2s)
	should.NoError(err)
	for _, order := range res.GetOrderList() {
		fmt.Printf("order: %+v\n", order)
	}
}
