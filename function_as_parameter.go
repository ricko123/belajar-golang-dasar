package main

import "fmt"

type Filter func(string) string //type declaration

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Ricko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
