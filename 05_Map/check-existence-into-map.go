package main

import "fmt"

func GetPrefix(name string, mustDel bool) (prefix string) {

	/**
	 * Short hand way to declare and initialize map
	 */
	prefixMap := map[string]string{
		"Punith":  "Data Scientist",
		"Abhi":    "Data Analyst",
		"Darshan": "Networking",
		"Dhanush": "AWS Production services", // comma is needed here
	}

	if mustDel {
		delete(prefixMap, "Punith")
	}

	/**
	 * Check if the value exist into the map or not.
	 */
	if _, exists := prefixMap[name]; exists {
		return prefixMap[name]
	} else {
		return "dude"
	}

}

func main() {

	fmt.Println("What is Punith's role? He is " + GetPrefix("Punith", false))

	fmt.Println("What is Punith's role? He is " + GetPrefix("Punith", true))

}
