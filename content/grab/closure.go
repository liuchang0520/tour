package main

import "sync"

func arrayMultiply(vals []int, multiplier int) []int {
	result := make([]int, len(vals))
	var wg sync.WaitGroup
	for i, v := range vals {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = v * multiplier
		}()
	}
	wg.Wait()
	return result
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7}
	mul := 2
	res := arrayMultiply(vals, mul)
	for i, v := range res {
		if v != (i+1)*2 {
			panic("something went wrong")
		}
	}
	println("all good!")
}
