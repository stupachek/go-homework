package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	var wg sync.WaitGroup

	// Ваша реалізація
	for i, j := range n {
		wg.Add(1)
		go func(i int, j []int) {
			defer wg.Done()
			fmt.Printf("slice %v: %v\n", i+1, sum(j))
		}(i, j)

	}
	wg.Wait()
}

func sum(arr []int) int {
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}
