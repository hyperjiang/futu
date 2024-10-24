package infra_test

import (
	"testing"

	"github.com/hyperjiang/futu/infra"
	"github.com/stretchr/testify/require"
)

func TestCrypto(t *testing.T) {
	should := require.New(t)

	key := []byte("3FA037BF519D18D5")
	c, err := infra.NewCrypto(key, nil)
	should.NoError(err)

	data := []byte("hello, world")
	ciphertext := c.Encrypt(data)
	plaintext := c.Decrypt(ciphertext)
	should.Equal(data, plaintext)

	_, err = infra.NewCrypto(nil, nil)
	should.Error(err)

	should.Nil(c.Decrypt(nil))
	ciphertext[len(ciphertext)-1] = 15
	should.Nil(c.Decrypt(ciphertext))
}
