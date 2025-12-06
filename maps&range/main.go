package main
import "fmt"
import "maps"

func main(){
	/////// maps & range in go ///////

	// maps -> key value pairs
	// unordered collection
	// dynamic size
	// keys must be unique
	// keys and values can be of any data type

	m:= make(map[string]string)

	m["name"] = "GG"
	m["course"] = "GoLang"
	m["website"] = "learncodeonline.in"

	// if key does not exist, returns zero value of value type
	fmt.Println("Name:", m["name"])
	fmt.Println("Age:", m["age"]) // returns ""

	delete(m, "website")
	fmt.Println("Map:", m)
	fmt.Println("Length:", len(m))
	clear(m) // maps get empty

	m1 := map[string]int{
		"maths": 90,
		"science": 85,
		"english": 88,
	}
	// fmt.Println("Map1:", m1)

	k, ok := m1["chemistry"]
	fmt.Println("Maths marks:", k, ok)

	m2 := map[string]int{
		"m": 90,
		"s": 85,
		"e": 88,
	}
	m3 := map[string]int{
		"m": 40,
		"s": 55,
		"e": 78,
	}

	// fmt.Println("Maps equal:", fmt.Sprintf("%v", m2) == fmt.Sprintf("%v", m3))
	fmt.Println(maps.Equal(m2,m3))



	////// range in go ///////

	// iterating over data structures 

	numss := []int{6,7,8}
	sum := 0

	for i, v:= range numss{
		sum+=v
		fmt.Println("Index:", i, "Value:", v)
	}

	for k, v:= range m1{
		fmt.Println("Key:", k, "Value:", v)
	}

	var s string = "hello"
	for i, v:= range s{
		fmt.Println("Index:", i, "Value:",v) // v is unicode code point rune
		                                     // starting byte of rune
											 // asci upto 255 take 1 byte
											 // beyond that take more than 1 byte
											 
	}

	 

	
}