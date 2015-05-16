package polish

import (
	"testing"
)

func TestGetCaseName(t *testing.T) {
	if GetCaseName(Instrumental) != "instrumental" {
		t.Fatalf("Expected 'instrumental', got '%s'\n",
			GetCaseName(Instrumental))
	}
}

func TestGetPolishYear(t *testing.T) {
	var standAloneYears = map[int]string{
		1999: "tysiąc dziewięćset dziewięćdziesiąt dziewięć",
		50:   "pięćdziesiąt",
		1766: "tysiąc siedemset sześćdziesiąt sześć",
		1254: "tysiąc dwieście pięćdziesiąt cztery",
		2055: "dwa tysiąc pięćdziesiąt pięć",
		2008: "dwa tysiąc osiem",
		2000: "dwutysięczny",
	}

	for year, expected := range standAloneYears {
		if result, _ := GetPolishYear(year, true); result != expected {
			t.Fatalf("Expected '%s' but got '%s'", expected, result)
		}
	}
}

func TestGetPolishDate(t *testing.T) {

	type DayMonth struct {
		day        int
		month      int
		standAlone bool
	}

	var dates = map[DayMonth]string{
		// Standalone dates
		DayMonth{1, 5, true}:   "pierwszy maja",
		DayMonth{27, 1, true}:  "dwudziesty siódmy stycznia",
		DayMonth{15, 12, true}: "piętnasty grudnia",

		// Non-standalone dates
		DayMonth{13, 7, false}:  "trzynastego lipca",
		DayMonth{31, 10, false}: "trzydziestego pierwszego października",
		DayMonth{14, 2, false}:  "czternastego lutego",
	}

	for date, expected := range dates {
		day := date.day
		month := date.month
		standAlone := date.standAlone
		str, _ := GetPolishDate(day, month, standAlone)
		if str != expected {
			t.Fatalf("Expected '%s' but got: '%s'", expected, str)
		}
	}
}

func TestDateQuiz(t *testing.T) {
	_, _ = DateQuiz()
}
