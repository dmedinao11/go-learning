package car

import "fmt"

type Car struct {
	Year            int
	Brand           string
	Name            string
	CurrentFuel     int
	MaxFuelCapacity int
}

func (myCar *Car) Refuel() {
	myCar.CurrentFuel = myCar.MaxFuelCapacity
}

func (myCar *Car) MultiplyCapacity() {
	myCar.MaxFuelCapacity *= 2
}

func (myCar *Car) ExpendFuel(fuelToUse int) {
	if myCar.canExpendFuel(fuelToUse) {
		myCar.CurrentFuel -= fuelToUse
		return
	}

	fmt.Println("Cannot expend fuel")
}

func (myCar Car) canExpendFuel(fuelToUse int) bool {
	myCar.Brand = "I don't do anything"
	return myCar.CurrentFuel-fuelToUse > 0
}
