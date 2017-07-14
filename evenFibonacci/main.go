package main

import (
	"fmt"
)

//Find a problem at projecteuler.net then create the solution. Add a comment beneath your solution that includes a description of the problem. Upload your solution to github. Tweet me a link to your solution.
//problem 2
//i copied a solution from http://project-euler-answers-in-go.blogspot.com/2011/07/problem-2.html
//i am going to comment on what i think its doing and if im wrong please tell me
func main() {
	a, b, sum := 1, 2, 0 //on this line we create 3 variables
	for b < 4000000 {    //for loop that will run until b is greater than 4000000 but should it be <= ?
		if b%2 == 0 { //there is an if that says add the sum to b and update sum value if b is an even number
			sum += b
		}
		a, b = b, a+b //this line assigns b to a and a+b to b
	}
	fmt.Println(sum) //this line just prints sum
}
