package futu_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hyperjiang/futu"
	"github.com/hyperjiang/futu/adapt"
	"github.com/hyperjiang/futu/client"
	"github.com/hyperjiang/futu/pb/notify"
	"github.com/hyperjiang/futu/pb/trdcommon"
	"github.com/hyperjiang/futu/protoid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

var (
	tobacco   = adapt.NewSecurity("HK.LIST1356")
	usAccount = adapt.NewTestingTradeAccount(1619199, trdcommon.TrdMarket_TrdMarket_US)
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

	err = ts.sdk.Subscribe(
		adapt.With("securityList", adapt.NewSecurities([]string{"HK.09988", "HK.00700"})),
		adapt.With("subTypeList", []int32{
			adapt.SubType_Basic,
			adapt.SubType_RT,
			adapt.SubType_KL_Day,
			adapt.SubType_KL_3Min,
			adapt.SubType_Ticker,
			adapt.SubType_OrderBook,
			adapt.SubType_Broker,
		}),
		adapt.With("isSubOrUnSub", true),
	)
	if err != nil {
		fmt.Println(err)
	}

	ts.sdk.RegisterHandler(protoid.Notify, func(s2c proto.Message) error {
		msg := s2c.(*notify.S2C)
		log.Info().Interface("s2c", msg).Msg("custom handler")
		return nil
	})
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

func (ts *SDKTestSuite) TestQotGetBasicQot() {
	should := require.New(ts.T())

	res, err := ts.sdk.GetBasicQot([]string{"HK.00700", "HK.09988"})
	should.NoError(err)
	for _, qot := range res {
		log.Info().Interface("qot", qot).Msg("GetBasicQot")
	}
}
