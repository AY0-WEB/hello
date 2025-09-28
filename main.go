package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Print("enter a number")
	fmt.Scan(&x)
	fmt.Print("enter the second number")
	fmt.Scan(&y)
	fmt.Println("the sum of the two numbers is", x+y)

}
