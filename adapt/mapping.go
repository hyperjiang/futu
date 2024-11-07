package adapt

import (
	"strings"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotgetreference"
	"github.com/hyperjiang/futu/pb/qotmodifyusersecurity"
	"google.golang.org/protobuf/proto"
)

// 行情接口的市场常量
const (
	// QotMarket_HK 香港市场
	QotMarket_HK = int32(qotcommon.QotMarket_QotMarket_HK_Security)

	// QotMarket_US 美国市场
	QotMarket_US = int32(qotcommon.QotMarket_QotMarket_US_Security)

	// QotMarket_SH 上海市场
	QotMarket_SH = int32(qotcommon.QotMarket_QotMarket_CNSH_Security)

	// QotMarket_SZ 深圳市场
	QotMarket_SZ = int32(qotcommon.QotMarket_QotMarket_CNSZ_Security)

	// QotMarket_SG 新加坡市场
	QotMarket_SG = int32(qotcommon.QotMarket_QotMarket_SG_Security)

	// QotMarket_JP 日本市场
	QotMarket_JP = int32(qotcommon.QotMarket_QotMarket_JP_Security)
)

var marketIDs = map[string]*int32{
	"HK": proto.Int32(QotMarket_HK),
	"US": proto.Int32(QotMarket_US),
	"SH": proto.Int32(QotMarket_SH),
	"SZ": proto.Int32(QotMarket_SZ),
	"SG": proto.Int32(QotMarket_SG),
	"JP": proto.Int32(QotMarket_JP),
}

// GetMarketID 根据市场名称返回市场ID
func GetMarketID(name string) *int32 {
	id, ok := marketIDs[strings.ToUpper(name)]
	if ok {
		return id
	}

	return proto.Int32(0)
}

// GetMarketName 根据市场ID返回市场名称
func GetMarketName(id int32) string {
	for k, v := range marketIDs {
		if *v == id {
			return k
		}
	}

	return "Unknown"
}

// 订阅类型
const (
	// SubType_None 未知订阅类型
	SubType_None = int32(qotcommon.SubType_SubType_None)

	// SubType_Basic 基础报价
	SubType_Basic = int32(qotcommon.SubType_SubType_Basic)

	// SubType_OrderBook 摆盘
	SubType_OrderBook = int32(qotcommon.SubType_SubType_OrderBook)

	// SubType_Ticker 逐笔
	SubType_Ticker = int32(qotcommon.SubType_SubType_Ticker)

	// SubType_RT 分时
	SubType_RT = int32(qotcommon.SubType_SubType_RT)

	// SubType_KL_Day 日K线
	SubType_KL_Day = int32(qotcommon.SubType_SubType_KL_Day)

	// SubType_KL_5Min 5分钟K线
	SubType_KL_5Min = int32(qotcommon.SubType_SubType_KL_5Min)

	// SubType_KL_15Min 15分钟K线
	SubType_KL_15Min = int32(qotcommon.SubType_SubType_KL_15Min)

	// SubType_KL_30Min 30分钟K线
	SubType_KL_30Min = int32(qotcommon.SubType_SubType_KL_30Min)

	// SubType_KL_60Min 60分钟K线
	SubType_KL_60Min = int32(qotcommon.SubType_SubType_KL_60Min)

	// SubType_KL_1Min 1分钟K线
	SubType_KL_1Min = int32(qotcommon.SubType_SubType_KL_1Min)

	// SubType_KL_Week 周K线
	SubType_KL_Week = int32(qotcommon.SubType_SubType_KL_Week)

	// SubType_KL_Month 月K线
	SubType_KL_Month = int32(qotcommon.SubType_SubType_KL_Month)

	// SubType_Broker 经纪队列
	SubType_Broker = int32(qotcommon.SubType_SubType_Broker)

	// SubType_KL_Qurater 季K线
	SubType_KL_Qurater = int32(qotcommon.SubType_SubType_KL_Qurater)

	// SubType_KL_Year 年K线
	SubType_KL_Year = int32(qotcommon.SubType_SubType_KL_Year)

	// SubType_KL_3Min 3分钟K线
	SubType_KL_3Min = int32(qotcommon.SubType_SubType_KL_3Min)
)

// 复权类型
const (
	// RehabType_None 不复权
	RehabType_None = int32(qotcommon.RehabType_RehabType_None)

	// RehabType_Forward 前复权
	RehabType_Forward = int32(qotcommon.RehabType_RehabType_Forward)

	// RehabType_Backward 后复权
	RehabType_Backward = int32(qotcommon.RehabType_RehabType_Backward)
)

