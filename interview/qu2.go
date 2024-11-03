package interview

import "sync"

func Pf(x int, wg sync.WaitGroup, channel chan int) int {
	defer wg.Done()
	channel <- x * x
	return x * x
}
