package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" //menggunakan underscor(_) agar dapat menjalankan function dari package tanpa menggunakan package tersebut
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
