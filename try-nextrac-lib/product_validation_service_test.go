package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestGenerateSecretKey(t *testing.T) {
	productKeyAndProductEncrypt := "NexsoftProductKeyAndProductEncrypt"
	signatureKey := "NexsoftSignatureKey"
	salt := "12345678"

	product := ProductConfiguration{ProductKey: "1234567890123456"}

	key, _ := secretKey(signatureKey, []byte(salt))
	fmt.Println("SecretKey:\t", base64.StdEncoding.EncodeToString(key))

	encryptedData := AESEncrypt(productKeyAndProductEncrypt, key, product)
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)

	fmt.Println("Encode Result:\t", encryptedString)

	fmt.Println("Decode Result:\t", AESDecrypt(encryptedString, key))
}

func TestEncodeInGo(t *testing.T) {
	signatureKey := "NexsoftSignatureKey"
	encode := base64.StdEncoding.EncodeToString([]byte(signatureKey))

	fmt.Println(encode)
}
