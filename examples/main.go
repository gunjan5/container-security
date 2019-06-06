package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var reasons = []string{
	"my cat is depressed",
	"my neighbor's mom is sick",
	"it's too cold outside",
}

func main() {
	http.HandleFunc("/", WFHServer)
	http.ListenAndServe(":8080", nil)

}

func WFHServer(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(reasons)
	fmt.Printf("[DEBUG]: Printing reason: %s\n", reasons[n])
	fmt.Fprintf(w, "Hi Ivan, I need to work from home today because: %s!", reasons[n])
}
