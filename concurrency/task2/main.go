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

	ch := make(chan []int)
	go write(ch, n)
	res := make([]int, 0)
	for v := range ch {
		res = append(res, sum(v))
	}
	fmt.Printf("result: %v", sum(res))

}

func write(ch chan []int, n [][]int) {
	for _, arr := range n {
		ch <- arr
	}
	close(ch)
}

func sum(arr []int) int {
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}
