package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		slice[i] = v * 2
	}

	fmt.Println(slice)
}
