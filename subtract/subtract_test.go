package subtract

import "testing"

func TestSubtract(t *testing.T) {
	result := Subtract(4,1)
	expected := 3
	if(result != expected) {
		t.Fatalf("expected %d, but got %d", expected, result)
	}
}