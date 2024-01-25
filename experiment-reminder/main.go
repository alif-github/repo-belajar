package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name  string
	Class string
}

func main() {
	var userNew user
	var userTemp = user{
		Name:  "Udin",
		Class: "12",
	}
	bytes, _ := json.Marshal(userTemp)
	fmt.Println(string(bytes))

	variabel := string(bytes)

	_ = json.Unmarshal([]byte(variabel), &userNew)
	fmt.Println(userNew.Name)
}
