package main

import "time"

type Courier struct {
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(p *Product) *Parcel {
	parcel := &Parcel{}

	parcel.Pdt = p
	parcel.ShippedTime = time.Now()

	return parcel
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()

	return p.Pdt
}
