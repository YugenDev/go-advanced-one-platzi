package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFib(n int) int {
	fmt.Printf("Calculating Expensive fib for %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	inProgress map[int]bool
	IsPending  map[int][]chan int
	lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.lock.RLock()
	exists := s.inProgress[job]
	if exists {
		s.lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.lock.Unlock()
		fmt.Printf("Waiting for response job: %d\n", job)
		res := <-response
		fmt.Printf("Response done, result: %d\n", res)
		return
	}

	s.lock.RUnlock()

	s.lock.Lock()
	s.inProgress[job] = true
	s.lock.Unlock()

	fmt.Printf("Calculate fib for %d\n", job)
	result := ExpensiveFib(job)

	s.lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
		fmt.Printf("Result sent to all pending workers ready, job: %d\n", job)
	}

	s.lock.Lock()
	s.inProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.lock.Unlock()

}

func NewService() *Service {
	return &Service{
		inProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	Service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 3, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, n := range jobs {
		go func(job int){
			defer wg.Done()
			Service.Work(job)
		}(n)
	}
	wg.Wait()
}
