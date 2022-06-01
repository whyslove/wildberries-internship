package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheImplementation struct {
	mu    sync.RWMutex
	cache map[string]int
}

func (ci *CacheImplementation) Write() {
	ci.mu.Lock()
	defer ci.mu.Unlock()
	ci.cache["123"] = 1
}

func (ci *CacheImplementation) Read() {
	ci.mu.RLock()
	defer ci.mu.RUnlock()
	fmt.Println(ci.cache["123"])
}

func main() {
	ci := CacheImplementation{cache: make(map[string]int)}
	go ci.Read()
	go ci.Write()
	go ci.Write()
	go ci.Read()
	go ci.Write()
	go ci.Read()
	go ci.Read()
	go ci.Write()

	time.Sleep(time.Minute)
}
