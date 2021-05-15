package main

import "fmt"

func main() {
	fmt.Println("\n############")
	fmt.Println("# NUMBERS  #")
	fmt.Println("############")

	/*
		Integers type:

		Unsigned Integer:
		Meaning: Contains only positive numbers or zero

		uint8
		uint16
		uint32
		uint64

		Signed Integer:

		int8
		int16
		int32
		int64

		Simple:

		int
	*/

	/*
		Floating type:

		Simple Floating:

		float
		float32
		float64

		Complex numbers:

		complex64
		complex128

	*/

	fmt.Println("\tInteger sum:\n\t1 + 1 =", 1+1)

	fmt.Println("\tFloating sum:\n\t1 + 1 =", 1.0+1.0)

	fmt.Println("\n############")
	fmt.Println("# STRINGS  #")
	fmt.Println("############")

	/*
		String type:
		Meaning: Sequence of characters

		"Hello World"
		`Hello World`
	*/

	fmt.Println(len(("Hello World")))

	fmt.Println("Hello World"[1])

	fmt.Println("Hello " + "World!")

	fmt.Println("\n############")
	fmt.Println("# BOOLEAN  #")
	fmt.Println("############")

	/*
		Booleans type:

		Logical operators:

		&& and
		|| or
		! not

		true
		false
	*/

	fmt.Println(true && true)

	fmt.Println(true && false)

	fmt.Println(true || true)

	fmt.Println(true && false)

	fmt.Println(!true)
}
