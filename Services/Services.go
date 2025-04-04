package Services

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"log"
	"net/http"
	"strings"
	"testContext/Structures"
	"time"
)

type value struct {
	StrValue   string
	FloatValue float32
}

func GetCodes(UrlGetCodes string) *Structures.Valuta {
	var codes Structures.Valuta
	rep := GetHttp(UrlGetCodes)
	Decoding(rep, &codes)
	defer rep.Body.Close()
	return &codes
}

// можно было использовать мапу для более быстрого поиска O(1), но решил сделать так, потому что данный код более читаемый

func CodeFinder(mstr *Structures.Valuta, actuallyName string) string {
	for _, v := range mstr.Items {
		// Пришлось использовать трим потому что коды шли с пробелом в конце....
		if strings.ToLower(v.Name) == strings.ToLower(actuallyName) {
			return strings.Trim(v.ParentCode, " ")
		}
	}
	log.Fatalln("Такого названия не существует. Попробуй посмотреть полынй перечень на: https://www.cbr.ru/scripts/XML_val.asp.")
	return ""
}

func GetHttp(url string) *http.Response {
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	client := &http.Client{
		Timeout: 30 * time.Second, // Добавляем таймаут
	}

	rep, _ := client.Do(request)

	return rep
}

// Т.к. ЦБ использует отличную кодировку от стандартного юникода необходимо создать кодировку, отличную от обычной маршаллы
// для этого был вынужден использовать стороннюю библиотеку charset, чтобы поменять кодировку

func Decoding(rep *http.Response, strct interface{}) {
	//defer rep.Body.Close()
	decoder := xml.NewDecoder(rep.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&strct); err != nil {
		log.Fatalln("Decoding", err)
	}
}
