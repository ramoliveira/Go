package main

import "fmt"

func main() {
	// For
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// Same as before
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// If-Else
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	// If-ElseIf
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("divisible by 2")
		} else if i%3 == 0 {
			fmt.Println("divisible by 3")
		} else if i%5 == 0 {
			fmt.Println("divisible by 5")
		}
	}

	// Switch
	for i := i; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		default:
			fmt.Println(i)
		}
	}
}
