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
