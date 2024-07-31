package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//ini bukan lagi function tapi method

func (customer Customer) sayhello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var cust Customer
	fmt.Println(cust)

	cust.Name = "Ricko"
	cust.Address = "Jakarta"
	cust.Age = 29

	fmt.Println(cust)
	fmt.Println(cust.Name)
	fmt.Println(cust.Address)
	fmt.Println(cust.Age)

	custom := Customer{
		Name:    "Ricko",
		Address: "Matraman",
		Age:     29,
	}

	fmt.Println(custom)

	budi := Customer{"Budi Santoso", "Manggarai", 30}

	fmt.Println(budi)

	budi.sayhello("agus")
	custom.sayhello("agus")
	cust.sayhello("agus")
}
