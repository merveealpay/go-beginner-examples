package main

import "fmt"

func main() {

	fmt.Println("please type a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 5
	fmt.Println("the output is: ", output)
}
