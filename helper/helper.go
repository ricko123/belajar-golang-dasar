package helper

import "fmt"

var version = "1.0.0"      //tidak bisa diakses dari luar package karena diawali huruf kecil
var Application = "golang" //dapat di akses di luar package karena diawali huruf besar

func sayGoodBye(name string) string { //nama function harus diawali huruf besar
	return "Goodbye" + name
}
func SayHello(name string) string { //nama function harus diawali huruf besar
	return "Hello" + name
}

func Contoh() {
	sayGoodBye("ricko")
	fmt.Println(version)
}
