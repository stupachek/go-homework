package main

import "fmt"

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	ch := make(chan int)
	for _, arr := range n {
		go sum(arr, ch)
	}
	res := 0
	for i := 0; i < len(n); i++ {
		res += <-ch
	}

	fmt.Printf("result: %v", res)

}

func sum(arr []int, ch chan int) {
	result := 0
	for _, i := range arr {
		result += i
	}
	ch <- result
}
