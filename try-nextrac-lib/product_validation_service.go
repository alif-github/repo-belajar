package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func secretKey(password string, salt []byte) ([]byte, []byte) {
	return pbkdf2.Key([]byte(password), salt, 1000, 16, sha1.New), salt
}

func AESEncrypt(plainText string, key []byte, _ ProductConfiguration) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error while generate NewCipher", err)
		return nil
	}
	if plainText == "" {
		fmt.Println("plainText content empty")
		return nil
	}

	// Todo productKey used as IV(initial vector) but should 16 byte
	ecb := cipher.NewCBCEncrypter(block, []byte("1234567890123456"))
	content := []byte(plainText)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(cipherText string, key []byte) (decryptedString string) {
	cipherTextDecoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	// Todo productKey used as IV(initial vector) but should 16 byte
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, []byte("1234567890123456"))
	mode.CryptBlocks(cipherTextDecoded, cipherTextDecoded)
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
