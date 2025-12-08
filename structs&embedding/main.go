package main


import "fmt"
import "time"

type customer struct {
	name string
	email string
}

type order struct {
	id string
	amt float32
	status string
    createdAt time.Time // nanosec precision
	customer customer
}

// method to change order status
func (o *order) changeStatus(status string){
	o.status = status // struct automatically dereferenced
}

func (o *order) changeAmt(amt float32){
	o.amt = amt
}

func newOrder(id string, amt float32, status string) *order{
	o := order{
		id: id,
		amt: amt,
		status: status,
		createdAt: time.Now(),
	}
	return &o

}

// use * when you want to modify the original struct
// without * you will be modifying a copy of the struct

func main(){
	// var order o =order{}
	// o := order{
	// 	id: "ord001",
	// 	amt: 250.75,
	// 	status: "pending",
	// 	createdAt: time.Now(),
	// }
	o := newOrder("ord001", 250.75, "pending")

	o2 := order{
		id: "ord002",
		amt: 500.00,
		status: "completed",
		createdAt: time.Now(),
	}
	o.amt += 100.25

	fmt.Println("Order:", o)
	// fmt.Println("Order ID:", o.id)
	// fmt.Println("Order Amount:", o.amt)
	// fmt.Println("Order Status:", o.status)
	// fmt.Println("Order Created At:", o.createdAt)

	// comparing two structs
	// if o == o2 {
	// 	fmt.Println("Orders are equal")
	// } else {
	// 	fmt.Println("Orders are not equal")
	// }

	o.changeStatus("shipped")
	o2.changeAmt(203.4)

	fmt.Println("Order after status change:", o.status)
	fmt.Println("Order2 after amount change:", o2.amt)
	

	lang := struct{ // anonymous struct (inline struct )
		name string
		isGood bool
	}{
		name: "GoLang",
		isGood: true,
	}

	fmt.Println("Anonymous Struct:", lang)


	// embedding structs
	// cust := customer{
	// 	name: "Alice Smith",
	// 	email: "gmail.com",
	// }
	newOrder := order{
		id: "ord003",
		amt: 750.50,
		status: "processing",
		// customer: cust,
		customer: customer{
			name: "John Doe",
			email: "john.doe@example.com",
		},
	}

	fmt.Println("New Order with Customer:", newOrder)
	fmt.Println("Customer details: ", newOrder.customer.email)
}