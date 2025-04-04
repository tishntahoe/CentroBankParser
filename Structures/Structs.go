package Structures

import (
	"encoding/xml"
	"log"
	"strconv"
	"strings"
)

var (
	// нашел список названий валют и их коды. С ними проще работать
	UrlGetCodes = "https://www.cbr.ru/scripts/XML_val.asp"
	Url         = "https://www.cbr.ru/scripts/XML_dynamic.asp?date_req1=[datefrom]&date_req2=[dateto]&VAL_NM_RQ=[code]"
)

// Базовый структуры

type ValCurs struct {
	XMLName xml.Name  `xml:"ValCurs"`
	Date    string    `xml:"Date,attr"`
	Valutes []Valutes `xml:"Record"`
}

type Valutes struct {
	Name       string `xml:"Name"`
	Code       string `xml:"Id,attr"`
	Value      string `xml:"Value"`
	ValueFloat float64
	Date       string `xml:"Date,attr"`
}

// Получение кодов

type Valuta struct {
	XMLName xml.Name `xml:"Valuta"`
	Items   []Item   `xml:"Item"`
}

type Item struct {
	Name       string `xml:"Name"`
	ParentCode string `xml:"ParentCode"`
}

func (st *ValCurs) ConvFloatString() {
	for i, v := range st.Valutes {
		digit, err := strconv.ParseFloat(strings.ReplaceAll(v.Value, ",", "."), 32)
		if err != nil {
			log.Fatalln("Проблема с конвертацией: ", err)
		}
		st.Valutes[i].ValueFloat = digit
	}
}
