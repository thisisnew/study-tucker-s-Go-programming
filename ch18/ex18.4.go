package main

import "fmt"

func changeArray(array [5]int) {
	array[2] = 200
}

func changeSlice(slice []int) {
	slice[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)
}
