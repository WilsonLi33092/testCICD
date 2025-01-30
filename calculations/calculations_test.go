package calculations 

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(2,3)
	expected := 6
	if result != expected {
		t.Fatalf("Expected %d, but got %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(4,2)
	expected := 2
	if result != expected {
		t.Fatalf("Expected %d, but got %d", expected, result)
	}

}

func TestMod(t *testing.T) {
	result := Mod(3,2)
	expected := 1
	if result != expected {
		t.Fatalf("Expected %d, but got %d", expected, result)
	}
}