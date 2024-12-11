package main

import (
	"errors"
	"fmt"
)

func main() {

	var mymap1 mymap

	mymap1 = make(mymap)

	mymap1["1"] = "bangalore"
	mymap1["2"] = "hyd"
	mymap1["3"] = "trivandrum"
	mymap1["4"] = "chennai"
	mymap1["5"] = "mumbai"
	mymap1["6"] = "delhi"

	for k, v := range mymap1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if keys, err := mymap1.GetKeys(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keys)
	}

	if valus, err := mymap1.GetValues(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valus)
	}

	var map1 map[string]any
	map1 = make(map[string]any)

	map1["1"] = "bangalore"
	map1["2"] = "hyd"
	map1["3"] = "trivandrum"
	map1["4"] = "chennai"
	map1["5"] = "mumbai"
	map1["6"] = "delhi"

	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	if keys, err := mymap(map1).GetKeys(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keys)
	}

	if valus, err := mymap(mymap1).GetValues(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valus)
	}

	var f1 myfunc = func(i1, i2 int) int {
		return i1 + i2
	}

	fmt.Println(f1.Sq(10, 20))

	var f2 myfunc = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Println(f2.Sq(10, 20))
}

type mymap map[string]any

func (m mymap) GetKeys() ([]string, error) {

	if m == nil {
		return nil, errors.New("nil map")
	}

	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys, nil
}

func (m mymap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}

	values := make([]any, len(m))

	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}

	return values, nil
}

type myfunc func(int, int) int

func (m myfunc) Sq(a, b int) int {
	return m(a, b) * m(a, b)
}
