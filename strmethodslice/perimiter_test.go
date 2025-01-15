package perimiter

import "testing"

func TestPerimiter(t *testing.T) {
	got := Perimiter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
