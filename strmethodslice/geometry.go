package geometry

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimiter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}
