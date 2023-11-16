package main

import (
	"fmt"
)

func main() {

	fmt.Println(">> Binary Search Algorithm")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 10
	index := binary_search(list, value)
	fmt.Printf(">> Binary Search Algorithm -> The index of value %d is: %d\n", value, index)
}

func binary_search(list []int, value int) int {
	low := 0
	high := len(list) - 1
	index := -1
	stop := 0

	for stop == 0 {
		middle := int((low + high) / 2)
		if value == list[middle] {
			fmt.Printf("[%d] Binary Search: value == list[middle]\n", middle)
			stop = 1
			index = middle
			break
		} else if value < list[middle] {
			fmt.Printf("[%d] Binary Search: value < list[middle]\n", middle)
			high = middle - 1
			continue
		} else {
			fmt.Printf("[%d] Binary Search: value > list[middle]\n", middle)
			low = middle + 1
			continue
		}
	}

	return index
}