// K线类型
const (
	// KLType_Unknown 未知
	KLType_Unknown = int32(qotcommon.KLType_KLType_Unknown)

	// KLType_1Min 1分钟K线
	KLType_1Min = int32(qotcommon.KLType_KLType_1Min)

	// KLType_Day 日K线
	KLType_Day = int32(qotcommon.KLType_KLType_Day)

	// KLType_Week 周K线
	KLType_Week = int32(qotcommon.KLType_KLType_Week)

	// KLType_Month 月K线
	KLType_Month = int32(qotcommon.KLType_KLType_Month)

	// KLType_Year 年K线
	KLType_Year = int32(qotcommon.KLType_KLType_Year)

	// KLType_5Min 5分钟K线
	KLType_5Min = int32(qotcommon.KLType_KLType_5Min)

	// KLType_15Min 15分钟K线
	KLType_15Min = int32(qotcommon.KLType_KLType_15Min)

	// KLType_30Min 30分钟K线
	KLType_30Min = int32(qotcommon.KLType_KLType_30Min)

	// KLType_60Min 60分钟K线
	KLType_60Min = int32(qotcommon.KLType_KLType_60Min)

	// KLType_3Min 3分钟K线
	KLType_3Min = int32(qotcommon.KLType_KLType_3Min)

	// KLType_Quarter 季K线
	KLType_Quarter = int32(qotcommon.KLType_KLType_Quarter)
)

// 证券类型
const (
	// SecurityType_Unknown 未知证券类型
	SecurityType_Unknown = int32(qotcommon.SecurityType_SecurityType_Unknown)

	// SecurityType_Bond 债券
	SecurityType_Bond = int32(qotcommon.SecurityType_SecurityType_Bond)

	// SecurityType_Bwrt 一揽子权证
	SecurityType_Bwrt = int32(qotcommon.SecurityType_SecurityType_Bwrt)

	// SecurityType_Eqty 正股
	SecurityType_Eqty = int32(qotcommon.SecurityType_SecurityType_Eqty)

	// SecurityType_Trust 信托或基金
	SecurityType_Trust = int32(qotcommon.SecurityType_SecurityType_Trust)

	// SecurityType_Warrant 窝轮
	SecurityType_Warrant = int32(qotcommon.SecurityType_SecurityType_Warrant)

	// SecurityType_Index 指数
	SecurityType_Index = int32(qotcommon.SecurityType_SecurityType_Index)

	// SecurityType_Plate 板块
	SecurityType_Plate = int32(qotcommon.SecurityType_SecurityType_Plate)

	// SecurityType_Drvt 期权
	SecurityType_Drvt = int32(qotcommon.SecurityType_SecurityType_Drvt)

	// SecurityType_PlateSet 板块集
	SecurityType_PlateSet = int32(qotcommon.SecurityType_SecurityType_PlateSet)

	// SecurityType_Future 期货
	SecurityType_Future = int32(qotcommon.SecurityType_SecurityType_Future)
)

// 板块集合类型
const (
	// PlateSetType_All 所有板块
	PlateSetType_All = int32(qotcommon.PlateSetType_PlateSetType_All)

	// PlateSetType_Industry 行业板块
	PlateSetType_Industry = int32(qotcommon.PlateSetType_PlateSetType_Industry)

	// PlateSetType_Region 地区板块，港美股市场的地域分类数据暂为空
	PlateSetType_Region = int32(qotcommon.PlateSetType_PlateSetType_Region)

	// PlateSetType_Concept 概念板块
	PlateSetType_Concept = int32(qotcommon.PlateSetType_PlateSetType_Concept)

	// PlateSetType_Other 其他板块，仅用于3207（获取股票所属板块）协议返回，不可作为其他协议的请求参数
	PlateSetType_Other = int32(qotcommon.PlateSetType_PlateSetType_Other)
)

