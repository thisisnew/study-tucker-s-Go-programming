package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(array)

	slice = append(slice, 100)
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(array)
}
