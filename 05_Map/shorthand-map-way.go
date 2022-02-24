package main

import "fmt"

func GetPrefix(name string) (prefix string) {

	/**
	 * Short hand way to declare and initialize map
	 */
	prefixMap := map[string]string{
		"Punith":  "Data Scientist",
		"Abhi":    "Data Analyst",
		"Darshan": "Cyber Security",
		"Ajay":    "AWS Solution Architect", // comma is needed here
	}

	return prefixMap[name]
}

func main() {

	fmt.Println("What is Punith's role? He is " + GetPrefix("Punith"))

}
