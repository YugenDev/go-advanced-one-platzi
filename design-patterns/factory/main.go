package main

import (
	"fmt"
)

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}
	if computerType == "desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("Invalid computer type")
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Name: %s, Stock: %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
