package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCountRepeat(t *testing.T) {
	input := "a"
	count := 5

	repeated := Repeat(input, count)
	outcount := strings.Count(repeated, input)

	if outcount != count {
		t.Errorf("expected a length of %q but got %q", count, outcount)
	}
}

func TestRepeatRepeat(t *testing.T) {
	strlib := strings.Repeat("a", 5)
	custom := Repeat("a", 5)

	if custom != strlib {
		t.Errorf("the stdlib Repeat output is %q, the custom Repeat func output is %q", strlib, custom)
	}
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
