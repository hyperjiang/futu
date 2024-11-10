package futu_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hyperjiang/futu"
	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/notify"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/protoid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

var pubKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAMkZDkCQzZbu0CWzeqPP/18HanqPIG4oD3EbmbqfFVxcUwWqIfcHF7N0
D+ZZ+XUioVpy/w0rV+8NaTOytypfqKyaDJv+K7wh0W4ERkcirK5TAQrU3rgTWtjb
295ZSpwLtaMCrC+ux99m6ptDM/YVXI8uVtP8X2ygFv0mgq5//meLAgMBAAE=
-----END RSA PUBLIC KEY-----`)

var priKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDJGQ5AkM2W7tAls3qjz/9fB2p6jyBuKA9xG5m6nxVcXFMFqiH3
BxezdA/mWfl1IqFacv8NK1fvDWkzsrcqX6ismgyb/iu8IdFuBEZHIqyuUwEK1N64
E1rY29veWUqcC7WjAqwvrsffZuqbQzP2FVyPLlbT/F9soBb9JoKuf/5niwIDAQAB
AoGBAJvq9lbvLsfwn6grkVIDig+nE1K1OELQgrCC4t2ETK6Q0roYoD8E28aCnXVP
m4/LaulTMheG3KX3cvLnhQawpnjUxm/3NZlVPj6EEjYepVyEBMLV2gBUzulUdTeZ
HM6hEBB3YQ8BnkJG1ajbr2lmilLenOaGTj2q6rxFz1n5dlWhAkEA7QaW0h8YrS6F
6ZRHcTui13ScwFxKAxuuOg9mbV9Y2EegDpAvhRdhvbx1pNCiD9vy46s6yAFtzNtF
+PtqnNASGwJBANkyMLusENpxZ1gucYd/RDwT0a9XMn6BAOPBJxLlhoKj1fI2YMoy
QJBHAFhh7BIt+U4XomXkhwTOUp67HPgc11ECQQC5QqUvps6Kzgos/5C3mH03GhZK
49eVhlUvXEoawqOWqKUZvOjnhdcHjf4FzGxfKPM3r+ZJ3ZQMwnZ2nUw/NQJxAkAi
jKpV4CwaI3n1/AVRMXxwNhLf2nYMy4aRtDL7/YjlFRy+V8oTv+SnTrQOWx1LUwba
VkYeATk9GXjpCQi1qxjRAkEA2jPfclINKKKfVPjys7R6Juq9sBFqJSmhcFYae8Xd
ywQCvmZiU66RGeo6pCSwdH0h4NeQ8w48SjhmRqswNKKr8g==
-----END RSA PRIVATE KEY-----`)

type SDKTestSuite struct {
	suite.Suite
	sdk *futu.SDK
}

// TestSDKTestSuite runs the http client test suite
func TestSDKTestSuite(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	suite.Run(t, new(SDKTestSuite))
}

// SetupSuite run once at the very start of the testing suite, before any tests are run.
func (ts *SDKTestSuite) SetupSuite() {
	var err error
	ts.sdk, err = futu.NewSDK(
		client.WithPrivateKey(priKey),
		client.WithPublicKey(pubKey),
	)
	if err != nil {
		ts.T().SkipNow()
	}
	ts.sdk.RegisterHandler(protoid.Notify, func(s2c proto.Message) error {
		msg := s2c.(*notify.S2C)
		log.Info().Interface("s2c", msg).Msg("custom handler")
		return nil
	})

	err = ts.sdk.Subscribe(
		[]string{"HK.09988", "HK.00700"},
		[]int32{
			adapt.SubType_Basic,
			adapt.SubType_RT,
			adapt.SubType_KL_Day,
			adapt.SubType_KL_3Min,
			adapt.SubType_Ticker,
			adapt.SubType_OrderBook,
			adapt.SubType_Broker,
		},
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
}

// TearDownSuite run once at the very end of the testing suite, after all tests have been run.
func (ts *SDKTestSuite) TearDownSuite() {
	if ts.sdk != nil {
		ts.sdk.Close()
	}
}

func (ts *SDKTestSuite) TestGetGlobalState() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetGlobalState()
	should.NoError(err)

	fmt.Println(res)
}

