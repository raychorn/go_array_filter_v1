package main

import (
    "testing"
)

func Test1(t *testing.T) {
    type test struct {
        input      []int
        expected   [][]int
    }

    tests := []test{
        {input: []int{1, 7, 3, 9}, expected: [][]int{{1, 7, 3, 9}, {1, 9, 3, 7}, {7, 9, 1, 3}, {3, 7, 1, 9}, {3, 9, 1, 7}} }, 
    } 

    for _, tc := range tests {
        // perform setUp before each test here
		actual := performTest(tc.input)
		for i, tcExpected := range tc.expected {
			//fmt.Println(fmt.Sprintf("(%d) expected '%s', got '%s'", i, prettyPrintTheGuess(tcExpected), prettyPrintTheGuess(actual)))
			if (!compareArrays(tcExpected, actual[i])) {
				t.Errorf("expected '%s', got '%s'", prettyPrintTheGuess(tcExpected), prettyPrintTheGuess(actual[i]))
			}
		}
    }
}