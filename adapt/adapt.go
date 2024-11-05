package adapt

import (
	"strings"

	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"google.golang.org/protobuf/proto"
)

// NewSecurity creates a new Security by a code string, e.g. "HK.00700".
func NewSecurity(code string) *qotcommon.Security {
	arr := strings.Split(code, ".")
	if len(arr) != 2 {
		return nil
	}

	return &qotcommon.Security{
		Market: GetMarketID(arr[0]),
		Code:   &arr[1],
	}
}

// NewSecurities creates a slice of Securities by a slice of code strings.
func NewSecurities(codes []string) []*qotcommon.Security {
	sa := make([]*qotcommon.Security, 0)
	for _, v := range codes {
		s := NewSecurity(v)
		if s != nil {
			sa = append(sa, s)
		}
	}

	return sa
}

// NewTradeAccount creates a new TrdHeader for a real trade account.
func NewTradeAccount(id uint64, market trdcommon.TrdMarket) *trdcommon.TrdHeader {
	return &trdcommon.TrdHeader{
		TrdEnv:    proto.Int32(int32(trdcommon.TrdEnv_TrdEnv_Real)),
		AccID:     proto.Uint64(id),
		TrdMarket: proto.Int32(int32(market)),
	}
}

// NewTestingTradeAccount creates a new TrdHeader for a testing trade account.
func NewTestingTradeAccount(id uint64, market trdcommon.TrdMarket) *trdcommon.TrdHeader {
	return &trdcommon.TrdHeader{
		TrdEnv:    proto.Int32(int32(trdcommon.TrdEnv_TrdEnv_Simulate)),
		AccID:     proto.Uint64(id),
		TrdMarket: proto.Int32(int32(market)),
	}
}
