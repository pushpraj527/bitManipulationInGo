/*
Given a number, check whether it is an even or odd
*/
package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var x int
	fmt.Scanln(&x)
	if isEven(x) == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}

func isEven(n int) int {
	return n & 1
}
