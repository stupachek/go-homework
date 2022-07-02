package main

import "fmt"

type Rectangle struct {
	a, b int
}

type Square struct {
	a int
}

func (r Rectangle) Drow(horizontally bool) {
	var width, height int

	if horizontally {
		width, height = MaxMin(r.a, r.b)
	} else {
		height, width = MaxMin(r.a, r.b)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(0)
		}
		fmt.Println()
	}

}

func (r *Rectangle) ChangeSize(multiplier int) {
	r.a *= multiplier
	r.b *= multiplier
}

func (r Rectangle) BiggerThan(compared Rectangle) bool {
	return r.GetArea() > compared.GetArea()
}

func (r Rectangle) GetArea() int {
	return r.a * r.b
}

func (r Rectangle) SquaresContain(square Square) int {
	squaresContain := 0
	width, heigth := MaxMin(r.a, r.b)

	for heigth >= square.a {
		heigth -= square.a

		for width >= square.a {
			width -= square.a
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
	fmt.Printf("Is rec > rec2? %v\n", rec.BiggerThan(rec2))
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
	fmt.Println("_______________________")

	rec = Rectangle{
		a: 5,
		b: 3,
	}
	fmt.Printf("How many squares in rec? - %v\n", rec.SquaresContain(square))
}
