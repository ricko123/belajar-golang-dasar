package main

import "fmt"

func getGoodBy(name string) string {
	return "Good Bye " + name
}

func main() {
	goodby := getGoodBy

	fmt.Println(goodby("Ricko"))
}
