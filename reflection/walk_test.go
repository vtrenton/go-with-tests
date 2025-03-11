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
	Age  int
	City string
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
}
