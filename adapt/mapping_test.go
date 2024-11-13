package adapt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMarketID(t *testing.T) {
	should := require.New(t)

	should.Equal(int32(1), *GetMarketID("HK"))
	should.Equal(int32(0), *GetMarketID("XX"))
	should.Equal(int32(0), *GetMarketID(""))
}

func TestGetMarketName(t *testing.T) {
	should := require.New(t)

	should.Equal("HK", GetMarketName(1))
	should.Equal("Unknown", GetMarketName(0))
}

func TestGetTrdMarketID(t *testing.T) {
	should := require.New(t)

	should.Equal(int32(1), GetTrdMarketID("HK"))
	should.Equal(int32(0), GetTrdMarketID("XX"))
	should.Equal(int32(0), GetTrdMarketID(""))
}
