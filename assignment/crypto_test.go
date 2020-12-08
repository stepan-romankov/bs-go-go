package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrypto(t *testing.T) {
	const key = "abcdefghijklmnopqrstuvwxyzABCDEF"
	crypt := NewCrypt(key)
	expectedPlain := "test"
	crypted, err := crypt.encrypt(expectedPlain)
	assert.Nil(t, err)
	plain, err := crypt.decrypt(crypted)
	assert.Nil(t, err)
	assert.Equal(t, expectedPlain, plain)
}
