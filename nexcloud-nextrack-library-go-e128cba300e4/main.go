package main

import (
	"encoding/json"
	"fmt"
	productnextrac "nextrack_library_go/product-nextrac"
	"os/exec"
)

func main() {
	licenseText := productnextrac.LicenseConfiguration{
		InstallationId:     1,
		ClientId:           "558444dc42e548848b6c034e1c29e5a1",
		ProductId:          "ND610",
		LicenseVariantName: "premium",
		LicenseTypeName:    "full",
		DeploymentMethod:   "C",
		NoOfUser:           10,
		UniqueId1:          "KJGJDH",
		UniqueId2:          "UYRTJH",
		LicenseStatus:      1,
		ModuleName1:        "cashbank",
		ModuleName2:        "",
		ModuleName3:        "",
		ModuleName4:        "",
		ModuleName5:        "",
		ModuleName6:        "",
		ModuleName7:        "",
		ModuleName8:        "",
		ModuleName9:        "",
		ModuleName10:       "",
		MaxOfflineDays:     10,
		ProductValidFrom:   "validfrom",
		ProductValidThru:   "validuntil",
	}

	licenseText.Component = []productnextrac.Component{
		{
			Name:  "a",
			Value: "a",
		},
		{
			Name:  "b",
			Value: "b",
		},
	}

	strLicenseText, _ := json.Marshal(licenseText)

	productText := productnextrac.ProductConfiguration{
		SignatureKey: "abc",
		ClientSecret: "bgd",
		EncryptKey:   "asdaw",
		HardwareId:   "ghj",
	}

	strProductText, _ := json.Marshal(productText)

	path := fmt.Sprintf(`C:\repo-belajar\repo-eksperimen-code\src\nexcloud-nextrack-library-go-e128cba300e4\generator.exe`)
	var args map[string]string

	args = make(map[string]string)
	args["args1"] = "ProductEncrypt"
	args["args2"] = string(strLicenseText)
	args["args3"] = string(strProductText)

	cmd := exec.Command(path, args["args1"], args["args2"], args["args3"])
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	println(string(strLicenseText))
	println("****************************************************************************************************************************************************************************************************")
	print(string(out))
	println("****************************************************************************************************************************************************************************************************")

	var response productnextrac.ValidationResponse
	_ = json.Unmarshal(out, &response)

	var productConfigurationDecrypt productnextrac.ProductConfiguration
	productConfigurationDecrypt = productnextrac.ProductConfiguration{
		SignatureKey:     productText.SignatureKey,  //di dapat dari request
		ProductSignature: response.ProductSignature, //di dapat dari response ~
		ClientId:         licenseText.ClientId,      //di dapat dari request
		ClientSecret:     productText.ClientSecret,  //di dapat dari request
		EncryptKey:       productText.EncryptKey,    //di dapat dari request
		HardwareId:       productText.HardwareId,    //di dapat dari request
		ProductKey:       response.ProductKey,       //di dapat dari response ~
		ProductId:        licenseText.ProductId,     //di dapat dari request
	}

	strProductLicenseDecrypt, _ := json.Marshal(productConfigurationDecrypt)

	var argsDecrypt map[string]string
	argsDecrypt = make(map[string]string)
	argsDecrypt["args1"] = "ProductDecrypt"
	argsDecrypt["args2"] = response.ProductEncrypt
	argsDecrypt["args3"] = string(strProductLicenseDecrypt)

	cmdDecrypt := exec.Command(path, argsDecrypt["args1"], argsDecrypt["args2"], argsDecrypt["args3"])
	out2, errs := cmdDecrypt.Output()
	if errs != nil {
		println(errs.Error())
		return
	}

	print(string(out2))
	println("****************************************************************************************************************************************************************************************************")
}
