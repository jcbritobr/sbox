package sbox

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

const (
	KeySize   = 32
	NonceSize = 24
)

var (
	ErrSeal = errors.New("secret: encryption failed")
	ErrOpen = errors.New("secret: decryption failed")
)

func GenNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, err
	}

	return nonce, nil
}

func Seal(key *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenNonce()
	if err != nil {
		return nil, ErrSeal
	}

	out := make([]byte, NonceSize)
	copy(out, nonce[:])
	out = secretbox.Seal(out, message, nonce, key)
	return out, nil
}

func Open(key *[KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (NonceSize + secretbox.Overhead) {
		return nil, ErrOpen
	}

	var nonce [NonceSize]byte
	copy(nonce[:], message[:NonceSize])
	out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
	if !ok {
		return nil, ErrOpen
	}

	return out, nil
}
