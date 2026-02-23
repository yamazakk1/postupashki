package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", userHandler)

	http.ListenAndServe(":8080", nil)

}

func userHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(7 * time.Second)
	fmt.Fprintf(w, "User page")
}