func (ts *SDKTestSuite) TestGetSubInfo() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetSubInfo()
	should.NoError(err)
	log.Info().Interface("result", res).Msg("GetSubInfo")
}

func (ts *SDKTestSuite) TestGetBasicQot() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetBasicQot([]string{"HK.00700", "HK.09988"})
	should.NoError(err)
	for _, qot := range res {
		log.Info().Interface("qot", qot).Msg("GetBasicQot")
	}
}

func (ts *SDKTestSuite) TestGetKL() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetKL(
		"HK.09988",
		adapt.KLType_Day,
		adapt.With("rehabType", adapt.RehabType_Forward),
		adapt.With("reqNum", 3),
	)
	should.NoError(err)
	for _, kl := range res.GetKlList() {
		log.Info().Interface("kline", kl).Msg("GetKL")
	}
}

func (ts *SDKTestSuite) TestGetRT() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetRT("HK.09988")
	should.NoError(err)
	should.Equal("09988", res.GetSecurity().GetCode())
	log.Info().Str("stock", res.GetName()).Int("num", len(res.GetRtList())).Msg("GetRT")
}

func (ts *SDKTestSuite) TestGetTicker() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetTicker("HK.09988")
	should.NoError(err)
	should.Equal("09988", res.GetSecurity().GetCode())
	log.Info().Str("stock", res.GetName()).Int("num", len(res.GetTickerList())).Msg("GetTicker")
}

func (ts *SDKTestSuite) TestGetOrderBook() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetOrderBook("HK.09988")
	should.NoError(err)
	should.Equal("09988", res.GetSecurity().GetCode())
	log.Info().Str("stock", res.GetName()).
		Int("实时卖盘", len(res.GetOrderBookAskList())).
		Int("实时买盘", len(res.GetOrderBookBidList())).
		Msg("GetOrderBook")
	for _, ask := range res.GetOrderBookAskList() {
		log.Info().Interface("data", ask).Msg("实时卖盘")
	}
	for _, bid := range res.GetOrderBookBidList() {
		log.Info().Interface("data", bid).Msg("实时买盘")
	}
}

func (ts *SDKTestSuite) TestGetBroker() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetBroker("HK.09988")
	should.NoError(err)
	should.Equal("09988", res.GetSecurity().GetCode())

	log.Info().Str("stock", res.GetName()).
		Int("实时经纪卖盘", len(res.GetBrokerAskList())).
		Int("实时经纪买盘", len(res.GetBrokerBidList())).
		Msg("GetBroker")
	for _, ask := range res.GetBrokerAskList() {
		log.Info().Interface("data", ask).Msg("实时经纪卖盘")
	}
	for _, bid := range res.GetBrokerBidList() {
		log.Info().Interface("data", bid).Msg("实时经纪买盘")
	}
}

func (ts *SDKTestSuite) TestRequestHistoryKL() {
	should := require.New(ts.T())

	res, err := ts.sdk.RequestHistoryKL(
		"HK.09988",
		adapt.KLType_Day,
		"2024-10-01",
		"2024-10-15",
		adapt.With("rehabType", adapt.RehabType_Forward),
		adapt.With("maxAckKLNum", 3), // 每次只取3条，模拟分页
	)
	should.NoError(err)
	should.Equal("09988", res.GetSecurity().GetCode())

	for _, kl := range res.GetKlList() {
		log.Info().Str("date", kl.GetTime()).Str("stock", res.GetName()).Float64("close", kl.GetClosePrice()).Msg("RequestHistoryKL")
	}

	next := res.GetNextReqKey()
	for len(next) > 0 {
		res, err = ts.sdk.RequestHistoryKL(
			"HK.09988",
			adapt.KLType_Day,
			"2024-10-01",
			"2024-10-15",
			adapt.With("rehabType", adapt.RehabType_Forward),
			adapt.With("maxAckKLNum", 3),
			adapt.With("nextReqKey", next),
		)
		should.NoError(err)

		for _, kl := range res.GetKlList() {
			log.Info().Str("date", kl.GetTime()).Str("stock", res.GetName()).Float64("close", kl.GetClosePrice()).Msg("RequestHistoryKL")
		}

		next = res.GetNextReqKey()
	}
}

