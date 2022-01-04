package main

import (
	"controllers/homecontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/home", homecontroller.Index)
	http.HandleFunc("/home/index", homecontroller.Index)

	http.ListenAndServe(":3000", nil)
}
