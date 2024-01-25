package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Envelope struct {
	XMLName xml.Name   `xml:"Envelope"`
	Cube    CubeLevel1 `xml:"Cube"`
}

type CubeLevel1 struct {
	XMLName xml.Name     `xml:"Cube"`
	Cube    []CubeLevel2 `xml:"Cube"`
}

type CubeLevel2 struct {
	XMLName xml.Name     `xml:"Cube"`
	Time    string       `xml:"time,attr"`
	Cube    []CubeDetail `xml:"Cube"`
}

type CubeDetail struct {
	XMLName  xml.Name `xml:"Cube"`
	Currency string   `xml:"currency,attr"`
	Rate     string   `xml:"rate,attr"`
}

func main() {
	var (
		url  string
		err  error
		resp *http.Response
		data Envelope
	)

	url = fmt.Sprintf(`https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml`)
	resp, err = http.Get(url)
	if err != nil {
		log.Println(fmt.Sprintf(`error: `), err)
	}

	defer resp.Body.Close()

	if err = xml.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(fmt.Sprintf(`error: `), err)
	}
}
