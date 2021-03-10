package main

import (
	"fmt"
	"strings"
	"testing"
)

type TestCases struct {
	expected string
	strings  []string
	flags    []int
	testName string
}

func Test(t *testing.T) {
	testCases := []TestCases{
		{
			`I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.`,
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]int{0, 0, 0, 0, 0, 0},
			"Test flag C",
		},
		{
			`3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.`,
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]int{1, 0, 0, 0, 0, 0},
			"Test flag C",
		},
		{
			`I love music.
I love music of Kartik.
I love music of Kartik.`,
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]int{0, 1, 0, 0, 0, 0},
			"Test flag D",
		},
		{
			`
Thanks.`,
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]int{0, 0, 1, 0, 0, 0},
			"Test flag D",
		},
		{
			`I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.`,
			[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."},
			[]int{0, 0, 0, 1, 0, 0},
			"Test flag I",
		},
		{
			`We love music.

I love music of Kartik.
Thanks.`,
			[]string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks."},
			[]int{0, 0, 0, 0, 1, 0},
			"Test flag F",
		},
		{
			`I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`,
			[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks."},
			[]int{0, 0, 0, 0, 0, 1},
			"Test flag F",
		},
		{
			`1 I love music.
1 
1 I love music of Kartik.
1 We love music of Kartik.
1 Thanks.`,
			[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks."},
			[]int{1, 0, 0, 0, 0, 1},
			"Test flag F",
		},
		{
			`1 I love music.
2 I love music of Kartik.`,
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.",
				"I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]int{1, 1, 0, 0, 0, 0},
			"Test flag D",
		},
	}

	for _, test := range testCases {
		result := correctUniqWork(&test.flags, &test.strings)
		res := strings.Join(result, "\n")
		if res != test.expected {
			t.Errorf(fmt.Sprintf("%v\nПолучили:\n%v\n\nОжидалось:\n%v", test.testName, res, test.expected))
		}
	}
}
