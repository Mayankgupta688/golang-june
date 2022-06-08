package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
}

func (shape Circle) getCircumference() float32 {
	var floatSmall float32 = math.Pi
	var floatBig complex128 = math.Pi

	fmt.Println(floatSmall)
	fmt.Println(floatBig)

	return 2 * math.Pi * shape.radius
}

type Rectangle struct {
	length float32
	breath float32
}

func (shape Rectangle) getCircumference() float32 {
	return 2 * (shape.length + shape.breath)
}

type Square struct {
	length float32
}

func (shape Square) getCircumference() float32 {
	return 4 * shape.length
}

func main() {
	myCircle := Circle{10}
	data := calculateData(myCircle)
	fmt.Println(data)

	myRectangle := Rectangle{10, 20}
	fmt.Println(calculateData(myRectangle))

	mySquare := Square{100}
	fmt.Println(calculateData(mySquare))
}

type IShapes interface {
	getCircumference() float32
}

func calculateData(shape IShapes) float32 {
	returnedData := shape.getCircumference()
	return returnedData
}
