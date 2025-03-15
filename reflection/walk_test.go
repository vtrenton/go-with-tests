package main

import (
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age      int
	Location string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Trent"},
			[]string{"Trent"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Trent", "Minnesota"},
			[]string{"Trent", "Minnesota"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Trent", 33},
			[]string{"Trent"},
		},
		{
			"nested fields",
			Person{
				"Trent",
				Profile{33, "Minnesota"},
			},
			[]string{"Trent", "Minnesota"},
		},
		{
			"pointers to things",
			&Person{
				"Trent",
				Profile{33, "Minnesota"},
			},
			[]string{"Trent", "Minnesota"},
		},
		{
			"slices",
			[]Profile{
				{33, "Minnesota"},
				{38, "Texas"},
			},
			[]string{"Minnesota", "Texas"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Minnesota"},
				{38, "Texas"},
			},
			[]string{"Minnesota", "Texas"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("Testing maps - dont worry about order", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Canada"}
			aChannel <- Profile{38, "Spain"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Canada", "Spain"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
