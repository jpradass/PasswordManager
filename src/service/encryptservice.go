package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func encrypt(chain string) ([]byte, error) {
	c, err := aes.NewCipher([]byte(conf.Key))
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
	return gcm.Seal(nonce, nonce, []byte(chain), nil), nil
}

func decrypt(crypted []byte) (string, error) {
	block, err := aes.NewCipher([]byte(conf.Key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(crypted) < nonceSize {
		return "", err
	}

	nonce, cryptedtext := crypted[:nonceSize], crypted[nonceSize:]
	text, err := gcm.Open(nil, nonce, cryptedtext, nil)
	if err != nil {
		return "", err
	}
	return string(text), nil
}
