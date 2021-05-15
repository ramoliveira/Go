package main

import "fmt"

func main() {
	// Array
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)

	// Len
	var total_2 float64 = 0
	for i := 0; i < len(y); i++ {
		total_2 += y[i]
	}
	fmt.Println(total_2 / float64(len(y)))

	// Another way of doing For
	var total_3 float64 = 0
	// for i, value := range x {} => Will not function, because i will not be used
	for _, value := range y {
		total_3 += value
	}
	fmt.Println(total_3 / float64(len(x)))

	// Another syntax of creating array
	z := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println(z)

	/*
		Slices

		A slice is a segment of an array. Like arrays slices are indexable
		and have a length. Unlike arrays this length is allowed to change.
	*/

	var a []float64
	b := make([]float64, 5)
	c := make([]float64, 5, 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Low-High Expression
	arr := [5]float64{1, 2, 3, 4, 5}
	d := arr[1:4]
	/*
		arr[0:] == arr[0:len(arr)]
		arr[:5] == arr[0:5]
		arr[:] == arr[0:len(arr)]
	*/

	fmt.Println(d)

	// Slice functions
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)

	/*
		Map

		Map is an unordered collection of key-value pairs. Also know as an
		associative array, a hash table or dictionary, maps are used to look up
		a value by its associated key.
	*/

	e := make(map[string]int)
	e["key"] = 10
	fmt.Println(e["key"])

	f := make(map[int]int)
	f[1] = 10
	fmt.Println(f[1])

	// Delete
	delete(f, 1)

	// When value is not found
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
