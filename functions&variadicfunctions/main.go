package main

import "fmt"

func add(a, b int) int {
	return a+b;
}

func getLangs() (string, string, string) {
	return "golang", "python", "java"
}

// func processIt(pn func(a int) int){
// 	pn(1)
// }
func processIt() func(a int) int {
	return func(a int) int {
		return a
	}
}




func sum(nums ...int) int{
	total :=0

	for _, num := range nums{
		total += num
	}
	return total
}

func main() {
	//////// funtions in go ////////
	res := add(2, 3)
	fmt.Println("Result:", res)
	// lang1, lang2, lang3 := getLangs()
	lang1, lang2, _ := getLangs()
	fmt.Println("Languages:", lang1, lang2)

	// fn := func(a int) int{
	// 	return 2
	// }

	// processIt(fn) // passing function as parameter
	
	fn := processIt() // returning function from function
	fmt.Println("Fn result:", fn(10))


	/////// variadic functions ///////

	// variadic function can take multiple arguments
	// all arguments are of same type
	
	nums := []int{1,2,3,4,5}
	total := sum(nums...)
	fmt.Println("Sum:", total)



}
