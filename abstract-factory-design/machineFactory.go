package main

import "fmt"

type machineFactory interface {
	makeCar() machineCar
	makeBike() machineBike
}

func getMachineFactory(brand string) (machineFactory, error) {
	if brand == "Volkswagen" {
		return &Volkswagen{}, nil
	}
	if brand == "BMW" {
		return &BMW{}, nil
	}
	return nil, fmt.Errorf("brand not supported yet")
}
