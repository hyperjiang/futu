package futu_test

import (
	"os"
	"testing"
	"time"

	"github.com/hyperjiang/futu"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
}

// TearDownSuite run once at the very end of the testing suite, after all tests have been run.
func (ts *FutuTestSuite) TearDownSuite() {
	time.Sleep(1 * time.Second)
	if ts.client != nil {
		ts.client.Close()
	}
}

func (ts *FutuTestSuite) TestGet() {
	should := require.New(ts.T())
	should.NoError(nil)
}
