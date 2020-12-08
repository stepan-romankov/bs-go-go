package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type Crypt struct {
	secret []byte
}

func NewCrypt(secret string) Crypt {
	return Crypt{secret: []byte(secret)}
}

func (crypt *Crypt) encrypt(plaintext string) ([]byte, error) {
	c, err := aes.NewCipher(crypt.secret)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (crypt *Crypt) decrypt(ciphertext []byte) (string, error) {
	c, err := aes.NewCipher(crypt.secret)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
