package client_test

import (
	"testing"

	"github.com/hyperjiang/futu/client"
	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	should := require.New(t)

	// default options
	opts := client.NewOptions()

	should.Equal(":11111", opts.Addr)
	should.Equal("futu-go", opts.ID)
	should.Nil(opts.PrivateKey)
	should.Nil(opts.PublicKey)
	should.True(opts.RecvNotify)
	should.Equal(100, opts.ResChanSize)

	// override default options

	opts2 := client.NewOptions(
		client.WithID("abc"),
		client.WithAddr(":8080"),
		client.WithPrivateKey([]byte("123")),
		client.WithPublicKey([]byte("123")),
		client.WithRecvNotify(false),
		client.WithResChanSize(10),
	)

	should.Equal("abc", opts2.ID)
	should.Equal(":8080", opts2.Addr)
	should.Equal([]byte("123"), opts2.PrivateKey)
	should.Equal([]byte("123"), opts2.PublicKey)
	should.False(opts2.RecvNotify)
	should.Equal(10, opts2.ResChanSize)
}
