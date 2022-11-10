package sbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	key = &[32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
)

func TestGenNonceMustHave24BytesAndErrorNil(t *testing.T) {
	nonce, err := GenNonce()
	println(nonce)
	assert.Nil(t, err)
	assert.Equal(t, NonceSize, len(nonce))
}

func TestSealOpenMessage(t *testing.T) {
	nonce, err := GenNonce()
	println(nonce)
	assert.Nil(t, err)
	message := "The quick brown fox jumps over the lazy dog"
	msgSealed, err := Seal(
		key,
		[]byte(message),
	)
	assert.Nil(t, err)
	assert.NotEqual(t, message, msgSealed)
	omessage, err := Open(key, msgSealed)
	assert.Nil(t, err)
	assert.Equal(t, message, string(omessage))
}