func (ts *SDKTestSuite) TestRequestHistoryKLQuota() {
	should := require.New(ts.T())

	res, err := ts.sdk.RequestHistoryKLQuota(
		adapt.With("bGetDetail", true), // 可选：返回详细拉取过的历史纪录
	)
	should.NoError(err)
	log.Info().Interface("result", res).Msg("RequestHistoryKLQuota")
}

func (ts *SDKTestSuite) TestRequestRehab() {
	should := require.New(ts.T())

	res, err := ts.sdk.RequestRehab("HK.09988")
	should.NoError(err)

	for _, rehab := range res.GetRehabList() {
		log.Info().Interface("rehab", rehab).Msg("RequestRehab")
	}
}

func (ts *SDKTestSuite) TestGetStaticInfo() {
	should := require.New(ts.T())

	// use securities to filter
	res, err := ts.sdk.GetStaticInfo(adapt.WithSecurities([]string{"HK.09988", "HK.00700"}))
	should.NoError(err)

	for _, info := range res {
		log.Info().Interface("info", info).Msg("GetStaticInfo by securities")
	}

	// use market and secType to filter
	res2, err := ts.sdk.GetStaticInfo(
		adapt.With("market", adapt.QotMarket_HK),
		adapt.With("secType", adapt.SecurityType_Eqty),
	)
	should.NoError(err)
	log.Info().Int("num", len(res2)).Msg("GetStaticInfo by market")
}

func (ts *SDKTestSuite) TestGetSecuritySnapshot() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetSecuritySnapshot([]string{"HK.09988", "HK.00700"})
	should.NoError(err)

	for _, snap := range res {
		log.Info().Interface("snapshot", snap).Msg("GetSecuritySnapshot")
	}
}

func (ts *SDKTestSuite) TestGetPlateSet() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetPlateSet(adapt.QotMarket_HK, adapt.PlateSetType_Industry)
	should.NoError(err)

	for _, plate := range res {
		log.Info().Str("name", plate.GetName()).
			Int32("type", plate.GetPlateType()).
			Interface("plate", plate.GetPlate()).
			Msg("GetPlateSet")
	}
}

func (ts *SDKTestSuite) TestGetPlateSecurity() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetPlateSecurity(
		"HK.LIST1059",
		adapt.With("sortField", adapt.SortField_Turnover),
		adapt.With("ascend", false),
	)
	should.NoError(err)

	for _, info := range res {
		log.Info().Interface("info", info).Msg("GetPlateSecurity")
	}
}

func (ts *SDKTestSuite) TestGetReference() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetReference("HK.09988", adapt.ReferenceType_Warrant)
	should.NoError(err)
	log.Info().Int("num", len(res)).Msg("GetReference")
}

func (ts *SDKTestSuite) TestGetOwnerPlate() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetOwnerPlate([]string{"HK.09988"})
	should.NoError(err)
	log.Info().Interface("data", res).Msg("GetOwnerPlate")
}

func (ts *SDKTestSuite) TestGetOptionChain() {
	should := require.New(ts.T())

	beginTime := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	endTime := time.Now().Format("2006-01-02")

	res, err := ts.sdk.GetOptionChain("HK.09988", beginTime, endTime)
	should.NoError(err)
	log.Info().Int("num", len(res)).Msg("GetOptionChain")
}

func (ts *SDKTestSuite) TestGetWarrant() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetWarrant(0, 3,
		adapt.With("owner", adapt.NewSecurity("HK.00981")),
		adapt.With("status", adapt.WarrantStatus_Normal),
	)
	should.NoError(err)
	log.Info().Int32("count", res.GetAllCount()).Msg("GetWarrant")
	for _, warrant := range res.GetWarrantDataList() {
		log.Info().Interface("warrant", warrant).Msg("GetWarrant")
	}
}

func (ts *SDKTestSuite) TestGetCapitalFlow() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetCapitalFlow(
		"HK.09988",
		adapt.With("beginTime", time.Now().AddDate(0, 0, -1).Format("2006-01-02")),
		adapt.With("endTime", time.Now().Format("2006-01-02")),
		adapt.With("periodType", adapt.PeriodType_DAY),
	)
	should.NoError(err)
	log.Info().Interface("data", res).Msg("GetCapitalFlow")
}

