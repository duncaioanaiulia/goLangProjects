package controlFlow

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	InitDefer()
	Robots()
	InitPanic()
	//PanicResponse()
	OrderOfExecution(false)
	InitRecover(true)
}

func InitDefer() {
	fmt.Println("TRAINING DEFER: Delay execution ")
	fmt.Println("WARNING: be careful in loops")
	a := "Start"
	defer fmt.Println(a)
	a = "End"
}

func Robots() {
	fmt.Println("TRAINING DEFER: Robots Example:")

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal((err))
	}
	fmt.Println(string(robots[0]))
}

func InitPanic() {
	fmt.Println("TRAINING PANIC: ")
	a, b := 1, 0
	//ans := a / b
	fmt.Println(a, b)
}

func PanicResponse() {
	fmt.Println("TRAINING PANIC: Panic Response")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func OrderOfExecution(isPanicking bool) {
	fmt.Println("TRAINING PANIC: Order Of Execution ")
	fmt.Println("WARNING: defer func it will be executed befor panic, to not  to lose resources")

	fmt.Println(("start"))
	defer fmt.Println(("defer"))
	if isPanicking {
		panic("panic")
	} else {
		fmt.Println(("end"))
	}
}

func InitRecover(isPanicking bool) {
	fmt.Println("TRAINING Recover: ")
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	if isPanicking {
		panic("PANIC!!!!!")
	} else {
		fmt.Println("done panic")
	}
}
