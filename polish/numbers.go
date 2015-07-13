package polish

import (
	"fmt"
	"math/rand"
	"strings"
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

type NounCase int

var CaseNames = []string{
	"nominative",
	"accusative",
	"genitive",
	"locative",
	"dative",
	"instrumental",
	"vocative",
}

var Months = map[NounCase][]string{
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

var NumberCardinalsZeroToTwenty = map[NounCase][]string{
	Nominative: []string{
		"zero",
		"jeden",
		"dwa",
		"trzy",
		"cztery",
		"pięć",
		"sześć",
		"siedem",
		"osiem",
		"dziewięć",
		"dziesięć",
		"jedenaście",
		"dwanaście",
		"trzynaście",
		"czternaście",
		"piętnaście",
		"szesnaście",
		"siedemnaście",
		"osiemnaście",
		"dziewięćnaście",
	},
	Genitive: []string{
		"zero",
		"jednego",
		"dwóch",
		"trzech",
		"czterech",
		"pięciu",
		"sześciu",
		"siedmiu",
		"ośmiu",
		"dziewięciu",
		"dwunastu",
		"trzynastu",
		"czternastu",
		"piętnastu",
		"szesnastu",
		"siedemnastu",
		"osiemnastu",
		"dziewięcnastu",
	},
}

var NumberCardinalsTens = map[NounCase][]string{
	Nominative: []string{
		"zero",
		"dziesięc",
		"dwadzieścia",
		"trzydzieści",
		"czterdzieści",
		"pięćdziesiąt",
		"sześćdziesiąt",
		"siedemdziesiąt",
		"osiemdziesiąt",
		"dziewięćdziesiąt",
	},
	Genitive: []string{
		"zero",
		"diesięciu",
		"dwudziestu",
		"trzydziestu",
		"czterdziestu",
		"pięćdziesiątego",
		"sześćdziesiątego",
		"siedemdziesiątego",
		"osiemdziesiątego",
		"dziewięćdziesiątego",
	},
}

var NumberCardinalsHundreds = map[NounCase][]string{
	Nominative: []string{
		"zero",
		"sto",
		"dwieście",
		"trzysta",
		"czterysta",
		"pięcset",
		"sześcset",
		"siedemset",
		"osiemset",
		"dziewięćset",
	},
}

// Dates by themselves use the masculine ordinal number
var DateOrdinals = map[NounCase][]string{
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

func GetPolishYear(year int, standAlone bool) (string, error) {
	if year < 1 || year > 2999 {
		return "", fmt.Errorf("Bad year given: %d", year)
	}

	var yearCase NounCase

	if standAlone {
		yearCase = Nominative
	} else {
		yearCase = Genitive
	}

	parts := make([]string, 0, 8)

	// Two thousand = special case
	if year == 2000 {
		return "dwutysięczny", nil
	}

	// Append thousands (always nominative)
	if year > 2000 {
		parts = append(parts, "dwa tysiąc")
	} else if year > 1000 {
		parts = append(parts, "tysiąc")
	}

	// Append century (always nominative)
	century := (year % 1000) / 100
	if century > 0 {
		parts = append(parts, NumberCardinalsHundreds[Nominative][century])

	}

	decade := (year % 100) / 10
	if decade > 1 {
		parts = append(parts, NumberCardinalsTens[yearCase][decade])
	}

	yearInDecade := year % 10
	if yearInDecade > 0 || decade == 1 {
		finalYear := yearInDecade
		if decade == 1 {
			finalYear += decade * 10
		}

		parts = append(parts, NumberCardinalsZeroToTwenty[yearCase][finalYear])
	}

	return strings.Join(parts, " "), nil
}

func GetCaseName(nounCase NounCase) string {
	return CaseNames[nounCase]
}

func GetPolishDate(day int, month int, standAlone bool) (string, error) {
	var dayCase NounCase
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

	utc, _ := time.LoadLocation("UTC")

	// 2016 is a leap year
	startDate := time.Date(2016, 1, 1, 0, 0, 0, 0, utc)
	dayAppend := rand.Int() % 366
	finalDate := startDate.AddDate(0, 0, dayAppend)

	const layout = "_2 January"
	dateString, _ := GetPolishDate(finalDate.Day(), int(finalDate.Month()), false)
	return finalDate.Format(layout), dateString
}
