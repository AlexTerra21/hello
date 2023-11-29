package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 100)

	for i, _ := range slice {
		slice[i] = i + 1
	}

	fmt.Println(slice)

	// slice1 := slice[10:90]
	slice = append(slice[0:10], slice[90:100]...)

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	fmt.Println(slice)

}
