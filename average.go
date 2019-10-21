package main

import "fmt"

func main() {
	x := [5]float64{98, 22, 36, 40, 52} //float64 e direk cast edebildik diziyi, dizi tanımlamak için rahat bir yol gibi.

	var sum float64 = 0
	for _, eleman := range x {
		sum += eleman
	}
	fmt.Println(sum / float64(len(x))) //aritmetik ortalamayı alıyoruz burda. len fonksiyonu x dizisinin uzunlugunu gostermek icin
	// eger float64 koymasaydık burda, derleyici hata verecekti len bize integer donuyor fakat x float64 tipli,tip donusumu yapmalıyız

}
