package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	splittedString := strings.Split(sentence, " ")

	sortedString := make([]string, 10)

	exp, _ := regexp.Compile(`\d`)

	for _, s := range splittedString {
		numberIndex := exp.FindStringIndex(s)
		number, _ := strconv.Atoi(s[numberIndex[0]:numberIndex[1]])
		sortedString[number] = s
	}

	return strings.Trim(strings.Join(sortedString, " "), " ")
}

func main() {
	fmt.Println(Order("s9cond th1s"))
}
