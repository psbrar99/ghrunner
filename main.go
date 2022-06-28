package main

import "fmt"

func main() {
	fmt.Println(Sum(2, 3, 4, 6, 9))
	fmt.Println(Sum(24, 43, 4, 66, 9))
	fmt.Println(Sum(22222, 355, 400, 6, 9))

}

func Sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s

}
