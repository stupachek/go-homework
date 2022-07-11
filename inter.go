package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	dogFood            = 2
	minDogNormalWeight = 3
	maxDogNormalWeight = 50
	catFood            = 7
	minCatNormalWeight = 2
	maxCatNormalWeight = 10
	cowFood            = 25
	minCowNormalWeight = 10
	maxCowNormalWeight = 50
)

type Dog struct {
	name, species string
	weight        float64
	eatable       bool
}

func (d Dog) getWeight() (float64, error) {
	weight := d.weight
	if weight < minDogNormalWeight {
		return 0, errors.New("вага тварини замала")
	} else if weight > maxDogNormalWeight {
		return 0, errors.New("вага тварини завелика")
	}
	return weight, nil
}

func (d Dog) getSpecies() string {
	return d.species
}
func (d Dog) getName() string {
	return d.name
}

func (d Dog) getEateble() (bool, error) {
	if d.eatable {
		return false, errors.New("собака вважається їстівним, але таким не має бути")
	}
	return d.eatable, nil
}

func (d Dog) String() string {
	return "собака"
}

func (d Dog) getFoodMonth() float64 {
	w, _ := d.getWeight()
	return w / dogFood
}

type Cat struct {
	name, species string
	weight        float64
	eatable       bool
}

func (c Cat) getWeight() (float64, error) {
	weight := c.weight
	if weight < minCatNormalWeight {
		return 0, errors.New("вага тварини замала")
	} else if weight > maxCatNormalWeight {
		return 0, errors.New("вага тварини завелика")
	}
	return weight, nil
}

func (c Cat) String() string {
	return "кіт"
}

func (c Cat) getSpecies() string {
	return c.species
}

func (c Cat) getName() string {
	return c.name
}

func (c Cat) getEateble() (bool, error) {
	if c.eatable {
		return false, errors.New("кіт вважається їстівним, але таким не має бути")
	}
	return c.eatable, nil
}

func (c Cat) getFoodMonth() float64 {
	w, _ := c.getWeight()
	return w / catFood
}

type Cow struct {
	name, species string
	weight        float64
	eatable       bool
}

func (c Cow) getWeight() (float64, error) {
	weight := c.weight
	if weight < minCowNormalWeight {
		return 0, errors.New("вага тварини замала")
	} else if weight > maxCowNormalWeight {
		return 0, errors.New("вага тварини завелика")
	}
	return weight, nil
}

func (c Cow) String() string {
	return "корова"
}

func (c Cow) getSpecies() string {
	return c.species
}

func (c Cow) getEateble() (bool, error) {
	if !c.eatable {
		return false, errors.New("корова вважається неїстівною, але такою не має бути")
	}
	return true, nil
}

func (c Cow) getName() string {
	return c.name
}

func (c Cow) getFoodMonth() float64 {
	w, _ := c.getWeight()
	return w / cowFood
}

type Animal interface {
	getEateble() (bool, error)
	getWeight() (float64, error)
	getFoodMonth() float64
	getName() string
	getSpecies() string
	fmt.Stringer
}

func getInfoAnimal(animals []Animal) (allFeed float64, err error) {
	for _, animal := range animals {
		if err := validate(animal); err != nil {
			fmt.Printf("для %v з ім'ям %v: %v\n\n", animal, animal.getName(), err)
		} else {
			w, _ := animal.getWeight()
			fmt.Printf("%v по імені \"%v\": важить %vкг, %vкг корму потрібно на місяць\n\n", animal.String(), animal.getName(), w, animal.getFoodMonth())
			allFeed += animal.getFoodMonth()
		}
	}
	return allFeed, nil
}

func validate(animal Animal) error {
	if err := validateType(animal); err != nil {
		return fmt.Errorf("провалена валідація: %v ", err)
	}

	if err := validateWeight(animal); err != nil {
		return fmt.Errorf("провалена валідація: %v ", err)
	}

	if err := validateEatable(animal); err != nil {
		return fmt.Errorf("провалена валідація: %v ", err)
	}

	return nil
}

func validateType(animal Animal) error {
	if animal.getSpecies() != animal.String() {
		return fmt.Errorf("%v насправді %v ", strings.ToLower(animal.getName()), animal.String())
	}
	return nil
}

func validateWeight(animal Animal) error {
	if _, err := animal.getWeight(); err != nil {
		return fmt.Errorf("вага не в нормі: %v ", err)
	}
	return nil
}

func validateEatable(animal Animal) error {
	if _, err := animal.getEateble(); err != nil {
		return fmt.Errorf("тварина не пройшла перевірки на їстівність: %v", err)
	}
	return nil
}

func main() {
	animals := []Animal{Cat{
		name:    "Васька",
		species: "собака",
		weight:  2,
		eatable: false,
	},
		Cat{
			name:    "Кузя",
			species: "кіт",
			weight:  5,
			eatable: false,
		},
		Dog{
			name:    "Гав",
			species: "собака",
			weight:  1234,
			eatable: false,
		},
		Cow{
			name:    "Мілка",
			species: "корова",
			weight:  27,
			eatable: false,
		},
		Dog{
			name:    "Іван",
			species: "собака",
			weight:  15,
			eatable: true,
		},
		Cow{
			name:    "Маруська",
			species: "корова",
			weight:  23,
			eatable: true,
		},
	}
	allFeed, _ := getInfoAnimal(animals)
	fmt.Printf("%vкг корму на ферму", allFeed)

}
