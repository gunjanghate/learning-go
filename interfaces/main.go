package main

import "fmt"


type paymentr interface{
	pay(amt float32)
	refund(amt float32, acc string)
}




type payment struct{
	// gateway stripe
	gateway paymentr
}


func (p payment) pay(amt float32) {
	// razorpaygw := razorpay{}
	// razorpaygw.pay(amt)
	// stripegw := stripe{}
	p.gateway.pay(amt)
}

type razorpayx struct{}

func (r razorpayx) pay(amt float32) {
	// logic to make payment using razorpay
	fmt.Println("Payment of", amt, "made using Razorpay")
}

func (r razorpayx) refund(amt float32, acc string) {
	// logic to refund payment using razorpay
	fmt.Println("Refund of", amt, "made to account", acc, "using Razorpay")
}
// type stripe struct{}

// func (s stripe) pay(amt float32) {
// 	// logic to make payment using stripe
// 	fmt.Println("Payment of", amt, "made using Stripe")
// }



// type fakepay struct{}

// func (f fakepay) pay(amt float32) {
// 	// logic to make payment using fakepay
// 	fmt.Println("Payment of", amt, "made using Fakepay")
// }

func main() {
    // payment := payment{}
	// payment.makePayment(1000)

	razorpayx := razorpayx{}
	


	// fakepay := fakepay{}
	// payment := payment{gateway: fakepay} // here we only can pass stripe as gateway
	                                     // because of static typing
										 // thats why we use interfaces
	// payment.makePayment(250)


	/////// Interfaces in Go ///////
	// An interface is a type that defines a set of method signatures
	// A type implements an interface by implementing its methods
	// Interfaces are used to achieve polymorphism in Go
	// A variable of interface type can hold any value that implements the interface

	newpay := payment{
		gateway: razorpayx,
	}
	newpay.pay(500)
}