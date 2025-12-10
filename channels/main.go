package main

import (
	"fmt"
	"time"
	// "math/rand"
)

// channels are used to communicate between goroutines
// they are like pipes through which data can be sent and received

// unbuffered channels - data is sent and received one at a time
// buffered channels - data is sent and received in batches 

// func processNum(numChan chan int){
// 	for num:= range numChan{
// 		fmt.Println("Processing number...", num)
// 		time.Sleep(time.Millisecond*500)
// 	}
// 	// fmt.Println("Processing number...", <-numChan)

// }

// func sum(res chan int, num1,num2 int){
// 	res <- num1 + num2
// }

func task(done chan bool){
	defer func(){
		done <- true
	}()

	fmt.Println("Task is being performed")
}



func emailSender(emailChan <-chan string, done chan<- bool){ // emailChan is a receive-only channel // done is a send-only channel
	defer func(){
		done <- true
	}()
	for email := range emailChan{
		fmt.Println("Sending email to:", email)
		time.Sleep(time.Second)
	}

	
	
}
func main() {

/////// Unbuffered Channels ///////

	// Creating Channel
	// messageChannel := make(chan string)

	// // Sending data in channel
	// messageChannel <- "ping" // blocking operation till there is a goroutine to receive the data

	// // recieve the data

	// msg := <-messageChannel

	// fmt.Println("Messgae:", msg)
	// Error
	// 	fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan send]:
	// main.main()
	//         D:/my/GUNJAN/Gunjan Go/learning/channels/main.go:10 +0x36
	// exit status 2

	// because there is no goroutine to receive the data from channel 


	// numChan := make(chan int)

	// go processNum(numChan)

	// for{
	// 	numChan <- rand.Intn(100)
	// }


	// time.Sleep(time.Second*2)


	// res := make(chan int)
	// go sum(res, 5, 7)

	// ress:= <-res  // blocking operation till there is data to receive
	// fmt.Println("Sum is:", ress)


	// channels for synchronization

	done := make(chan bool)

	// go task(done)

	// <- done  // blocking operation till there is data to receive
	// fmt.Println("Task completed")


/////// Buffered Channels ///////
// we can send multiple data without waiting for receiving it

emailChan:=make(chan string, 10)

// emailChan <- "1@gmail.com"
// emailChan <- "2@gmail.com"

// fmt.Println(<-emailChan)
// fmt.Println(<-emailChan)

go emailSender(emailChan, done)

for i:=0;i<5;i++{
	emailChan <- fmt.Sprintf("%d@gmail.com", i)
}

fmt.Println("done sending..")

// close Channel (important to close channel to avoid deadlock in receiver goroutine)
close(emailChan)
<- done



chan1 := make(chan int)
chan2 := make(chan string)

go func(){
	chan1 <- 10
}()
go func(){
	chan2 <- "hello"
}()

for i:= 0; i<2; i++{
	select{
	case chan1Val := <- chan1:
		fmt.Println("Received from chan1:", chan1Val)
	case chan2Val := <- chan2:
		fmt.Println("Received from chan2:", chan2Val)	
	}
}

}
