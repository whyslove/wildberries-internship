package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	c := Counter{value: 0}
	wg := sync.WaitGroup{}
	wg.Add(6)
	go func() {
		c.mu.Lock()
		c.value++
		defer func() { c.mu.Unlock(); wg.Done() }()
	}()
	go func() {
		c.mu.Lock()
		c.value++
		defer func() { c.mu.Unlock(); wg.Done() }()
	}()
	go func() {
		c.mu.Lock()
		c.value++
		defer func() { c.mu.Unlock(); wg.Done() }()
	}()
	go func() {
		c.mu.Lock()
		c.value++
		defer func() { c.mu.Unlock(); wg.Done() }()
	}()
	go func() {
		c.mu.Lock()
		c.value++
		defer func() { c.mu.Unlock(); wg.Done() }()
	}()
	go func() {
		c.mu.Lock()
		c.value++
		defer func() { c.mu.Unlock(); wg.Done() }()
	}()
	wg.Wait()
	fmt.Printf("counter %d\n", c.value)

}
