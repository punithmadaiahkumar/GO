package main

import "./greet"

func main() {

	var github = greet.Profile{"Punith Gowda", "Punith", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	greet.Greet(github, greet.CreatePrintFunction("!"), true)
}
