package main

import (
	"fmt"
)

func main() {
	fmt.Println(">> Linear Search Algorithm")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 10
	index := linear_search(list, value)
	fmt.Printf(">> Linear Search Algorithm -> The index of value %d is: %d\n", value, index)
}

func linear_search(list []int, value int) int {
	index := -1
	for i, v := range list {
		fmt.Printf("[%d] Linear Search: v == value\n", i)
		if v == value {
			index = i
			break
		}
	}
	return index
}
