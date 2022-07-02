package main

import "fmt"

type Rectangle struct {
	a, b int
}

type Square struct {
	a int
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

func (rectangle Rectangle) SquaresContain(square Square) int {
	squaresContain := 0
	for rectangle.a >= square.a {
		rectangle.a -= square.a
		for rectangle.b >= square.a {
			rectangle.b -= square.a
			squaresContain++
		}
	}
	return squaresContain
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
	fmt.Println("_______________________")

	rec.Drow(true)
	fmt.Println("_______________________")

	rec2 := Rectangle{
		a: 2,
		b: 4,
	}
	fmt.Printf("Is rec > rec2? %v\n", rec.IsBiggerArea(rec2))
	fmt.Println("_______________________")

	square := Square{
		a: 2,
	}
	fmt.Printf("How many squares in rec2? - %v\n", rec2.SquaresContain(square))
	fmt.Println("_______________________")

	square2 := Square{
		a: 3,
	}
	fmt.Printf("How many squares2 in rec? - %v\n", rec.SquaresContain(square2))
}
