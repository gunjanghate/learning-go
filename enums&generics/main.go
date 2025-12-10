package main

import "fmt"


// enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota // predefined constant starting from 0
	Confirmed // 1
	Prepared // 2
	Delivered // 3
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status to", status)
}


func main(){

	changeOrderStatus(Confirmed)
}

