package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // m  sync.Mu/tex

func a() {
	fmt.Println("a")

	defer wg.Done()
}

func b() {
	fmt.Println("b")

	defer wg.Done()
}

func c() {
	fmt.Println("c")

	defer wg.Done()
}

func main() {
	wg.Add(1)
	go a()
	wg.Wait()
	wg.Add(1)
	go b()
	wg.Wait()
	wg.Add(1)
	go c()
	wg.Wait()
}