// 排序字段
const (
	// SortField_Unknow 未知排序字段
	SortField_Unknow = int32(qotcommon.SortField_SortField_Unknow)

	// SortField_Code 代码
	SortField_Code = int32(qotcommon.SortField_SortField_Code)

	// SortField_CurPrice 当前价格
	SortField_CurPrice = int32(qotcommon.SortField_SortField_CurPrice)

	// SortField_PriceChangeVal 价格变动值
	SortField_PriceChangeVal = int32(qotcommon.SortField_SortField_PriceChangeVal)

	// SortField_ChangeRate 涨跌幅
	SortField_ChangeRate = int32(qotcommon.SortField_SortField_ChangeRate)

	// SortField_Status 状态
	SortField_Status = int32(qotcommon.SortField_SortField_Status)

	// SortField_BidPrice 买价
	SortField_BidPrice = int32(qotcommon.SortField_SortField_BidPrice)

	// SortField_AskPrice 卖价
	SortField_AskPrice = int32(qotcommon.SortField_SortField_AskPrice)

	// SortField_BidVol 买量
	SortField_BidVol = int32(qotcommon.SortField_SortField_BidVol)

	// SortField_AskVol 卖量
	SortField_AskVol = int32(qotcommon.SortField_SortField_AskVol)

	// SortField_Volume 成交量
	SortField_Volume = int32(qotcommon.SortField_SortField_Volume)

	// SortField_Turnover 成交额
	SortField_Turnover = int32(qotcommon.SortField_SortField_Turnover)

	// SortField_Amplitude 振幅
	SortField_Amplitude = int32(qotcommon.SortField_SortField_Amplitude)

	// 以下排序字段只支持用于 Qot_GetWarrant 协议

	// SortField_Score 综合评分
	SortField_Score = int32(qotcommon.SortField_SortField_Score)

	// SortField_Premium 溢价
	SortField_Premium = int32(qotcommon.SortField_SortField_Premium)

	// SortField_EffectiveLeverage 有效杠杆
	SortField_EffectiveLeverage = int32(qotcommon.SortField_SortField_EffectiveLeverage)

	// SortField_Delta 对冲值
	SortField_Delta = int32(qotcommon.SortField_SortField_Delta)

	// SortField_ImpliedVolatility 隐含波动率
	SortField_ImpliedVolatility = int32(qotcommon.SortField_SortField_ImpliedVolatility)

	// SortField_Type 类型
	SortField_Type = int32(qotcommon.SortField_SortField_Type)

	// SortField_StrikePrice 行使价
	SortField_StrikePrice = int32(qotcommon.SortField_SortField_StrikePrice)

	// SortField_BreakEvenPoint 打和点
	SortField_BreakEvenPoint = int32(qotcommon.SortField_SortField_BreakEvenPoint)

	// SortField_MaturityTime 到期时间
	SortField_MaturityTime = int32(qotcommon.SortField_SortField_MaturityTime)

	// SortField_ListTime 上市时间
	SortField_ListTime = int32(qotcommon.SortField_SortField_ListTime)

	// SortField_LastTradeTime 最后交易时间
	SortField_LastTradeTime = int32(qotcommon.SortField_SortField_LastTradeTime)

	// SortField_Leverage 杠杆比率
	SortField_Leverage = int32(qotcommon.SortField_SortField_Leverage)

	// SortField_InOutMoney 实际杠杆
	SortField_InOutMoney = int32(qotcommon.SortField_SortField_InOutMoney)

	// SortField_RecoveryPrice 回收价
	SortField_RecoveryPrice = int32(qotcommon.SortField_SortField_RecoveryPrice)

	// SortField_ChangePrice 涨跌价
	SortField_ChangePrice = int32(qotcommon.SortField_SortField_ChangePrice)

	// SortField_Change 涨跌幅
	SortField_Change = int32(qotcommon.SortField_SortField_Change)

	// SortField_StreetRate 街货比率
	SortField_StreetRate = int32(qotcommon.SortField_SortField_StreetRate)

	// SortField_StreetVol 街货量
	SortField_StreetVol = int32(qotcommon.SortField_SortField_StreetVol)

	// SortField_WarrantName 窝轮名称
	SortField_WarrantName = int32(qotcommon.SortField_SortField_WarrantName)

	// SortField_Issuer 发行人
	SortField_Issuer = int32(qotcommon.SortField_SortField_Issuer)

	// SortField_LotSize 每手数量
	SortField_LotSize = int32(qotcommon.SortField_SortField_LotSize)

	// SortField_IssueSize 发行量
	SortField_IssueSize = int32(qotcommon.SortField_SortField_IssueSize)

	// SortField_UpperStrikePrice 上限价
	SortField_UpperStrikePrice = int32(qotcommon.SortField_SortField_UpperStrikePrice)

	// SortField_LowerStrikePrice 下限价
	SortField_LowerStrikePrice = int32(qotcommon.SortField_SortField_LowerStrikePrice)

	// SortField_InLinePriceStatus 价格状态
	SortField_InLinePriceStatus = int32(qotcommon.SortField_SortField_InLinePriceStatus)

	// 以下排序字段只支持用于 Qot_GetPlateSecurity 协议，并仅支持美股

	// SortField_PreCurPrice 盘前价格
	SortField_PreCurPrice = int32(qotcommon.SortField_SortField_PreCurPrice)

	// SortField_AfterCurPrice 盘后价格
	SortField_AfterCurPrice = int32(qotcommon.SortField_SortField_AfterCurPrice)

	// SortField_PrePriceChangeVal 盘前涨跌额
	SortField_PrePriceChangeVal = int32(qotcommon.SortField_SortField_PrePriceChangeVal)

	// SortField_AfterPriceChangeVal 盘后涨跌额
	SortField_AfterPriceChangeVal = int32(qotcommon.SortField_SortField_AfterPriceChangeVal)

	// SortField_PreChangeRate 盘前涨跌幅
	SortField_PreChangeRate = int32(qotcommon.SortField_SortField_PreChangeRate)

	// SortField_AfterChangeRate 盘后涨跌幅
	SortField_AfterChangeRate = int32(qotcommon.SortField_SortField_AfterChangeRate)

	// SortField_PreAmplitude 盘前振幅
	SortField_PreAmplitude = int32(qotcommon.SortField_SortField_PreAmplitude)

	// SortField_AfterAmplitude 盘后振幅
	SortField_AfterAmplitude = int32(qotcommon.SortField_SortField_AfterAmplitude)

	// SortField_PreTurnover 盘前成交额
	SortField_PreTurnover = int32(qotcommon.SortField_SortField_PreTurnover)

	// SortField_AfterTurnover 盘后成交额
	SortField_AfterTurnover = int32(qotcommon.SortField_SortField_AfterTurnover)

	// 以下排序字段只支持用于 Qot_GetPlateSecurity 协议，并仅支持期货

	// SortField_LastSettlePrice 上次结算价
	SortField_LastSettlePrice = int32(qotcommon.SortField_SortField_LastSettlePrice)

	// SortField_Position 持仓量
	SortField_Position = int32(qotcommon.SortField_SortField_Position)

	// SortField_PositionChange 持仓量变化
	SortField_PositionChange = int32(qotcommon.SortField_SortField_PositionChange)
)

