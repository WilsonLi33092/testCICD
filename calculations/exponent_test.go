package calculations

import "testing"

func TestExponent(t *testing.T) {
	result := Exponent(2,3)
	expected := 8
	if result != expected {
		t.Fatalf("Expected %d, but got %d", expected, result)
	}
}