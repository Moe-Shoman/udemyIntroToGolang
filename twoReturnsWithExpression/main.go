package main

import "fmt"

//Modify the  program in twoReturns to use a func expression.

//the prior programs instructions
//Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let’s us know whether or not the argument was even.

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	h, even := half(5)
	fmt.Println(h, even)
}
