package futu_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hyperjiang/futu"
	"github.com/hyperjiang/futu/pb/qotcommon"
	"github.com/hyperjiang/futu/pb/qotsub"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

var (
	alibaba = &qotcommon.Security{
		Market: (*int32)(qotcommon.QotMarket_QotMarket_HK_Security.Enum()),
		Code:   proto.String("09988"),
	}
	tencent = &qotcommon.Security{
		Market: (*int32)(qotcommon.QotMarket_QotMarket_HK_Security.Enum()),
		Code:   proto.String("00700"),
	}
)

type FutuTestSuite struct {
	suite.Suite
	client *futu.Client
}

// TestFutuTestSuite runs the http client test suite
func TestFutuTestSuite(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	suite.Run(t, new(FutuTestSuite))
}

// SetupSuite run once at the very start of the testing suite, before any tests are run.
func (ts *FutuTestSuite) SetupSuite() {
	var err error
	ts.client, err = futu.NewClient()
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
		},
		IsSubOrUnSub: proto.Bool(true),
	}

	should := require.New(ts.T())
	err = ts.client.QotSub(c2s)
	should.NoError(err)
}

// TearDownSuite run once at the very end of the testing suite, after all tests have been run.
func (ts *FutuTestSuite) TearDownSuite() {
	// time.Sleep(10 * time.Minute)
	if ts.client != nil {
		ts.client.Close()
	}
}

func (ts *FutuTestSuite) TestGetGlobalState() {
	should := require.New(ts.T())

	res, err := ts.client.GetGlobalState()
	should.NoError(err)

	fmt.Println(res)
}
