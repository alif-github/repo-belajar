package product_nextrac

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func secretKey(password string, salt []byte) ([]byte, []byte) {
	return pbkdf2.Key([]byte(password), salt, 1000, 16, sha1.New), salt
}

func AESEncrypt(plainText string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error while generate NewCipher", err)
	}
	if plainText == "" {
		fmt.Println("plainText content empty")
	}

	// todo  productKey used as IV(initial vector) but should 16 byte
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
	// todo  productKey used as IV(initial vector) but should 16 byte
	block, err := aes.NewCipher(key)
	if err != nil {
		// todo handle this error
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, []byte("1234567890123456"))
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func ProductEncrypt(licenseConfiguration LicenseConfiguration, productConfiguration ProductConfiguration) ValidationResponse {

	requestKey := productConfiguration.ClientId + "|" + productConfiguration.ClientSecret + "|" + productConfiguration.ProductKey + "|" + productConfiguration.EncryptKey + "|" + productConfiguration.HardwareId
	key, _ := secretKey(requestKey, []byte(productConfiguration.ProductKey))

	configJsonRequest, err := json.Marshal(licenseConfiguration)
	if err != nil {
		fmt.Println(err)
		return ValidationResponse{MessageCode: "FAILED",
			Message:              err.Error(),
			Notification:         "Terjadi Kesalahan dalam sistem, permintaan decrypt product encrypt gagal",
			Configuration:        licenseConfiguration,
			ProductConfiguration: productConfiguration,
			ProductSignature:     productConfiguration.ProductSignature,
			ProductKey:           productConfiguration.ProductKey,
		}
	}

	fmt.Println("JSON MARSHAL ", string(configJsonRequest))
	encryptedData := AESEncrypt(string(configJsonRequest), key)
	productKeyAndProductEncrypt := base64.StdEncoding.EncodeToString(encryptedData)

	return ValidationResponse{MessageCode: "SUCCESS",
		Message:              productKeyAndProductEncrypt,
		Notification:         "",
		ProductConfiguration: productConfiguration,
		ProductSignature:     productKeyAndProductEncrypt,
		ProductEncrypt:       productKeyAndProductEncrypt,
		ProductKey:           productConfiguration.ProductKey,
	}

}

func ProductDecrypt(decryptedText string, productConfiguration ProductConfiguration) ValidationResponse {
	requestKey := productConfiguration.ClientId + "|" + productConfiguration.ClientSecret + "|" + productConfiguration.ProductKey + "|" + productConfiguration.EncryptKey + "|" + productConfiguration.HardwareId
	key, _ := secretKey(requestKey, []byte(productConfiguration.ProductKey))
	productEncrypt := AESDecrypt(decryptedText, key)
	data := LicenseConfiguration{}
	err := json.Unmarshal([]byte(productEncrypt), &data)
	if err != nil {
		return ValidationResponse{MessageCode: "FAILED",
			Message:              err.Error(),
			Notification:         "Terjadi Kesalahan dalam sistem, permintaan decrypt product encrypt gagal",
			ProductConfiguration: productConfiguration,
			ProductSignature:     productConfiguration.ProductSignature,
			ProductKey:           productConfiguration.ProductKey,
		}
	}
	return ValidationResponse{MessageCode: "SUCCESS",
		Notification:  "",
		Configuration: data,
		ProductKey:    productConfiguration.ProductKey,
	}

}
