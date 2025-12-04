package main

import "fmt"


// const page = 1 // allowed
func main(){
	/////// Variables in Go /////

	// var name = "Gopher" // type inference
	// name := "Gopher"    // shorthand syntax
	
	// var name string
	// name = "Gopher"
	
    // var price float32 
	// price = 19.99
	
	var name string = "Gopher"
	var age int = 5
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	/////// Constants in Go /////

	// const constName = "GoLang"
	// const constAge int64 = 10

	const (
		constName = "GoLang"
		constAge  = 10
	)

	fmt.Println(constName, ":", constAge )
}