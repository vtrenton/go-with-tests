package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("standard repeat validation using fixed string", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("validate counted output matches input", func(t *testing.T) {
		input := "a"
		count := 5

		repeated := Repeat(input, count)
		outcount := strings.Count(repeated, input)

		if outcount != count {
			t.Errorf("expected a length of %q but got %q", count, outcount)
		}
	})

	t.Run("validate output matches string.Repeat", func(t *testing.T) {
		strlib := strings.Repeat("a", 5)
		custom := Repeat("a", 5)

		if custom != strlib {
			t.Errorf("the stdlib Repeat output is %q, the custom Repeat func output is %q", strlib, custom)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("a", 5)
	fmt.Println(out)
	// Output: aaaaa
}
