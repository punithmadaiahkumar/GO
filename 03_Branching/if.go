package main

import "./greeting";

func main() {

	var github = greeting.Profile{"Punith Gowda", "PunithGowda", "Welcome to Go world!"}

	/**
	 * Call the function and pass the data to the function
	 */
	greeting.Greeting(github, greeting.CreatePrintFunction("!"), false)