package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}

/// hati2 dalam penggunakan closure karna tidak tahu si counter ini di panggil berulang jadi jumlah counternya terus bertambah
