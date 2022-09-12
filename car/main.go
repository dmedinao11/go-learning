package main

import (
	"fmt"

	"github.com/dmedinao11/go-learning/car"
)

func main() {
	myCar := car.Car{Year: 2022, Brand: "chevrolet", Name: "sail", CurrentFuel: 100, MaxFuelCapacity: 1000}

	myCar.Refuel()
	fmt.Println(myCar)

	myCar.ExpendFuel(1200)
	fmt.Println(myCar)
}
