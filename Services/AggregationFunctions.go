package Services

import (
	"strings"
	"testContext/Structures"
)

// Во всех функциях можно было использовать параллелизм, но именно в этом таске смысла в этом нет, т.к. очень мало данных
// решил просто использовать базовые алгоритмы

// Решил сделать поиск O(n/2). Немного устал от линейного поиска.
// В конце отдельным выражением работаем с последним элементом, в случае если число возможных итераций нечетно

func Max(curs *Structures.ValCurs) (float64, string) {
	if len(curs.Valutes) == 0 {
		return 0, ""
	}
	res := curs.Valutes[0].ValueFloat
	var actDate string

	for i := 0; i < len(curs.Valutes)-1; i += 2 {
		if curs.Valutes[i].ValueFloat > res {
			res = curs.Valutes[i].ValueFloat
			actDate = curs.Valutes[i].Date
		}
		if curs.Valutes[i+1].ValueFloat > res {
			res = curs.Valutes[i+1].ValueFloat
			actDate = curs.Valutes[i].Date
		}
	}

	if len(curs.Valutes)%2 == 1 {
		last := curs.Valutes[len(curs.Valutes)-1].ValueFloat
		if last > res {
			res = last
			actDate = curs.Valutes[len(curs.Valutes)-1].Date
		}
	}
	return res, actDate
}

// Просто меняем знаки

func Min(curs *Structures.ValCurs) (float64, string) {
	if len(curs.Valutes) == 0 {
		return 0, ""
	}
	res := curs.Valutes[0].ValueFloat
	var actDate string
	for i := 0; i < len(curs.Valutes)-1; i += 2 {
		if curs.Valutes[i].ValueFloat < res {
			res = curs.Valutes[i].ValueFloat
			actDate = curs.Valutes[i].Date
		}
		if curs.Valutes[i+1].ValueFloat < res {
			res = curs.Valutes[i+1].ValueFloat
			actDate = curs.Valutes[i].Date
		}
	}

	if len(curs.Valutes)%2 == 1 {
		last := curs.Valutes[len(curs.Valutes)-1].ValueFloat
		if last < res {
			res = last
			actDate = curs.Valutes[len(curs.Valutes)-1].Date
		}
	}
	return res, actDate
}

func Avg(curs *Structures.ValCurs) float64 {

	var (
		res   float64
		count float64
	)
	if len(curs.Valutes) == 0 {
		return 0
	}
	for _, v := range curs.Valutes {
		count++
		res += v.ValueFloat
	}
	return res / count
}

func AvgRUB(val *Structures.Valuta, df string, dt string) float64 {
	var (
		vc    Structures.ValCurs
		res   float64
		count float64
	)

	for _, v := range val.Items {

		replacer := strings.NewReplacer(
			"[datefrom]", df,
			"[dateto]", dt,
			"[code]", strings.Trim(v.ParentCode, " "),
		)
		url := replacer.Replace(Structures.Url)
		rep := GetHttp(url)
		Decoding(rep, &vc)

		rep.Body.Close()
		vc.ConvFloatString()
		res += Avg(&vc)
		count++

	}

	return res / count
}
