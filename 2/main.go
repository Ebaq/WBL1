package main

import (
	"fmt"
	"sync"
)

func square(index, num int, res []int, wg *sync.WaitGroup) {
	defer wg.Done()
	res[index] = num * num
}

func concurrentSquare(nums []int) []int {
	res := make([]int, len(nums))
	var wg sync.WaitGroup

	for i, num := range nums {
		wg.Add(1)
		go square(i, num, res, &wg)
	}

	wg.Wait()
	return res
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	res := concurrentSquare(nums)

	fmt.Println("result:", res)
}