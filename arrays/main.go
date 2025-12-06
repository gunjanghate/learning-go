package main

import "fmt"

func main() {

	/////// arrays in go ///////

	// - fixed size
	// - same data type
	// - Constant time access
	// - memory optimized
	// - Memory stored in contiguous locations
	
	
	var nums [4]int
	// will get automatically initialized to zero
	// false for bool, "" for string
	fmt.Println(len(nums))

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	fmt.Println(nums)

	nums1:= [3]int{10,20,30}

	fmt.Println(nums1)

	// multidimensional array
	// var matrix [2][2]int
	// matrix[0][0] = 1
	// matrix[0][1] = 2
	// matrix[1][0] = 4
	// matrix[1][1] = 5

	// fmt.Println(matrix)

	m:= [2][2]int{
		{1,2},
		{3,4},
	}
	fmt.Println(m)



}