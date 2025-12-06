package main

import (
	"fmt"
	"slices"
)


func main(){
	/////// slices in go ///////

	// - dynamic size
	// - same data type
	// - useful methods built on arrays
	// - more convenient to use than arrays
    

	// uninitialised slice is nil
	// var nums []int
	// fmt.Println(nums, len(nums), nums==nil)

	            //make(type, initial size, capacity)
	var nums = make([]int, 0, 5)
	fmt.Println("Array" ,nums, "Capacity:", cap(nums), "Length:", len(nums), "IsNil:", nums==nil)

	nums= append(nums, 1)
	nums= append(nums, 2)
	nums= append(nums, 3)
	nums= append(nums, 4)
	// nums= append(nums, 5) 
	// so basically when we exceed capacity, go creates a new array with double the capacity and copies the elements

	// var nums = make([]int, 1, 5)
	// nums[0] = 1
	// nums[1] = 2 // will give runtime error: index out of range
	// nums[10] = 10 // will give runtime error: index out of range
	
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// copy function
	newNums := make([]int, len(nums), cap(nums)*2)
	copy(newNums, nums)
	fmt.Println("New Slice", newNums, "Capacity:", cap(newNums), "Length:", len(newNums), "IsNil:", newNums==nil)

	// slice operator
	var a = []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Original Slice", a)
	// b := a[2:5] // from index 2 to index 4
	// b := a[:5] // from index 0 to index 4
	// c := a[4:] // from index 4 to end
	// fmt.Println("Sliced Slice b:", b)
	// fmt.Println("Sliced Slice c:", c)


	//slice
	var n1 = []int{1,2}
	var n2 = []int{3,4}

	fmt.Println(slices.Equal(n1,n2))
	fmt.Println(slices.BinarySearch(n1, 3))
}