package main

import "fmt"

func main() {
	// Declaration and attributing
	var x string = "Hello World"

	fmt.Println(x)

	// Declaration and then attributing
	var y string
	y = "Hello World"
	fmt.Println(y)

	// Declaration and changing value
	var z string
	z = "First"
	fmt.Println(z)
	z = "Second"
	fmt.Println(z)

	// Sum of strings
	var a string
	a = "first"
	fmt.Println(a)
	a = a + " second"
	fmt.Println(a)

	// Same as before
	var b string
	b = "first"
	fmt.Println(a)
	b += " second"
	fmt.Println(a)

	// Creating and starting value
	c := "Hello World"
	fmt.Println(c)

	// Constants
	const d string = "Hello World"
	fmt.Println(d)

	// Defining Multiple Variables
	var (
		e = 5
		f = 10
		g = 15
	)
	fmt.Println(e + f + g)

	// Defining Multiple Constants
	const (
		h = 1
		j = 2
		k = 3
	)
	fmt.Println(h + j + k)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
