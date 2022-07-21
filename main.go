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
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	for i := range n {
		wg.Add(1)
		go func(i int) {
		  defer wg.Done()
		  fmt.Printf("slice %v: %v\n", i+1, sum(n[i]))
		}(i)
	}
	wg.Wait()
}

func sum(s []int) int {
	var sm int
	for _, v := range s {
		sm += v
	}
	return sm
}
