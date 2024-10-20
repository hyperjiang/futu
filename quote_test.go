package futu_test

import (
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
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
