package main

import "fmt"

func insertionSort(a []int) {
	for i := range a {
		j := i
		for j > 0 && a[j-1] > a[j] {
			swap(a, j, j-1)
			j--
		}
	}
}

func swap(a []int, i int, j int) {
	aux := a[i]
	a[i] = a[j]
	a[j] = aux
}

func main() {
	toSort := [4]int{3, 2, 1, 4}
	insertionSort(toSort[:])
	fmt.Println(toSort)
}
