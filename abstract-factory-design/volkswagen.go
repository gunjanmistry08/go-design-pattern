package main

type Volkswagen struct {
}

func (f *Volkswagen) makeCar() machineCar {
	return &VolkswagenCar{
		Car: Car{
			company:  "Volkswagen",
			color:    "red",
			topSpeed: 200,
		},
	}
}

func (f *Volkswagen) makeBike() machineBike {
	return &VolkswagenBike{
		Bike: Bike{
			company:  "Volkswagen",
			color:    "red",
			topSpeed: 400,
		},
	}
}
