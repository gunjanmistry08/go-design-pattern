package main

import "fmt"

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) Builder {
	switch builderType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	default:
		fmt.Println("ERROR in builder Type")
		return nil
	}

}
