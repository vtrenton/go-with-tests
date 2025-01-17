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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
