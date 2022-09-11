package main

import "fmt"

func reverseArray(a []int32) []int32 {
	// Write your code here
	reversedArray := []int32{}

	for i := len(a) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, a[i])
	}

	return reversedArray
}

func main() {

	toReverse := []int32{1, 2, 3}
	fmt.Println(reverseArray(toReverse))

}
