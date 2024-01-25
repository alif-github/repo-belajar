package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Member struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
}

type Audit struct {
	KeyId  string `json:"key_id"`
	KeyEn  string `json:"key_en"`
	Before string `json:"before"`
	After  string `json:"after"`
}

func main() {
	var (
		a     Member
		b     Member
		c     = fmt.Sprintf(`{"id":12,"first_name":"Aldi","last_name":"Lukito Diono","number":null}`)
		d     = fmt.Sprintf(`{"id":14,"first_name":"Aldi","last_name":null,"number":1}`)
		audit []Audit
	)

	_ = json.Unmarshal([]byte(c), &a)
	_ = json.Unmarshal([]byte(d), &b)

	lastName := checkAndCompareHistory("Nama Belakang", "Lastname", a.LastName, b.LastName)
	if lastName.KeyId != "" {
		audit = append(audit, lastName)
	}

	firstName := checkAndCompareHistory("Nama Depan", "Firstname", a.FirstName, b.FirstName)
	if firstName.KeyId != "" {
		audit = append(audit, firstName)
	}

	number := checkAndCompareHistory("Nomor", "Number", a.Number, b.Number)
	if number.KeyId != "" {
		audit = append(audit, number)
	}

	byt, ers := json.Marshal(audit)
	if ers != nil {
		fmt.Println("Error: ", ers.Error())
		return
	}

	fmt.Println(string(byt))
}

func checkAndCompareHistory(keyId, KeyEn string, before, after interface{}) Audit {
	var (
		recordBefore string
		recordAfter  string
	)

	recordBefore = parseInterfaceToString(before)
	recordAfter = parseInterfaceToString(after)

	if recordBefore == recordAfter {
		return Audit{}
	}

	return Audit{
		KeyId:  keyId,
		KeyEn:  KeyEn,
		Before: recordBefore,
		After:  recordAfter,
	}
}

func parseInterfaceToString(value interface{}) (record string) {
	switch value.(type) {
	case int64:
		num := value.(int64)
		record = strconv.Itoa(int(num))
	case string:
		record = value.(string)
	case bool:
		bol := value.(bool)
		record = strconv.FormatBool(bol)
	default:
	}

	return
}
