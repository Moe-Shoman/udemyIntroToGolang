package main

import "fmt"

//practicing recursion. take a number and subtract 1 from it and multiply it against the original number and then keep lookiping until the number is equal to zero

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(10))

}