func (ts *SDKTestSuite) TestGetCapitalDistribution() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetCapitalDistribution("HK.09988")
	should.NoError(err)
	log.Info().Interface("data", res).Msg("GetCapitalDistribution")
}

func (ts *SDKTestSuite) TestGetUserSecurity() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetUserSecurity("特别关注")
	should.NoError(err)
	log.Info().Int("count", len(res)).Msg("GetUserSecurity")
}

func (ts *SDKTestSuite) TestModifyUserSecurity() {
	should := require.New(ts.T())

	err := ts.sdk.ModifyUserSecurity(
		"特别关注",
		[]string{"HK.09988"},
		adapt.ModifyUserSecurityOp_Add,
	)
	should.Error(err) // 仅支持修改自定义分组，不支持修改系统分组
}

func (ts *SDKTestSuite) TestStockFilter() {
	should := require.New(ts.T())

	// f := &qotstockfilter.BaseFilter{
	// 	FieldName:  proto.Int32(int32(qotstockfilter.StockField_StockField_MarketVal)),
	// 	FilterMin:  proto.Float64(10000000000),
	// 	SortDir:    proto.Int32(int32(qotstockfilter.SortDir_SortDir_Ascend)),
	// 	IsNoFilter: proto.Bool(false),
	// }
	f := adapt.NewBaseFilter(
		qotstockfilter.StockField_StockField_MarketVal,
		10000000000,
		0,
		qotstockfilter.SortDir_SortDir_Ascend,
	)

	res, err := ts.sdk.StockFilter(
		adapt.QotMarket_HK,
		adapt.With("begin", 0),
		adapt.With("num", 10),
		adapt.WithBaseFilters(f),
	)
	should.NoError(err)
	log.Info().Int("count", int(res.GetAllCount())).Msg("StockFilter")
	for _, stock := range res.GetDataList() {
		log.Info().Interface("stock", stock).Msg("StockFilter")
	}
}

func (ts *SDKTestSuite) TestGetIpoList() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetIpoList(adapt.QotMarket_HK)
	should.NoError(err)

	for _, ipo := range res {
		log.Info().Interface("ipo", ipo).Msg("GetIpoList")
	}
}

func (ts *SDKTestSuite) TestGetFutureInfo() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetFutureInfo([]string{"HK.TCHmain"})
	should.NoError(err)
	log.Info().Interface("data", res).Msg("GetFutureInfo")
}

func (ts *SDKTestSuite) TestRequestTradeDate() {
	should := require.New(ts.T())

	res, err := ts.sdk.RequestTradeDate(adapt.QotMarket_HK, "", "2024-12-01", "2024-12-10")
	should.NoError(err)
	log.Info().Interface("data", res).Msg("RequestTradeDate")
}

func (ts *SDKTestSuite) TestSetPriceReminder() {
	should := require.New(ts.T())

	res, err := ts.sdk.SetPriceReminder(
		"HK.09988",
		adapt.SetPriceReminderOp_Add,
		adapt.With("type", adapt.PriceReminderType_PriceDown),
		adapt.With("freq", adapt.PriceReminderFreq_OnlyOnce),
		adapt.With("value", 80),
		adapt.With("note", "go sdk"),
	)
	should.NoError(err)
	log.Info().Int64("result", res).Msg("SetPriceReminder")
}

func (ts *SDKTestSuite) TestGetPriceReminder() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetPriceReminder("", adapt.QotMarket_HK)
	should.NoError(err)
	log.Info().Interface("data", res).Msg("GetPriceReminder")

	// remove all the reminders
	for _, reminder := range res {
		_, err := ts.sdk.SetPriceReminder(
			adapt.SecurityToCode(reminder.GetSecurity()),
			adapt.SetPriceReminderOp_DelAll,
		)
		should.NoError(err)
	}
}

func (ts *SDKTestSuite) TestGetUserSecurityGroup() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetUserSecurityGroup(adapt.GroupType_System)
	should.NoError(err)
	log.Info().Interface("data", res).Msg("GetUserSecurityGroup")
}
