package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("This is foo function.")
	wg.Done()
}

func bar() {
	fmt.Println("This is bar function.")
	wg.Done()
}
