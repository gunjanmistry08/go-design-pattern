package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "SNOW"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "SNOW"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 2
}

func (b *iglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
