package main

import "fmt"

//callback function takes in two arguments a slice of type int and a function which takes an argument of type int
//in main we call the function (named "foo") and gave it the arguments of a slice with 10,20,30 and anonymous function which takes n as the int
//foo takes each int in slice and assigns it to n which is then printed to screen
func foo(aSlice []int, hollaBack func(int)) {
	for _, n := range aSlice {
		hollaBack(n)
	}
}

func main() {
	foo([]int{10, 20, 30}, func(n int) {
		fmt.Println(n)
	})

}
