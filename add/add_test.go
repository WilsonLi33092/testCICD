package Add 

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2,7)
	expected := 9

	if sum != expected {
		t.Fatalf("expected %d but got %d", expected, sum)
	}
}