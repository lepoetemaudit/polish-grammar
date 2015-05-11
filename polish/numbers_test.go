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
		day   int
		month int
	}

	var dates = map[DayMonth]string{
		DayMonth{1, 5}:   "pierwszy maja",
		DayMonth{27, 1}:  "dwudziesty siódmy stycznia",
		DayMonth{15, 12}: "piętnasty grudnia",
	}

	for date, expected := range dates {
		day := date.day
		month := date.month
		str, _ := GetPolishDate(day, month, true)
		if str != expected {
			t.Fatalf(str)
		}
	}

}
