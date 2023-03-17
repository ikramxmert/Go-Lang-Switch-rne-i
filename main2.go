package main

import "fmt"

func main() {

	/*Burada kullanılan giriş için ikram mert in sadece ikram kısmını alır
	reader := bufio.NewReader(os.Stdin) //bufio giriş çıkış işlemleri içindir.. yeni okuyucu os pakedi işletim sistemi yani içinde std in input u okur
	fmt.Print("Enter text : ") //println yazsaydık alt satıra inerdi
	var str string
	fmt.Scanf("%s", &str)
	*/

	/*Bu girdi ise ikram mert in tamamını alır
	reader := bufio.NewReader(os.Stdin) //bufio giriş çıkış işlemleri içindir.. yeni okuyucu os pakedi işletim sistemi yani içinde std in input u okur
	fmt.Print("Enter text : ")          //println yazsaydık alt satıra inerdi
	str, _ := reader.ReadString('\n')   //eror da döndüğü için black identifier kullandık
	fmt.Println(str)
	*/

	//Switch ile Alakalı güzel bir örnek
	var derece string
	print("Lütfen Notunuzu Giriniz : ")
	var n int
	fmt.Scanf("%v", &n)

	var ort float64
	ort = float64(n)
	switch {
	case ort < 50:
		derece = "FF"
	case ort < 60:
		derece = "DD"
	case ort < 70:
		derece = "CC"
	case ort < 85:
		derece = "BB"
	case ort <= 100:
		derece = "AA"
	default:
		derece = "HATALI SONUÇ"
	}
	fmt.Printf("Harf Notunuz : %s", derece)

}
