package integers

import "testing"

func TestHelloWorld(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, sum)
	}
}
