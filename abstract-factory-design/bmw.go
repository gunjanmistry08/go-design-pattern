package main

type BMW struct {
}

func (f *BMW) makeCar() machineCar {
	return &BMWCar{
		Car: Car{
			company:  "BMW",
			color:    "red",
			topSpeed: 200,
		},
	}
}

func (f *BMW) makeBike() machineBike {
	return &BMWBike{
		Bike: Bike{
			company:  "BMW",
			color:    "red",
			topSpeed: 400,
		},
	}
}
