package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello World!"
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", helloWorld)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
