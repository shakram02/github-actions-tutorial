package main

import (
	"fmt"
	"net/http"
)

type Greeter struct{}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("hello: %v", name)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var greeter = Greeter{}

	name := r.FormValue("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, greeter.Greet(name))

	// fmt.Fprintf(w, greeter.Greet("world"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9010", nil)
}
