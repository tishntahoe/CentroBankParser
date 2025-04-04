package main

import (
	"flag"
	"fmt"
	"strings"
	"testContext/Services"
	"testContext/Structures"
	"time"
)

func main() {
	var (
		valstruct Structures.ValCurs
		dateTo    = time.Now().Format("02/01/2006")
		dateFrom  = time.Now().AddDate(0, 0, -90).Format("02/01/2006")
	)
	AllCodes := Services.GetCodes(Structures.UrlGetCodes)

	// не использовал библиотеки для построения консольного приложения, потому что не вижу смысла))
	help := flag.Bool("help", false, "Помощь по программе")
	name := flag.String("name", "", "Название валюты (пример: Австрийский шиллинг, Азербайджанский манат и тд. Полный перечень можно посмотреть на: https://www.cbr.ru/scripts/XML_val.asp)")
	type_find := flag.String("type", "avg", "Тип результата. Всего их 4: max, min, avg, ")
	flag.Parse()

	if *help {
		fmt.Println("Использование:")
		fmt.Println("--help          Показать справку")
		fmt.Println("--name [value]  Название валюты (пример: Австрийский шиллинг, Азербайджанский манат и тд. Полный перечень можно посмотреть на: https://www.cbr.ru/scripts/XML_val.asp)")
		fmt.Println("--name [value] --type [value]  Тип выполняемой задачи. MAX, MIN, AVG(среднее значение курса рубля к name [value] валюте)")
		fmt.Println("--type [value]  Тип выполняемой задачи. AVGRUB (среднее значение курса рубля за весь период по всем валютам) используется без --name. В случае если будет использоваться эта команда - name будет игнорироваться")
		return
	}

	if strings.ToLower(*type_find) == "avgrub" {
		result := Services.AvgRUB(AllCodes, dateFrom, dateTo)
		fmt.Println("Среднее значение по всем валютам: " + fmt.Sprint(result))
		return
	} else if *type_find == "" {
		fmt.Println("Вы не указали тип выполняемой задачи (max,min,avg)")
		fmt.Println("Полная справка --help")
	}
	if *name == "" {
		fmt.Println("Вы не указали имя")
		fmt.Println("Полная справка --help")
	}

	code := Services.CodeFinder(AllCodes, *name)

	replacer := strings.NewReplacer(
		"[datefrom]", dateFrom,
		"[dateto]", dateTo,
		"[code]", code,
	)
	url := replacer.Replace(Structures.Url)
	rep := Services.GetHttp(url)

	Services.Decoding(rep, &valstruct)
	rep.Body.Close()

	valstruct.ConvFloatString()

	switch strings.ToLower(*type_find) {
	case "max":
		result, date := Services.Max(&valstruct)
		fmt.Println("Результат Maximum функции: " + fmt.Sprint(result))
		fmt.Println("Название валюты: " + *name + ". Дата: " + date)
	case "min":
		result, date := Services.Min(&valstruct)
		fmt.Println("Результат Minimal функции: " + fmt.Sprint(result))
		fmt.Println("Название валюты: " + *name + ". Дата: " + date)
	case "avg":
		result := Services.Avg(&valstruct)
		fmt.Println("Результат Average функции: " + fmt.Sprint(result))

	default:
		fmt.Println("Функции с названием " + *type_find + " не существует. --help полная справка.")
	}
}
