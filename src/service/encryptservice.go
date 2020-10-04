package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func encrypt(chain string) (string, error) {
	block, err := aes.NewCipher([]byte(conf.Key))
	if err != nil {
		return "", err
	}

	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	content := pkcs5padding([]byte(chain), block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return string(crypted), nil
}

func decrypt(crypted string) (string, error) {
	block, err := aes.NewCipher([]byte(conf.Key))
	if err != nil {
		return "", err
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	decrypted := make([]byte, len(crypted))
	ecb.CryptBlocks(decrypted, []byte(crypted))

	return string(pkcs5trimming(decrypted)), nil
}

func pkcs5padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
