package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Shape interface { //Элемент
	getType() string
	accept(Visitor)
}

type Square struct { //Конкретный элемент
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct { //Конкретный элемент
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct { //Конкретный элемент
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

type Visitor interface { //Посетитель
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}

type AreaCalculator struct { //Конкретный посетитель
	area int // Считает площадь

}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct { //Конкретный посетитель
	x int //Считает координату центар фигуры
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() { //main
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
