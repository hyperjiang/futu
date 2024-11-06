package client_test

import (
	"context"
	"time"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetbasicqot"
	"github.com/hyperjiang/futu/pb/qotgetbroker"
	"github.com/hyperjiang/futu/pb/qotgetcapitaldistribution"
	"github.com/hyperjiang/futu/pb/qotgetcapitalflow"
	"github.com/hyperjiang/futu/pb/qotgetfutureinfo"
	"github.com/hyperjiang/futu/pb/qotgetipolist"
	"github.com/hyperjiang/futu/pb/qotgetkl"
	"github.com/hyperjiang/futu/pb/qotgetmarketstate"
	"github.com/hyperjiang/futu/pb/qotgetoptionchain"
	"github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate"
	"github.com/hyperjiang/futu/pb/qotgetorderbook"
	"github.com/hyperjiang/futu/pb/qotgetownerplate"
	"github.com/hyperjiang/futu/pb/qotgetplatesecurity"
	"github.com/hyperjiang/futu/pb/qotgetplateset"
	"github.com/hyperjiang/futu/pb/qotgetpricereminder"
	"github.com/hyperjiang/futu/pb/qotgetreference"
	"github.com/hyperjiang/futu/pb/qotgetrt"
	"github.com/hyperjiang/futu/pb/qotgetsecuritysnapshot"
	"github.com/hyperjiang/futu/pb/qotgetstaticinfo"
	"github.com/hyperjiang/futu/pb/qotgetsubinfo"
	"github.com/hyperjiang/futu/pb/qotgetticker"
	"github.com/hyperjiang/futu/pb/qotgetusersecurity"
	"github.com/hyperjiang/futu/pb/qotgetusersecuritygroup"
	"github.com/hyperjiang/futu/pb/qotgetwarrant"
	"github.com/hyperjiang/futu/pb/qotmodifyusersecurity"
	"github.com/hyperjiang/futu/pb/qotrequesthistorykl"
	"github.com/hyperjiang/futu/pb/qotrequesthistoryklquota"
	"github.com/hyperjiang/futu/pb/qotrequestrehab"
	"github.com/hyperjiang/futu/pb/qotrequesttradedate"
	"github.com/hyperjiang/futu/pb/qotsetpricereminder"
	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func (ts *ClientTestSuite) TestQotGetSubInfo() {
	should := require.New(ts.T())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	info, err := ts.client.QotGetSubInfo(ctx, &qotgetsubinfo.C2S{})
	should.NoError(err)
	log.Info().Str("data", info.String()).Msg("QotGetSubInfo")
}

func (ts *ClientTestSuite) TestQotGetBasicQot() {
	should := require.New(ts.T())

	c2s := &qotgetbasicqot.C2S{
		SecurityList: []*qotcommon.Security{alibaba},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	s2c, err := ts.client.QotGetBasicQot(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", s2c.GetBasicQotList()).Msg("QotGetBasicQot")
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
	log.Info().Interface("data", s2c.GetKlList()).Msg("QotGetKL")
}

func (ts *ClientTestSuite) TestQotGetRT() {
	should := require.New(ts.T())

	c2s := &qotgetrt.C2S{
		Security: tencent,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s2c, err := ts.client.QotGetRT(ctx, c2s)
	should.NoError(err)
	should.Equal(tencent.GetCode(), s2c.GetSecurity().GetCode())
	log.Info().Str("stock", s2c.GetName()).Int("num", len(s2c.GetRtList())).Msg("QotGetRT")
}

func (ts *ClientTestSuite) TestQotGetTicker() {
	should := require.New(ts.T())

	c2s := &qotgetticker.C2S{
		Security:  tencent,
		MaxRetNum: proto.Int32(100),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s2c, err := ts.client.QotGetTicker(ctx, c2s)
	should.NoError(err)
	should.Equal(tencent.GetCode(), s2c.GetSecurity().GetCode())
	log.Info().Str("stock", s2c.GetName()).Int("num", len(s2c.GetTickerList())).Msg("QotGetTicker")
}

func (ts *ClientTestSuite) TestQotGetOrderBook() {
	should := require.New(ts.T())

	c2s := &qotgetorderbook.C2S{
		Security: tencent,
		Num:      proto.Int32(10),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s2c, err := ts.client.QotGetOrderBook(ctx, c2s)
	should.NoError(err)
	should.Equal(tencent.GetCode(), s2c.GetSecurity().GetCode())
	log.Info().Str("stock", s2c.GetName()).
		Int("实时卖盘", len(s2c.GetOrderBookAskList())).
		Int("实时买盘", len(s2c.GetOrderBookBidList())).
		Msg("QotGetOrderBook")
	for _, ask := range s2c.GetOrderBookAskList() {
		log.Info().Interface("data", ask).Msg("实时卖盘")
	}
	for _, bid := range s2c.GetOrderBookBidList() {
		log.Info().Interface("data", bid).Msg("实时买盘")
	}
}

func (ts *ClientTestSuite) TestQotGetBroker() {
	should := require.New(ts.T())

	c2s := &qotgetbroker.C2S{
		Security: tencent,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s2c, err := ts.client.QotGetBroker(ctx, c2s)
	should.NoError(err)
	should.Equal(tencent.GetCode(), s2c.GetSecurity().GetCode())
	log.Info().Str("stock", s2c.GetName()).
		Int("实时经纪卖盘", len(s2c.GetBrokerAskList())).
		Int("实时经纪买盘", len(s2c.GetBrokerBidList())).
		Msg("QotGetBroker")
	for _, ask := range s2c.GetBrokerAskList() {
		log.Info().Interface("data", ask).Msg("实时经纪卖盘")
	}
	for _, bid := range s2c.GetBrokerBidList() {
		log.Info().Interface("data", bid).Msg("实时经纪买盘")
	}
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

func (ts *ClientTestSuite) TestQotRequestHistoryKLQuota() {
	should := require.New(ts.T())
	c2s := &qotrequesthistoryklquota.C2S{
		BGetDetail: proto.Bool(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotRequestHistoryKLQuota(ctx, c2s)
	should.NoError(err)
	log.Info().Str("data", res.String()).Msg("QotRequestHistoryKLQuota")
}

func (ts *ClientTestSuite) TestQotRequestRehab() {
	should := require.New(ts.T())
	c2s := &qotrequestrehab.C2S{
		Security: tencent,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotRequestRehab(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetRehabList()).Msg("QotRequestRehab")
}

func (ts *ClientTestSuite) TestQotGetStaticInfo() {
	should := require.New(ts.T())
	c2s := &qotgetstaticinfo.C2S{
		// Market:       proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		// SecType:      proto.Int32(int32(qotcommon.SecurityType_SecurityType_Eqty)),
		SecurityList: []*qotcommon.Security{alibaba, tencent},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetStaticInfo(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetStaticInfoList()).Msg("QotGetStaticInfo")
}

func (ts *ClientTestSuite) TestQotGetPlateSet() {
	should := require.New(ts.T())
	c2s := &qotgetplateset.C2S{
		Market:       proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		PlateSetType: proto.Int32(int32(qotcommon.PlateSetType_PlateSetType_Industry)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetPlateSet(ctx, c2s)
	should.NoError(err)
	for _, plate := range res.GetPlateInfoList() {
		plate.GetPlate()
		log.Info().Str("name", plate.GetName()).
			Int32("type", plate.GetPlateType()).
			Interface("plate", plate.GetPlate()).
			Msg("QotGetPlateSet")
	}
}

func (ts *ClientTestSuite) TestQotGetPlateSecurity() {
	should := require.New(ts.T())
	c2s := &qotgetplatesecurity.C2S{
		Plate: tobacco,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetPlateSecurity(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetStaticInfoList()).Msg("QotGetPlateSecurity")
}

func (ts *ClientTestSuite) TestQotGetReference() {
	should := require.New(ts.T())
	c2s := &qotgetreference.C2S{
		Security:      tencent,
		ReferenceType: proto.Int32(int32(qotgetreference.ReferenceType_ReferenceType_Warrant)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := ts.client.QotGetReference(ctx, c2s)
	should.NoError(err)
	log.Info().Int("num", len(res.GetStaticInfoList())).Msg("QotGetReference")
	// for _, info := range res.GetStaticInfoList() {
	// 	log.Info().Interface("basic", info.GetBasic()).
	// 		Interface("future", info.GetFutureExData()).
	// 		Interface("warrant", info.GetWarrantExData()).
	// 		Interface("option", info.GetOptionExData()).
	// 		Msg("QotGetReference")
	// }
}

func (ts *ClientTestSuite) TestQotGetOwnerPlate() {
	should := require.New(ts.T())
	c2s := &qotgetownerplate.C2S{
		SecurityList: []*qotcommon.Security{tencent},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetOwnerPlate(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetOwnerPlateList()).Msg("QotGetOwnerPlate")
}

func (ts *ClientTestSuite) TestQotGetOptionChain() {
	should := require.New(ts.T())
	c2s := &qotgetoptionchain.C2S{
		Owner:     tencent,
		BeginTime: proto.String(time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")),
		EndTime:   proto.String(time.Now().Format("2006-01-02 15:04:05")),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetOptionChain(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetOptionChain()).Msg("QotGetOptionChain")
}

func (ts *ClientTestSuite) TestQotGetWarrant() {
	should := require.New(ts.T())
	c2s := &qotgetwarrant.C2S{
		Owner:     tencent,
		Begin:     proto.Int32(0),
		Num:       proto.Int32(3),
		SortField: proto.Int32(int32(qotcommon.SortField_SortField_Code)),
		Ascend:    proto.Bool(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetWarrant(ctx, c2s)
	should.NoError(err)
	log.Info().Int32("count", res.GetAllCount()).Msg("QotGetWarrant")
	for _, warrant := range res.GetWarrantDataList() {
		log.Info().Interface("warrant", warrant).Msg("QotGetWarrant")
	}
}

func (ts *ClientTestSuite) TestQotGetCapitalFlow() {
	should := require.New(ts.T())
	c2s := &qotgetcapitalflow.C2S{
		Security:   tencent,
		PeriodType: proto.Int32(int32(qotcommon.PeriodType_PeriodType_DAY)),
		BeginTime:  proto.String("2024-11-01"),
		EndTime:    proto.String("2024-11-02"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	res, err := ts.client.QotGetCapitalFlow(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetFlowItemList()).Msg("QotGetCapitalFlow")
}

func (ts *ClientTestSuite) TestQotGetCapitalDistribution() {
	should := require.New(ts.T())
	c2s := &qotgetcapitaldistribution.C2S{
		Security: tencent,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetCapitalDistribution(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res).Msg("QotGetCapitalDistribution")
}

func (ts *ClientTestSuite) TestQotGetUserSecurity() {
	should := require.New(ts.T())
	c2s := &qotgetusersecurity.C2S{
		GroupName: proto.String("特别关注"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetUserSecurity(ctx, c2s)
	should.NoError(err)
	log.Info().Int("count", len(res.GetStaticInfoList())).Msg("QotGetUserSecurity")
}

func (ts *ClientTestSuite) TestQotModifyUserSecurity() {
	should := require.New(ts.T())
	c2s := &qotmodifyusersecurity.C2S{
		GroupName:    proto.String("特别关注"),
		SecurityList: []*qotcommon.Security{alibaba},
		Op:           proto.Int32(int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_Add)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := ts.client.QotModifyUserSecurity(ctx, c2s)
	should.Error(err) // 仅支持修改自定义分组，不支持修改系统分组
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
		Begin:          &begin,
		Num:            &num,
		Market:         (*int32)(qotcommon.QotMarket_QotMarket_HK_Security.Enum()),
		BaseFilterList: []*qotstockfilter.BaseFilter{f},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotStockFilter(ctx, c2s)
	should.NoError(err)

	log.Info().Int32("count", res.GetAllCount()).Msg("QotStockFilter")

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
		log.Info().Interface("data", snapshot.GetSnapshotList()[0]).Msg("QotGetSecuritySnapshot")
	}
}

func (ts *ClientTestSuite) TestQotGetIpoList() {
	should := require.New(ts.T())
	c2s := &qotgetipolist.C2S{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetIpoList(ctx, c2s)
	should.NoError(err)
	for _, ipo := range res.GetIpoList() {
		log.Info().Interface("data", ipo).Msg("QotGetIpoList")
	}
}

func (ts *ClientTestSuite) TestQotGetFutureInfo() {
	future := &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		Code:   proto.String("TCHmain"), // 腾讯主连
	}

	should := require.New(ts.T())
	c2s := &qotgetfutureinfo.C2S{
		SecurityList: []*qotcommon.Security{future},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetFutureInfo(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res).Msg("QotGetFutureInfo")
}

func (ts *ClientTestSuite) TestQotRequestTradeDate() {
	should := require.New(ts.T())
	c2s := &qotrequesttradedate.C2S{
		Market:    proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		BeginTime: proto.String(time.Now().Format("2006-01-02")),
		EndTime:   proto.String(time.Now().Add(time.Hour * 24 * 7).Format("2006-01-02")),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotRequestTradeDate(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetTradeDateList()).Msg("QotRequestTradeDate")
}

func (ts *ClientTestSuite) TestQotSetPriceReminder() {
	should := require.New(ts.T())
	c2s := &qotsetpricereminder.C2S{
		Security: tencent,
		Op:       proto.Int32(int32(qotsetpricereminder.SetPriceReminderOp_SetPriceReminderOp_Add)),
		Type:     proto.Int32(int32(qotcommon.PriceReminderType_PriceReminderType_PriceDown)),
		Freq:     proto.Int32(int32(qotcommon.PriceReminderFreq_PriceReminderFreq_OnlyOnce)),
		Value:    proto.Float64(300),
		Note:     proto.String("go sdk"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotSetPriceReminder(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res).Msg("QotSetPriceReminder")
}

func (ts *ClientTestSuite) TestQotGetPriceReminder() {
	should := require.New(ts.T())
	c2s := &qotgetpricereminder.C2S{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetPriceReminder(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetPriceReminderList()).Msg("QotGetPriceReminder")
}

func (ts *ClientTestSuite) TestQotGetUserSecurityGroup() {
	should := require.New(ts.T())
	c2s := &qotgetusersecuritygroup.C2S{
		GroupType: proto.Int32(int32(qotgetusersecuritygroup.GroupType_GroupType_Custom)),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetUserSecurityGroup(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetGroupList()).Msg("QotGetUserSecurityGroup")
}

func (ts *ClientTestSuite) TestQotGetMarketState() {
	should := require.New(ts.T())
	c2s := &qotgetmarketstate.C2S{
		SecurityList: []*qotcommon.Security{tencent},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetMarketState(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetMarketInfoList()).Msg("QotGetMarketState")
}

func (ts *ClientTestSuite) TestQotGetOptionExpirationDate() {
	should := require.New(ts.T())
	c2s := &qotgetoptionexpirationdate.C2S{
		Owner: tencent,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.QotGetOptionExpirationDate(ctx, c2s)
	should.NoError(err)
	log.Info().Interface("data", res.GetDateList()).Msg("QotGetOptionExpirationDate")
}
