package main

import "fmt"

func main() {
	var x [5]float64 // 5 boyutlu float64 tipli x dizisi tanımladık default olarak 0 0 0 0 0
	x[0] = 50
	x[1] = 60
	x[2] = 70
	x[3] = 80
	x[4] = 90

	var toplam float64 = 0 //toplam:0

	for i := 0; i < 5; i++ {
		toplam += x[i] //toplam = 50+60+70+80+90
	}

	fmt.Println(toplam / 5) // 70
}
