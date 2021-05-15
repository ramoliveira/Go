package main

import "fmt"

func zero(x int) {
	x = 0
}

func zero_pointer(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)

	zero_pointer(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
