package bitutil

import (
	"testing"
)

var hamPassingTests = []struct {
	a        string
	b        string
	expected int
}{
	{
		"this is a test",
		"wokka wokka!!!",
		37,
	},
	{
		"00000000 00000000 00000000 00000000 00000000 00000000",
		"11111111 11111111 11111111 11111111 11111111 11111111",
		48,
	},
}

var hamFailingTests = []struct {
	a string
	b string
}{
	{
		"this is a test",
		"that is supposed to fail",
	},
}

func TestHamming(t *testing.T) {
	for _, tt := range hamPassingTests {
		result, err := Hamming([]byte(tt.a), []byte(tt.b))
		if err != nil {
			t.Fatal(err)
		} else if result != tt.expected {
			t.Errorf("Unexpected value")
		}
	}

	for _, tt := range hamFailingTests {
		_, err := Hamming([]byte(tt.a), []byte(tt.b))
		if err == nil {
			t.Errorf("Expected error for strings of unequal lengths")
		}
	}
}
