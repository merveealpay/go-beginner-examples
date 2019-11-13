package main

import "fmt"

func main() {
	//slice Array'in bir kesiti,parçası, array gibi indeksi ve boyu var ama bir farkla, boyu degisebilir
	var x []float64           //x in boyu sıfır , bir array gibi
	x := make([]float64, 5)   // make ile boyut veriyoruz, tipi float64 olan 5 elemanlı dizi.
	x := make(float64, 5, 10) //kapasitesi 10 olan array'den 5lik dilim

	array := [5]float64{1, 2, 3, 4, 5}
	x := array[0:3] // 1,2,3. Arrayin 0.elemanından itibaren 0.dahil , index 3e gelene kadar yazdır demek.
	array[0:5]      // 1,2,3,4,5
	array[1:4]      //2,3,4

	array[:]  // 1,2,3,4,5
	array[0:] //index 0'dan başla, arrayin sonunaa kadar
	array[2:] //index 2denn başla arrayin sonuna kadar
	array[:2] //index 0dan 2ye kadar 1,2

	array[0:len(array)] //1,2,3,4,5
	array[:len(array)]  //1,2,3,4,5

	//append, slice ın sonuna eleman eklemek icin kullanırız.
	slice1 := []int{1, 2, 3}       //1,2,3
	slice2 := append(slice1, 4, 5) //1,2,3,4,5
	fmt.Println(slice2)

	//copy iki argüman alır dst ve src. src dakiler dst a kopyalanıyor direk üstüne yazılır hatta, varolan silinir
	// eger src dst dan büyükse dst nin kapasitesi kadar kopyalanır.
	slice1 := []int{1, 2, 3} //1,2,3
	slice2 := make([]int, 2) //2 boyutlu dizi olusturduk
	copy(slice2, slice1)     //slice1'i slice2'ye kopyala diyoruz
	fmt.Println(slice2)      //1,2 yazdıracak cunku slice2 2 boyutlu

	//genel olarak dizi mantıgı hep n.indeksteki elemanı ver seklindedir fakat ya indeks sayısal olmasa? mesela soyad gibi string bir ifade index olsa ve
	//soyad a gore aramak istesek, burda map karsımıza cıkıyor.

}
