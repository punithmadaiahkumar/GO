package main

import "fmt"

type Profile struct {
	name string
	role string
}

func main() {

	slice := []Profile{
		{"Punith", "Data Scientist"},
		{"Kumar", "Data Engineerr"},
		{"Abhi", "Data Analyst"},
		{"Ajay", "Python Developer"}, // comma is needed here
	}

	/**
	 * _ is index; we are not using it; thus replacing with _
	 */
	for _, employee := range slice {
		fmt.Println("Profile: " + employee.name + " (" + employee.role + ")")
	}

}
