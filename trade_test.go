package futu_test

import (
	"fmt"

	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/pb/trdgetacclist"
	"github.com/hyperjiang/futu/pb/trdgetfunds"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (ts *FutuTestSuite) TestTrdGetAccList_TrdGetFunds() {
	should := require.New(ts.T())

	c2s := &trdgetacclist.C2S{
		TrdCategory: proto.Int32(int32(trdcommon.TrdCategory_TrdCategory_Security)),
	}

	res, err := ts.client.TrdGetAccList(c2s)
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

		res, err := ts.client.TrdGetFunds(c2s)
		should.NoError(err)
		fmt.Println(res.GetFunds())
	}
}
