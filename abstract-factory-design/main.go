package main

import (
	"fmt"
)

func main() {

	BMWFactory, err := getMachineFactory("BMW")
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	Bike1 := BMWFactory.makeBike()
	Bike1.printDetails()

	Car1 := BMWFactory.makeCar()
	Car1.printDetails()

	VolkswagenFactory, err := getMachineFactory("Volkswagen")
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	Bike2 := VolkswagenFactory.makeBike()
	Bike2.printDetails()

	Car2 := VolkswagenFactory.makeCar()
	Car2.printDetails()
}
