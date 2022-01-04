package channels

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("TRAINING DEFER: Delay execution ")
	fmt.Println("Channels are for sending data")
	fmt.Printf("Sintax: \n Create: make(chan val-type) \n Send : channel <- \n Receives:  <-channel \n ")

	SimpleChannel()
	ChannelBuffering()
	CreateAChannel()

	done := make(chan bool, 1)
	go ChannelSynchronization(done)
	<-done
}

func SimpleChannel() {
	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)
}

func ChannelBuffering() {
	message := make(chan string, 2)
	message <- "first"
	message <- "second"

	fmt.Println(<-message)
	fmt.Println(<-message)
}

func CreateAChannel() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func ChannelSynchronization(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
