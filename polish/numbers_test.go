package polish

import (
	"testing"
)

func TestGetCaseName(t *testing.T) {
	if GetCaseName(Instrumental) != "instrumental" {
		t.Fatalf("Expected 'instrumental', got '%s'\n", GetCaseName(Instrumental))
	}
}

func test_polish_date_format() {

}
