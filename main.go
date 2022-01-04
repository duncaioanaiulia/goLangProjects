package main

import ()

func main() {
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServer(":5000", server))
}
