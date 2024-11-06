package adapt

import (
	"strings"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"google.golang.org/protobuf/proto"
)

var marketIDs = map[string]*int32{
	"HK": proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
	"US": proto.Int32(int32(qotcommon.QotMarket_QotMarket_US_Security)),
	"SH": proto.Int32(int32(qotcommon.QotMarket_QotMarket_CNSH_Security)),
	"SZ": proto.Int32(int32(qotcommon.QotMarket_QotMarket_CNSZ_Security)),
	"SG": proto.Int32(int32(qotcommon.QotMarket_QotMarket_SG_Security)),
	"JP": proto.Int32(int32(qotcommon.QotMarket_QotMarket_JP_Security)),
}

// GetMarketID returns the market ID by the market name.
func GetMarketID(name string) *int32 {
	id, ok := marketIDs[strings.ToUpper(name)]
	if ok {
		return id
	}

	return proto.Int32(0)
}

// GetMarketName returns the market name by the market ID.
func GetMarketName(id int32) string {
	for k, v := range marketIDs {
		if *v == id {
			return k
		}
	}

	return "Unknown"
}

const (
	SubType_None       = int32(qotcommon.SubType_SubType_None)       // 未知
	SubType_Basic      = int32(qotcommon.SubType_SubType_Basic)      // 基础报价
	SubType_OrderBook  = int32(qotcommon.SubType_SubType_OrderBook)  // 摆盘
	SubType_Ticker     = int32(qotcommon.SubType_SubType_Ticker)     // 逐笔
	SubType_RT         = int32(qotcommon.SubType_SubType_RT)         // 分时
	SubType_KL_Day     = int32(qotcommon.SubType_SubType_KL_Day)     // 日K
	SubType_KL_5Min    = int32(qotcommon.SubType_SubType_KL_5Min)    // 5分钟K
	SubType_KL_15Min   = int32(qotcommon.SubType_SubType_KL_15Min)   // 15分钟K
	SubType_KL_30Min   = int32(qotcommon.SubType_SubType_KL_30Min)   // 30分钟K
	SubType_KL_60Min   = int32(qotcommon.SubType_SubType_KL_60Min)   // 60分钟K
	SubType_KL_1Min    = int32(qotcommon.SubType_SubType_KL_1Min)    // 1分钟K
	SubType_KL_Week    = int32(qotcommon.SubType_SubType_KL_Week)    // 周K
	SubType_KL_Month   = int32(qotcommon.SubType_SubType_KL_Month)   // 月K
	SubType_Broker     = int32(qotcommon.SubType_SubType_Broker)     // 经纪队列
	SubType_KL_Qurater = int32(qotcommon.SubType_SubType_KL_Qurater) // 季K
	SubType_KL_Year    = int32(qotcommon.SubType_SubType_KL_Year)    // 年K
	SubType_KL_3Min    = int32(qotcommon.SubType_SubType_KL_3Min)    // 3分钟K
)

const (
	RehabType_None     = int32(qotcommon.RehabType_RehabType_None)     // 不复权
	RehabType_Forward  = int32(qotcommon.RehabType_RehabType_Forward)  // 前复权
	RehabType_Backward = int32(qotcommon.RehabType_RehabType_Backward) // 后复权
)

const (
	KLType_Unknown = int32(qotcommon.KLType_KLType_Unknown) // 未知
	KLType_1Min    = int32(qotcommon.KLType_KLType_1Min)    // 1分钟K
	KLType_Day     = int32(qotcommon.KLType_KLType_Day)     // 日K
	KLType_Week    = int32(qotcommon.KLType_KLType_Week)    // 周K
	KLType_Month   = int32(qotcommon.KLType_KLType_Month)   // 月K
	KLType_Year    = int32(qotcommon.KLType_KLType_Year)    // 年K
	KLType_5Min    = int32(qotcommon.KLType_KLType_5Min)    // 5分钟K
	KLType_15Min   = int32(qotcommon.KLType_KLType_15Min)   // 15分钟K
	KLType_30Min   = int32(qotcommon.KLType_KLType_30Min)   // 30分钟K
	KLType_60Min   = int32(qotcommon.KLType_KLType_60Min)   // 60分钟K
	KLType_3Min    = int32(qotcommon.KLType_KLType_3Min)    // 3分钟K
	KLType_Quarter = int32(qotcommon.KLType_KLType_Quarter) // 季K
)
