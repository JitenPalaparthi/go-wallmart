package main

import "fmt"

func main() {

	var map1 map[string]any // declaration but not a instantiation

	map1 = make(map[string]any, 100)

	map1["560086"] = "Bangalore-Yeshvantpur"
	map1["560096"] = "Bangalore-Rajaji nagar"
	map1["560042"] = "Bangalore-BTM"
	// map1["560082"] = true
	// e08d99312033ce18299e 10c77822ec28a2bf7560
	// segment               offset
	// num<8                 index
	// 4                     2

	map2 := map1

	var mapptr *map[string]any = &map1
	(*mapptr)["560086"] = "Something else"

	v, ok := map1["560086"]
	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("no key is available")
	}

	v, ok = map1["560123"]
	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("no key is available")
	}

	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Println("Len:", len(map1))
	//fmt.Println("Cap:", cap(map1))
	//clear(map1)
	delete(map2, "560086")
	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}
	clear(map1)
	fmt.Println("after clear")
	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Println("Len:", len(map1))
}

// any type that implements == can be a key
// hashfunc
// hashkey
// map creats 8 buckets , bucket is a slice
// can check whether map is nil or not
// maps are not ordered
// map is very efficinet
// map is not thread safe
// use make function to instantiate the map
// use range loop to fetch the map
