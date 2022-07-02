package main

import "fmt"

type Rectangle struct {
	a, b int
}

// position == true drow horizontally or position == false drow vertically
func (rectangle Rectangle) Drow(position bool) {
	var width, height int

	if position {
		width, height = MaxMin(rectangle.a, rectangle.b)
	} else {
		height, width = MaxMin(rectangle.a, rectangle.b)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(0)
		}
		fmt.Println()
	}

}

func (rectangle *Rectangle) ChangeSize(multiplier int) {
	rectangle.a *= multiplier
	rectangle.b *= multiplier
}

func (rectangle Rectangle) IsBiggerArea(compared Rectangle) bool {
	return rectangle.GetArea() > compared.GetArea()
}

func (rectangle Rectangle) GetArea() int {
	return rectangle.a * rectangle.b
}

func MaxMin(a, b int) (max, min int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	rec := Rectangle{
		a: 3,
		b: 5,
	}
	rec.Drow(false)
	rec.ChangeSize(2)
	rec.Drow(true)
}
