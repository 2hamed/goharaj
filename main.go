package main

import (
	"net/http"
)

func main() {
	r := GetRouter()
	h := Apply(r)
	http.ListenAndServe(":8000", h)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
