package main

import "fmt"

type machineCar interface {
	setColor(color string)
	setTopSpeed(speed int)
	getColor() string
	getTopSpeed() int
	getCompany() string
	printDetails()
}

type Car struct {
	company  string
	color    string
	topSpeed int
}

func (c *Car) setColor(color string) {
	c.color = color
}

func (c *Car) setTopSpeed(speed int) {
	c.topSpeed = speed
}

func (c *Car) getColor() string {
	return c.color
}

func (c *Car) getTopSpeed() int {
	return c.topSpeed
}

func (c *Car) getCompany() string {
	return c.company
}

func (c *Car) printDetails() {
	fmt.Println("[CAR] Company:", c.company, "Color:", c.color, "TopSpeed:", c.topSpeed)
}
