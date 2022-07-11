package main

import "fmt"

const (
	dogFood = 2
	catFood = 7
	cowFood = 25
)

type Dog struct {
	weight float64
}

func (d Dog) getWeight() float64 {
	return d.weight
}

func (d Dog) String() string {
	return "dog"
}

func (d Dog) getFoodMonth() float64 {
	return d.getWeight() / dogFood
}

type Cat struct {
	weight float64
}

func (c Cat) getWeight() float64 {
	return c.weight
}

func (c Cat) String() string {
	return "cat"
}

func (c Cat) getFoodMonth() float64 {
	return c.getWeight() / catFood
}

type Cow struct {
	weight float64
}

func (c Cow) getWeight() float64 {
	return c.weight
}

func (c Cow) String() string {
	return "cow"
}

func (c Cow) getFoodMonth() float64 {
	return c.getWeight() / cowFood
}

type Animal interface {
	getWeight() float64
	getFoodMonth() float64
	fmt.Stringer
}

func getInfoAnimal(animals []Animal) (allFeed float64) {
	for _, animal := range animals {
		fmt.Printf("%s: важить %vкг, %vкг корму потрібно на місяць\n", animal.String(), animal.getWeight(), animal.getFoodMonth())
		allFeed += animal.getFoodMonth()
	}
	return allFeed
}

func main() {
	animals := []Animal{Dog{
		weight: 5,
	}, Cat{
		weight: 3,
	}, Cow{
		weight: 26,
	}, Cow{
		weight: 30,
	}}
	allFeed := getInfoAnimal(animals)
	fmt.Printf("%vкг корму на ферму", allFeed)
}
