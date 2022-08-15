package main

import "fmt"

type machineBike interface {
	setColor(color string)
	setTopSpeed(speed int)
	getColor() string
	getTopSpeed() int
	getCompany() string
	printDetails()
}

type Bike struct {
	company  string
	color    string
	topSpeed int
}

func (b *Bike) setColor(color_ string) {
	b.color = color_
}

func (b *Bike) setTopSpeed(topSpeed_ int) {
	b.topSpeed = topSpeed_
}

func (b *Bike) getColor() string {
	return b.color
}

func (b *Bike) getTopSpeed() int {
	return b.topSpeed
}

func (b *Bike) getCompany() string {
	return b.company
}

func (b *Bike) printDetails() {
	fmt.Println("[BIKE] Company:", b.company, "Color:", b.color, "TopSpeed:", b.topSpeed)
}
