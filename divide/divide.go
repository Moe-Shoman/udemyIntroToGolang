package main

import "fmt"

//Create a program that prints to the terminal asking for a user to enter a small number and a larger number. Print the remainder of the larger number divided by the smaller number.

func main() {
	var numOne int
	var numTwo int
	fmt.Print("Please enter in a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter in a smaller number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "/", numTwo, "=", numOne/numTwo)

}
