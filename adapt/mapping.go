package adapt

import (
	"strings"

	"github.com/hyperjiang/futu/pb/qotcommon"
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
