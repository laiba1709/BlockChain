package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {

	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"John", 82000, "Web Developer"}
	emp3 := employee{"Doe", 80000, "Back-end Developer"}

	emplys := []employee{emp1, emp2, emp3}
	
	comp := company{companyName: "Tetra", employees: emplys}

	Valprint(&comp)
	

}

func Valprint(comp *company) {
	fmt.Println("Company Name: ",comp.companyName)
}