package geometry

import "testing"

func TestPerimiter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimiter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		got := rect.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %v want %v\n", got, want)
		}
	})
}
