package polish

import (
	"testing"
)

func TestGetCaseName(t *testing.T) {
	if GetCaseName(Instrumental) != "instrumental" {
		t.Fatalf("Expected 'instrumental', got '%s'\n", GetCaseName(Instrumental))
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
