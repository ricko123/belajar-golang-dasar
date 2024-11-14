package main

import (
	"fmt"
	"belajar-golang-dasar/helper"
)

func main (){
	result := helper.SayHello("Ricko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye("ricko"))
}