package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()

	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock() 
	}
	
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var wg sync.WaitGroup

	cache := NewCache(GetFibonacci)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter numbers separated by spaces to calculate Fibonacci (e.g., 42 40 41):")
	input, _ := reader.ReadString('\n')
	fmt.Println("")
	input = strings.TrimSpace(input)
	numbers := strings.Split(input, " ")

	for _, numStr := range numbers {
		wg.Add(1)

		n, err := strconv.Atoi(numStr)
		if err != nil {
			log.Println("Invalid number:", numStr)
			continue
		}

		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(index, value, time.Since(start))
		}(n)

	}
	wg.Wait()
}
