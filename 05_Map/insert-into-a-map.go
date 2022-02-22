package main

import "fmt"

func GetPrefix(name string) (prefix string) {

	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	/**
	 * Insert key and value into the map.
	 *
	 * <map-variable>[<key>] = <value>
	 */
	prefixMap["Punith"] = "Data Scientist"
	prefixMap["Abhi"] = "Data Analyst"
	prefixMap["Deepak"] = "Sr. Solution Architect"
	prefixMap["Darshan"] = "Cyber Security"

	return prefixMap[name]
}

func main() {

	fmt.Println("What is Punith's role? He is " + GetPrefix("Punith"))

}
