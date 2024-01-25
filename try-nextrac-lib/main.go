package main

import "fmt"

type School struct {
	Name  string
	Class string
}

func main() {
	var schoolData School
	coba(&schoolData)
	fmt.Println(schoolData)

	check := schoolData
	check.Name = "Budi"
	fmt.Println(schoolData)
	fmt.Println(check)

	//plainText := "SemangatMenjalankanHari"
	//signatureKey := "Signature"
	//salt := "Salt"
	//
	//product := ProductConfiguration{ProductKey: "ProductKey"}
	//
	//key, _ := secretKey(signatureKey, []byte(salt))
	//
	////--------- Encrypt
	//encrypt := AESEncrypt(plainText, key, product)
	//fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	//
	////--------- Decrypt
	//fmt.Println(AESDecrypt(base64.StdEncoding.EncodeToString(encrypt), key))
}

func coba(data *School) {
	*data = School{
		Name:  "udin",
		Class: "13",
	}
}
