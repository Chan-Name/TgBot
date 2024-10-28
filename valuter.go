package telegram

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  []struct {
		NumCode   string `xml:"NumCode"`
		CharCode  string `xml:"CharCode"`
		Nominal   string `xml:"Nominal"`
		Name      string `xml:"Name"`
		Value     string `xml:"Value"`
		VunitRate string `xml:"VunitRate"`
	} `xml:"Valute"`
}

func valuterStart() string {
	url := fmt.Sprintf("https://cbr.ru/scripts/XML_daily.asp?date_req=%v", currentDate())
	data := bytes.Buffer{}

	client := http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "curl/7.54.1")

	res, _ := client.Do(req)

	io.Copy(&data, res.Body)
	var valCurs ValCurs

	decoder := xml.NewDecoder(&data)
	decoder.CharsetReader = charset.NewReaderLabel

	decoder.Decode(&valCurs)
	var z string
	for _, valute := range valCurs.Valute {
		z += fmt.Sprintf("%s: 1 %s = %v\n", valute.CharCode, valute.Name, valute.Value)
	}
	return z
}

func currentDate() string {
	currentTime := time.Now()
	return currentTime.Format("02/01/2006")
}
