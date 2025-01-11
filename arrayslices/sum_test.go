package arrslice

import "testing"

func TestSum(t *testing.T) {
	t.Run("Test slice with size of 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %d", got, want, numbers)
		}
	})

	//t.Run("Test slice with size of 3", func(t *testing.T) {
	//	numbers := []int{1, 2, 3}

	//	got := Sum(numbers)
	//	want := 6

	//	if got != want {
	//		t.Errorf("got %d want %d given, %d", got, want, numbers)
	//	}
	//})
}
