package ex5

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex = sync.RWMutex{}
var wg = sync.WaitGroup{}
var countUp = 0
var countDown = 100

func main() {
	BackAndForth()
	//	DecrementIncrement()
}

func DecrementIncrement() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 100; i++ {
		wg.Add(2)
		mutex.RLock()
		go decrement()
		mutex.Lock()
		go increment()
	}
	wg.Wait()
}

func decrement() {
	countDown--
	fmt.Printf("Decrement %v   ;  ", countDown)
	mutex.RUnlock()
	wg.Done()
}

func increment() {
	countUp++
	fmt.Printf("Increment %v \n", countUp)
	mutex.Unlock()
	wg.Done()
}

func BackAndForth() {
	chBack := make(chan int, 10)
	chForth := make(chan int, 10)

	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go back(chBack, i)
		mutex.Lock()
		go forth(chBack, chForth)
		if i == 10 {
			close(chBack)
			close(chForth)
		}
	}
	wg.Wait()

}

func back(ba chan<- int, msg int) {
	ba <- msg
	fmt.Printf("Back %v\n", msg)
	mutex.RUnlock()
	wg.Done()
}

func forth(ba <-chan int, fo chan<- int) {
	msg := <-ba
	fo <- msg
	fmt.Printf("Forth %v\n", msg)
	mutex.Unlock()
	wg.Done()
}
