package polish

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Nominative = iota
	Accusative
	Genitive
	Locative
	Dative
	Instrumental
	Vocative
)

var CaseNames = []string{
	"nominative",
	"accusative",
	"genitive",
	"locative",
	"dative",
	"instrumental",
	"vocative",
}

var Months = map[int][]string{
	Nominative: []string{
		"styczeń",
		"luty",
		"marzec",
		"kwiecień",
		"maj",
		"czerwiec",
		"lipiec",
		"sierpień",
		"wrzesień",
		"październik",
		"listopad",
		"grudzień",
	},

	Genitive: []string{
		"stycznia",
		"lutego",
		"marcu",
		"kwietnia",
		"maja",
		"czerwca",
		"lipca",
		"sierpnia",
		"września",
		"października",
		"listopada",
		"grudnia",
	},

	Instrumental: []string{
		"styczniu",
		"lutym",
		"marcu",
		"kwietniu",
		"maju",
		"czerwcu",
		"wrześniu",
		"październik",
		"listopadzie",
		"grudnie",
	},
}

// Dates by themselves use the masculine ordinal number
var DateOrdinals = map[int][]string{
	Nominative: []string{
		"zero",
		"pierwszy",
		"drugi",
		"trzeci",
		"czwarty",
		"piąty",
		"szósty",
		"siódmy",
		"ósmy",
		"dziewiąty",
		"dziesiąty",
		"jedenasty",
		"dwunasty",
		"trzynasty",
		"czternasty",
		"piętnasty",
		"szesnasty",
		"siedemnasty",
		"osiemnasty",
		"dziewiętnasty",
		"dwudziesty",
		"dwudziesty pierwszy",
		"dwudziesty drugi",
		"dwudziesty trzeci",
		"dwudziesty czwarty",
		"dwudziesty piąty",
		"dwudziesty szósty",
		"dwudziesty siódmy",
		"dwudziesty ósmy",
		"dwudziesty dziewiąty",
		"trzydziesty",
		"trzydziesty pierwszy",
	},
	Genitive: []string{
		"zero",
		"pierwszego",
		"drugiego",
		"trzeciego",
		"czwartego",
		"piątego",
		"szostego",
		"siódmego",
		"ósmego",
		"dziewiątego",
		"dziesiątego",
		"jedenastego",
		"dwunastego",
		"trzynastego",
		"czternastego",
		"piątnastego",
		"szesnastego",
		"siedemnastego",
		"osiemnastego",
		"dziewiętnatego",
		"dwudziestego",
		"dwudziestego pierwszego",
		"dwudziestego drugiego",
		"dwudziestego trzeciego",
		"dwudziestego czwartego",
		"dwudziestego piątego",
		"dwudziestego szóstego",
		"dwudziestego siódmego",
		"dwudziesty ósmego",
		"dwudziesty dziewiątego",
		"trzydziestego",
		"trzydziestego pierwszego",
	},
}

func GetCaseName(nounCase int) string {
	return CaseNames[nounCase]
}

func GetPolishDate(day int, month int, standAlone bool) (string, error) {
	var dayCase int
	if standAlone {
		dayCase = Nominative
	} else {
		dayCase = Genitive
	}
	dayString := DateOrdinals[dayCase][day]
	monthString := Months[Genitive][month-1]
	dateString := fmt.Sprintf("%s %s", dayString, monthString)

	return dateString, nil
}

func DateQuiz() (string, string) {
	// TODO: if options is nil, allow randoms

	utc, _ := time.LoadLocation("utc")

	// 2016 is a leap year
	startDate := time.Date(2016, 1, 1, 0, 0, 0, 0, utc)
	dayAppend := rand.Int() % 366
	finalDate := startDate.AddDate(0, 0, dayAppend)

	const layout = "_2 January"
	dateString, _ := GetPolishDate(finalDate.Day(), int(finalDate.Month()), false)
	return finalDate.Format(layout), dateString
}
