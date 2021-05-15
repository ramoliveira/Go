package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Function returning multiple values
func f() (int, int) {
	return 5, 6
}

// Variadic functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Function that returns function
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Defer

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func third() {
	defer second()
	first()
	fmt.Println("3rd")
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	x, y := f()
	fmt.Println(x, y)

	z := []int{1, 2, 3, 4}
	fmt.Println(add(z...))

	fmt.Println(add(1, 2, 3, 4, 5))

	// Closure
	minus := func(x, y int) int {
		return x - y
	}
	fmt.Println(minus(2, 1))

	a := 0
	increment := func() int {
		a++
		return a
	}
	fmt.Println(increment())

	nexEven := makeEvenGenerator()
	fmt.Println(nexEven()) // 0
	fmt.Println(nexEven()) // 2
	fmt.Println(nexEven()) // 4

	fmt.Println(factorial(4))

	third()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
