package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

// Same as circleArea function, but uses Circle struct
func cArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func c2Area(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// Struct
type Circle struct {
	x float64
	y float64
	r float64
}

/*
	Can also be:

	type Circle struct {
		x, y, r float64
	}
*/

// Methods of Circle
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

/*
	Embedded types can turn the struct above into:

	type Android struct {
		Person
		Model string
	}
*/

type EmbeddedAndroid struct {
	Person
	Model string
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	/*
		Ways of declaring a new Circle struct

		var c1 Circle
		c2 := new(Circle)
		c3 := Circle{x: 0, y: 0, r: 5}
		c4 := Circle{0, 0, 5}
	*/

	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(cArea(c))

	fmt.Println(c2Area(&c))

	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	a := new(Android)
	a.Person.Talk()

	a2 := new(EmbeddedAndroid)
	a2.Talk()

	fmt.Println(totalArea(&c, &r))
}
