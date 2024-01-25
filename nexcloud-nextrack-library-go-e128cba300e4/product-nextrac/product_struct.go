package product_nextrac

type ProductConfiguration struct {
	SignatureKey     string `json:"signatureKey"`
	ProductSignature string `json:"productSignature"`
	ClientId         string `json:"clientId"`
	ClientSecret     string `json:"clientSecret"`
	EncryptKey       string `json:"encryptKey"`
	HardwareId       string `json:"hardwareId"`
	ProductKey       string `json:"productKey"`
	ProductId        string `json:"productId"`
}

type LicenseConfiguration struct {
	InstallationId     int64       `json:"installationId"`
	ClientId           string      `json:"clientId"`
	ProductId          string      `json:"productId"`
	LicenseVariantName string      `json:"licenseVariantName"`
	LicenseTypeName    string      `json:"licenseTypeName"`
	DeploymentMethod   string      `json:"deploymentMethod"`
	NoOfUser           int64       `json:"noOfUser"`
	UniqueId1          string      `json:"uniqueId1"`
	UniqueId2          string      `json:"uniqueId2"`
	LicenseStatus      int64       `json:"licenseStatus"`
	ModuleName1        string      `json:"moduleName1"`
	ModuleName2        string      `json:"moduleName2"`
	ModuleName3        string      `json:"moduleName3"`
	ModuleName4        string      `json:"moduleName4"`
	ModuleName5        string      `json:"moduleName5"`
	ModuleName6        string      `json:"moduleName6"`
	ModuleName7        string      `json:"moduleName7"`
	ModuleName8        string      `json:"moduleName8"`
	ModuleName9        string      `json:"moduleName9"`
	ModuleName10       string      `json:"moduleName10"`
	MaxOfflineDays     int64       `json:"maxOfflineDays"`
	ProductValidFrom   string      `json:"productValidFrom"`
	ProductValidThru   string      `json:"productValidThru"`
	Component          []Component `json:"component"`
}

type ValidationResponse struct {
	MessageCode          string
	Message              string
	Notification         string
	Configuration        LicenseConfiguration
	ProductConfiguration ProductConfiguration

	ProductSignature string
	ProductEncrypt   string
	ProductKey       string
}

type Component struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
