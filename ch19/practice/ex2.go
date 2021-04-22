package main

import "fmt"

type Parkinglot struct {
	LotSize int
}

func ParkCar(lot *Parkinglot, carSize int) {
	lot.LotSize -= carSize
}

func (lot *Parkinglot) ParkCar(carSize int) {
	lot.LotSize -= carSize
}

func main() {
	lot := &Parkinglot{LotSize: 100}
	ParkCar(lot, 10)

	lot.ParkCar(10)

	fmt.Println(lot.LotSize)
}
