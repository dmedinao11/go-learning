package main

import "testing"

func TestIsEven(t *testing.T) {
	isTwoEven := isEven(2)

	if !isTwoEven {
		t.Errorf("Number two must be even")
	}
}

func TestTableIsEven(t *testing.T) {
	table := []struct {
		number int
		isEven bool
	}{
		{number: 5, isEven: false},
		{number: 2, isEven: true},
		{number: 5, isEven: true},
		{number: 7, isEven: false},
	}

	for _, testParam := range table {
		if isNumberEven := isEven(testParam.number); isNumberEven != testParam.isEven {
			t.Errorf("Test for number %v fails. isEven? %v | result: %v", testParam.number, testParam.isEven, isNumberEven)
		}
	}
}
