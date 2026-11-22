// Package run has the helper every benchmark uses to fan requests out.
package run

import "sync"

// Map applies f to every item using the given number of workers and keeps the order.
func Map[T, R any](items []T, workers int, f func(int, T) R) []R {
	out := make([]R, len(items))
	jobs := make(chan int)
	var wg sync.WaitGroup
	for range max(workers, 1) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				out[i] = f(i, items[i])
			}
		}()
	}
	for i := range items {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return out
}
