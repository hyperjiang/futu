package client_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

var (
	alibaba = &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		Code:   proto.String("09988"),
	}
	tencent = &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		Code:   proto.String("00700"),
	}
	apple = &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_US_Security)),
		Code:   proto.String("AAPL"),
	}
	tobacco = &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_HK_Security)),
		Code:   proto.String("LIST1356"),
	}
	usAccount = &trdcommon.TrdHeader{
		TrdEnv:    proto.Int32(int32(trdcommon.TrdEnv_TrdEnv_Simulate)),
		AccID:     proto.Uint64(1619199),
		TrdMarket: proto.Int32(int32(trdcommon.TrdMarket_TrdMarket_US)),
	}
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

type ClientTestSuite struct {
	suite.Suite
	client *client.Client
}

// TestClientTestSuite runs the http client test suite
func TestClientTestSuite(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	suite.Run(t, new(ClientTestSuite))
}

// SetupSuite run once at the very start of the testing suite, before any tests are run.
func (ts *ClientTestSuite) SetupSuite() {
	var err error
	ts.client, err = client.New(
		client.WithPrivateKey(priKey),
		client.WithPublicKey(pubKey),
	)
	if err != nil {
		ts.T().SkipNow()
	}

	c2s := &qotsub.C2S{
		SecurityList: []*qotcommon.Security{alibaba, tencent},
		SubTypeList: []int32{
			int32(qotcommon.SubType_SubType_Basic),
			int32(qotcommon.SubType_SubType_RT),
			int32(qotcommon.SubType_SubType_KL_Day),
			int32(qotcommon.SubType_SubType_KL_3Min),
			int32(qotcommon.SubType_SubType_Ticker),
			int32(qotcommon.SubType_SubType_OrderBook),
			int32(qotcommon.SubType_SubType_Broker),
		},
		IsSubOrUnSub: proto.Bool(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = ts.client.QotSub(ctx, c2s)
	if err != nil {
		fmt.Println(err)
	}
}

// TearDownSuite run once at the very end of the testing suite, after all tests have been run.
func (ts *ClientTestSuite) TearDownSuite() {
	// time.Sleep(10 * time.Minute)
	if ts.client != nil {
		ts.client.Close()
	}
}

func (ts *ClientTestSuite) TestGetGlobalState() {
	should := require.New(ts.T())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := ts.client.GetGlobalState(ctx)
	should.NoError(err)

	fmt.Println(res)
}
