package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("----- TRAINING GOROUTINE:")

	CreateGoroutine()
	SynchronizeMultipleGoroutines()
	SynchronizeMultipleGoroutinesWithMutex()
	PrintThreads()
}

var wg = sync.WaitGroup{}
var counter = 0

func SynchronizeMultipleGoroutines() {
	fmt.Println("----- Synchronize Multiple Goroutines")
	counter = 0
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go printCounter()
		go increment()
	}
	wg.Wait()
}

func printCounter() {
	fmt.Printf("Number #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func CreateGoroutine() {
	fmt.Println("----- Create Goroutines")

	var msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

var mutex = sync.RWMutex{}
var wgMutex = sync.WaitGroup{}
var counterMutex = 0

func SynchronizeMultipleGoroutinesWithMutex() {
	fmt.Println("----- Synchronize Multiple Goroutines With Mutex")
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wgMutex.Add(2)
		mutex.RLock()
		go printCounterUsingMutex()
		mutex.Lock()
		go incrementUsingMutex()
	}
	wgMutex.Wait()
}

func printCounterUsingMutex() {
	fmt.Printf("#%v\n", counterMutex)
	mutex.RUnlock()
	wgMutex.Done()
}

func incrementUsingMutex() {
	counterMutex++
	mutex.Unlock()
	wgMutex.Done()
}

func PrintThreads() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
