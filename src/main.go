package main

import (
	"fmt"

	"github.com/miguelfuentesgo/practices/modules/to_do_list_file"
)

func main() {
	fmt.Println("BIENVENIDO A ESTA PRACTICA ")

	// //Using module searcher
	// filteredItems := searcher.SearchValue("guaca")
	// fmt.Println(filteredItems)

	// //Using module goroutines
	// goroutines.RunGoroutines()

	// //Using module mutex
	// mutex.RunMutextExample()

	// //Using module calculator
	// err := calculator.Run()

	// if err != nil {
	// 	fmt.Println("ERROR", err)
	// }

	// //Using module random_number
	// random_number.Run()

	// Using module password_generator
	// password_generator.Run()

	//Using module to_do_list_file
	to_do_list_file.Run()
}
