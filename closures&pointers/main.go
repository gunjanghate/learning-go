package main

import "fmt"

func counter() func() int {
	var cnt int = 0

	return func() int {
		cnt++
		return cnt
	}
}

func changeNum(num int) {  // taking by value
	num = 20
	fmt.Println("In changeNum: " ,num)
}

func changeNumByRef(num *int) {  // taking by reference
	*num = 30 // dereferencing pointer to get the value
	fmt.Println("In changeNumByRef: " ,*num)
}

func main() {
	/////// closures in go ///////

	// if any inner function refers to the variables defined in the outer function,
	// then it is known as closure in Go

	// in call stack, the outer function's variables will be alive as long as the inner function is alive

	// here, counter is outer function and anonymous function is inner function
	// anonymous function refers to cnt variable of counter function

	// so, anonymous function is closure

	// each time we call counter, a new instance of cnt variable is created
	// so, each counter instance will have its own cnt variable

	count := counter()
	println(count())
	println(count())
	println(count())

	/////// pointers in go ///////

	num:=1
	fmt.Println("Before changeNum:", num) // 1
	changeNum(num) // 20
	fmt.Println("After changeNum:", num)// 1
	// why? because in go, parameters are passed by value
	// so, changeNum function gets a copy of num variable
	// hence, original num variable remains unchanged

	fmt.Println("Before changeNumByRef:", num) // 1
	changeNumByRef(&num) // 30
	fmt.Println("After changeNumByRef:", num) // 30
	// why? because we passed the address of num variable
	// so, changeNumByRef function can modify the original num variable
}