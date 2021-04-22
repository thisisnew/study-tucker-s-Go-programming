package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name        string
	Age         int
	Point       int
	SuccessRate float64
}

type Players []Player

func (p Players) Len() int {
	return len(p)
}

func (p Players) Less(i int, j int) bool {
	return p[i].Point < p[j].Point
}

func (p Players) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	p := []Player{
		{Name: "나통키", Age: 13, Point: 45, SuccessRate: 78.4},
		{Name: "오맹태", Age: 16, Point: 24, SuccessRate: 67.4},
		{Name: "오동도", Age: 18, Point: 54, SuccessRate: 50.8},
		{Name: "황금산", Age: 16, Point: 36, SuccessRate: 89.7},
	}

	sort.Sort(sort.Reverse(Players(p)))
	fmt.Println(p)
}
