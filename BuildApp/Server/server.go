package server

import "net/http"
import "fmt"
import "strings"

// type Handler interface{
// 	ServeHTTP(ResponseWriter, *Request)
// }

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}
}
