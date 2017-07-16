package main

import (
	"fmt"
	"sort"
)

//take a slice and sort it.
//finally at end reverse sort

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	family := people{"yasin", "awni", "dena", "dunya", "mosho", "samir", "diana", "delal"}
	fmt.Println(family)
	sort.Sort(family)
	fmt.Println(family) //this is the result of them sorted
	family.Swap(0, 7)   //i wanted to play with the methods of the interface
	fmt.Println(family)
	a := family.Len()
	fmt.Println(a)
	sort.Sort(sort.Reverse(family)) //reverse the order of the slice
	fmt.Println(family)
}