// 参考类型
const (
	// ReferenceType_Unknow 未知
	ReferenceType_Unknow = int32(qotgetreference.ReferenceType_ReferenceType_Unknow)

	// ReferenceType_Warrant 正股相关的窝轮
	ReferenceType_Warrant = int32(qotgetreference.ReferenceType_ReferenceType_Warrant)

	// ReferenceType_Future 期货主连的相关合约
	ReferenceType_Future = int32(qotgetreference.ReferenceType_ReferenceType_Future)
)

// 窝轮类型
const (
	// WarrantType_Unknown 未知
	WarrantType_Unknown = int32(qotcommon.WarrantType_WarrantType_Unknown)

	// WarrantType_Buy 认购
	WarrantType_Buy = int32(qotcommon.WarrantType_WarrantType_Buy)

	// WarrantType_Sell 认沽
	WarrantType_Sell = int32(qotcommon.WarrantType_WarrantType_Sell)

	// WarrantType_Bull 看涨型窝轮
	WarrantType_Bull = int32(qotcommon.WarrantType_WarrantType_Bull)

	// WarrantType_Bear 看跌型窝轮
	WarrantType_Bear = int32(qotcommon.WarrantType_WarrantType_Bear)

	// WarrantType_InLine 界内证
	WarrantType_InLine = int32(qotcommon.WarrantType_WarrantType_InLine)
)

// 窝轮状态
const (
	// WarrantStatus_Unknow 未知
	WarrantStatus_Unknow = int32(qotcommon.WarrantStatus_WarrantStatus_Unknow)

	// WarrantStatus_Normal 正常
	WarrantStatus_Normal = int32(qotcommon.WarrantStatus_WarrantStatus_Normal)

	// WarrantStatus_Suspend 停牌
	WarrantStatus_Suspend = int32(qotcommon.WarrantStatus_WarrantStatus_Suspend)

	// WarrantStatus_StopTrade 终止交易
	WarrantStatus_StopTrade = int32(qotcommon.WarrantStatus_WarrantStatus_StopTrade)

	// WarrantStatus_PendingListing 待上市
	WarrantStatus_PendingListing = int32(qotcommon.WarrantStatus_WarrantStatus_PendingListing)
)

// 周期类型
const (
	// PeriodType_Unknown 未知
	PeriodType_Unknown = int32(qotcommon.PeriodType_PeriodType_Unknown)

	// PeriodType_INTRADAY 实时
	PeriodType_INTRADAY = int32(qotcommon.PeriodType_PeriodType_INTRADAY)

	// PeriodType_DAY 日
	PeriodType_DAY = int32(qotcommon.PeriodType_PeriodType_DAY)

	// PeriodType_WEEK 周
	PeriodType_WEEK = int32(qotcommon.PeriodType_PeriodType_WEEK)

	// PeriodType_MONTH 月
	PeriodType_MONTH = int32(qotcommon.PeriodType_PeriodType_MONTH)
)

// 自选股操作类型
const (
	// ModifyUserSecurityOp_Unknown 未知
	ModifyUserSecurityOp_Unknown = int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_Unknown)

	// ModifyUserSecurityOp_Add 新增自选
	ModifyUserSecurityOp_Add = int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_Add)

	// ModifyUserSecurityOp_Del 删除自选
	ModifyUserSecurityOp_Del = int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_Del)

	// ModifyUserSecurityOp_MoveOut 移出分组
	ModifyUserSecurityOp_MoveOut = int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_MoveOut)
)
