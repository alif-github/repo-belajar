package product_nextrac

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestGenerateSecretKey(t *testing.T) {
	productKeyAndProductEncrypt := "NexsoftProductKeyAndProductEncrypt"
	signatureKey := "NexsoftSignatureKey"
	salt := "12345678"

	key, _ := secretKey(signatureKey, []byte(salt))
	fmt.Println("SecretKey:\t", base64.StdEncoding.EncodeToString(key))

	encryptedData := AESEncrypt(productKeyAndProductEncrypt, key)
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)

	fmt.Println("Encode Result:\t", encryptedString)

	fmt.Println("Decode Result:\t", AESDecrypt(encryptedString, key))

}

func TestEncodeInGo(t *testing.T) {
	signatureKey := "NexsoftSignatureKey"
	encode := base64.StdEncoding.EncodeToString([]byte(signatureKey))

	fmt.Println(encode)
}

func TestProductEncrypt(t *testing.T) {

	productConfig := ProductConfiguration{
		SignatureKey:     "signatureKey",
		ProductSignature: "harusnya diremove ga guna ini",
		ClientId:         "1",
		ClientSecret:     "1",
		EncryptKey:       "1",
		HardwareId:       "1",
		ProductId:        "1",
	}

	licenseConfig := LicenseConfiguration{
		InstallationId:     1,
		ClientId:           productConfig.ClientId,
		ProductId:          productConfig.ProductId,
		LicenseVariantName: "1",
		LicenseTypeName:    "1",
		DeploymentMethod:   "C",
		NoOfUser:           1,
		UniqueId1:          "1",
		UniqueId2:          "1",
		LicenseStatus:      1,
		MaxOfflineDays:     10,
	}
	response := ProductEncrypt(licenseConfig, productConfig)

	fmt.Println("response.ProductEncrypt :\t", response.ProductEncrypt)
	fmt.Println("response.ProductSignature :\t", response.ProductSignature)

}

func TestProductDecrypt(t *testing.T) {
	productConfig := ProductConfiguration{
		SignatureKey:     "signatureKey",
		ProductSignature: "harusnya diremove ga guna ini",
		ClientId:         "1",
		ClientSecret:     "1",
		EncryptKey:       "1",
		HardwareId:       "1",
		ProductId:        "1",
	}

	licenseConfig := LicenseConfiguration{
		InstallationId:     1,
		ClientId:           productConfig.ClientId,
		ProductId:          productConfig.ProductId,
		LicenseVariantName: "1",
		LicenseTypeName:    "1",
		DeploymentMethod:   "C",
		NoOfUser:           1,
		UniqueId1:          "1",
		UniqueId2:          "1",
		LicenseStatus:      1,
		MaxOfflineDays:     10,
	}
	encrypt := ProductEncrypt(licenseConfig, productConfig)

	response := ProductDecrypt(encrypt.ProductSignature, productConfig)
	fmt.Println("response.ProductEncrypt :\t", response)
}
