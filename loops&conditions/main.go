package main

import (
	"fmt"
	"time"
)

// for -> only construct in go for looping
func main() {
	/////// loops in go ///////

	// while
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("Loop finished")

	// infinite loop
	// for{
	// 	fmt.Println("infinite loop")
	// }

	//Classic for loop
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	// 1.22 Range
	for k := range 3 {
		fmt.Println(k)
	} // print 0, 1 ,2

	// conditionals in go ///////

	// If Else
	num := 10
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// we can declare variable in if condition
	if age := 15; age >= 18 {
		fmt.Println("You can vote", age)
	} else if age >= 12 {
		fmt.Println("You cannot vote", age)
	}
	// If Else If
	marks := 85
	if marks >= 85 {
		fmt.Println("Grade A")
	} else if marks >= 70 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// Switch Case

	// Simple Switch
	a := 5

	switch a {
	case 1:
		fmt.Println("One") // no need to write break, its by default
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number") // default also optional
	}

	// Multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's a Weekday")
	}

	// type Switch
	whoAmI := func(i interface{}) {
		switch b := i.(type) {
		case int:
			fmt.Printf("I'm an Integer and my value is %d\n", b)
		case string:
			fmt.Printf("I'm a String and my value is %s\n", b)
		}
	}

	whoAmI(42)
	whoAmI("Hello Go")

}
