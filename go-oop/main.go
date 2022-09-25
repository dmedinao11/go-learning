package main

import (
	"fmt"

	"github.com/dmedinao11/go-learning/go-oop/employee"
)

func main() {
	myEmployee := employee.Employee{}
	myEmployee.SetId(1)
	myEmployee.SetName("Daniel")
	fmt.Printf("{name: %s, id: %d}\n", myEmployee.GetName(), myEmployee.GetId())

	myPerson := employee.Person{}
	myPerson.SetName("Miguel")
	fmt.Printf("{name: %s}\n", myPerson.GetName())

}
