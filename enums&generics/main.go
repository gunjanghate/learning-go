package main

import "fmt"


// enumerated types

type OrderStatus string

const (
	// Received OrderStatus = iota // predefined constant starting from 0
	// Confirmed // 1
	// Prepared // 2
	// Delivered // 3
	Rece OrderStatus = "Received"
	Conf OrderStatus = "Confirmed"
	Prep OrderStatus = "Prepared"
	Deliv OrderStatus = "Delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status to", status)
}


func printSlice[T comparable, V string](items []T, name V){ // comparable constraint ensures that the type T supports comparison operators eg., strings, integers, booleans, etc.
	for _, item:= range items{
		fmt.Println(item, name)
	}
}


type Stack[T any] struct {
	elements []T
}

func main(){

	changeOrderStatus(Conf)


	/////// Generics in go 1.18+
	ints := []int{1,2,3,4,5}
	// strings :=[]string{"a","b","c","d","e"}

	// printSlice(strings)
	// float32s := []float32{1.1,2.2,3.3,4.4,5.5}
	// printSlice(float32s)
	printSlice(ints, "GG")

	st := Stack[int]{
		elements: []int{10,20,30,40,50},
	}

	for _, num := range st.elements{
		fmt.Println(num)
	}


}


