package main

import "fmt"

func main() {
	num1 := 0
	num2 := 1
	sum := 0
	i := 0
	for i < 4 {
		sum = num1 + num2
		if sum/2 == 0 {
			fmt.Println(sum)
		}
		fmt.Println(sum, " -- this is an odd number")
	}

}
