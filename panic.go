package main

import "fmt"

func endApp() {
	fmt.Println("End go")
	message := recover()
	/*
		dengan menggunakan recover dapat mendapatkan pesan error dari panic sehingga program tidak berhenti total
		dan bisa berjalan ke code berikutnya
	*/
	fmt.Println("Terjadi erro", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Ricko sapto Prihartanto")
}
