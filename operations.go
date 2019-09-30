package main

import "fmt"

func main() {

	a := 8
	b := 5
	total := a + b

	fmt.Println("total is: ", total)

	total *= 10
	fmt.Println("multiple x10 : ", total)

	total /= 2
	fmt.Println("divide /2 :", total)

	total -= 5
	fmt.Println("subtraction -5 :", total)

	total++
	fmt.Println("total++ :", total)

}
