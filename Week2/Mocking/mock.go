package mock

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const write = "write"
const sleep = "sleep"
const countdownStart = 3

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

type CountdownOperaxtionsSpy struct {
	Calls []string
}

func (s *CountdownOperaxtionsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperaxtionsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {

}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}
	for i := countdownStart; i > 0; i-- {
		fmt.Fprint(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
