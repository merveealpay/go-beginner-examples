package main

import "fmt"

//diziler , sabit boyu olan elemanların durdugu bir taşıyıcı, 0 indeksli yani ilk elemanı almak için 0.'yı istiyoruz.

func main() {
	var x [5]int   //burda x bir array 5 boyutlu ve tipi int. default olarak 0 0 0 0 0
	x[4] = 100     // 0 indeksli oldugu icin, aslında 5.elemanı 100 yaptı
	fmt.Println(x) // sonuç: 0 0 0 0 100
}
