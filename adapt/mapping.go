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

// 订阅类型
const (
	SubType_None       = int32(qotcommon.SubType_SubType_None)
	SubType_Basic      = int32(qotcommon.SubType_SubType_Basic)
	SubType_OrderBook  = int32(qotcommon.SubType_SubType_OrderBook)
	SubType_Ticker     = int32(qotcommon.SubType_SubType_Ticker)
	SubType_RT         = int32(qotcommon.SubType_SubType_RT)
	SubType_KL_Day     = int32(qotcommon.SubType_SubType_KL_Day)
	SubType_KL_5Min    = int32(qotcommon.SubType_SubType_KL_5Min)
	SubType_KL_15Min   = int32(qotcommon.SubType_SubType_KL_15Min)
	SubType_KL_30Min   = int32(qotcommon.SubType_SubType_KL_30Min)
	SubType_KL_60Min   = int32(qotcommon.SubType_SubType_KL_60Min)
	SubType_KL_1Min    = int32(qotcommon.SubType_SubType_KL_1Min)
	SubType_KL_Week    = int32(qotcommon.SubType_SubType_KL_Week)
	SubType_KL_Month   = int32(qotcommon.SubType_SubType_KL_Month)
	SubType_Broker     = int32(qotcommon.SubType_SubType_Broker)
	SubType_KL_Qurater = int32(qotcommon.SubType_SubType_KL_Qurater)
	SubType_KL_Year    = int32(qotcommon.SubType_SubType_KL_Year)
	SubType_KL_3Min    = int32(qotcommon.SubType_SubType_KL_3Min)
)
