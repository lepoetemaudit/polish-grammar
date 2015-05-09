package polish

import (
	"fmt"
	"time"
)

type DateOptions struct {
	toPolish bool
}

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

func GetCaseName(nounCase int) string {
	return CaseNames[nounCase]
}

func GetPolishDate(date time.Time, nounCase int) (string, error) {
	supportedCases := []int{Nominative, Genitive, Instrumental}
	supported := false

	for _, c := range supportedCases {
		if nounCase == c {
			supported = true
			break
		}
	}

	if !supported {
		return "", fmt.Errorf("Unsupported number case: %s", GetCaseName(nounCase))
	}

	return "nope", nil
}

func DateQuiz(options DateOptions) (string, string) {
	// TODO: if options is nil, allow randoms

	// For now, assume dates between 1000 and 2050
	return "nope", "nope"
}
