package main

import "fmt"

type Brand string

// brand adında yeni bir tip

const (
	FACEBOOK Brand = "Facebook" //tipi brand olan bir FACEBOOK , sabit bu aynı enumlar gibi oluşturuluyor.
	GOOGLE   Brand = "Google"
)

func PrintBrand(brand Brand) { //buraya brandi nesne olarak yolladım ve bunu ekrana basan fonksiyon yazdım
	fmt.Println(brand)
}

func main() {
	PrintBrand(GOOGLE) //olusturdugum fonksiyonu GOOGLE parametresiyle cagırdım
}
